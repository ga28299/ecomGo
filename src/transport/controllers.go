package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ga28299/ecomGo/db"
	"github.com/ga28299/ecomGo/generate"
	"github.com/ga28299/ecomGo/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var Validator = validator.New()

func Signup(c *gin.Context) {

	var _, cancel = context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error is here": err.Error()})
		return
	}

	validationErr := Validator.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error is there": validationErr.Error()})
		return
	}

	existingUser, err := db.GetUserByEmail(*user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists", "existing_user_email": existingUser.Email})
		return
	}

	existingPhone, err := db.GetUserByPhone(*user.Phone)
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existingPhone != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists", "existing_user_phone": existingPhone.Phone})
		return
	}

	pwd, err := HashPassword(*user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Check your password"})
		return
	}
	user.Password = &pwd

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = uuid.New()
	user.UserID = user.ID.String()

	// Generate tokens
	token, refreshToken, err := generate.TokenGen(*user.Email, *user.FirstName, *user.LastName, user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	user.Token = &token
	user.RefreshToken = &refreshToken
	user.UserCart = make([]models.Product, 0)
	user.Addresses = make([]models.Address, 0)
	user.Orders = make([]models.Order, 0)
	fmt.Println("Here is whats going tom the db -->>")
	fmt.Printf("id: %v \n", user.ID)
	fmt.Printf("fname:%v \n", *user.FirstName)
	fmt.Printf("lname %v \n", *user.LastName)
	fmt.Printf("email   %v \n", *user.Email)
	fmt.Printf("password   %v \n", *user.Password)
	fmt.Printf("phone   %v \n", *user.Phone)
	fmt.Printf("token   %v \n", *user.Token)
	fmt.Printf("refreshtoken: %v \n", *user.RefreshToken)
	fmt.Printf("created_at   %v \n", user.CreatedAt)
	fmt.Printf("updated_at   %v \n", user.UpdatedAt)
	fmt.Printf("user_id   %v \n", user.UserID)

	err = db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer cancel()
	CookieToken(c, token, refreshToken)
	c.JSON(http.StatusCreated, "Sucessfully created user")

}

func Login(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := db.GetUserByEmail(*user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "That combination not there"})
		return
	}

	valid, msg := VerifyPassword(*existingUser.Password, *user.Password)

	defer cancel()

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		fmt.Println(msg)
		return
	}

	token, refreshToken, err := generate.TokenGen(*existingUser.Email, *existingUser.FirstName, *existingUser.LastName, existingUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	err = db.UpdateTokensForID(existingUser.UserID, token, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	CookieToken(c, token, refreshToken)

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})

}

func ProductViewerAdmin(c *gin.Context) {

}

func SearchProduct(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte("Hello World"))

}

func SearchProductQuery(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func HashPassword(pwd string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(hashed), err
}

func VerifyPassword(userpassword string, givenpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(givenpassword))
	if err != nil {
		return false, "Incorrect Password"
	}
	return true, "Correct Password"

}

func CookieToken(c *gin.Context, token string, refreshToken string) {
	c.SetCookie("token", token, int(time.Hour.Seconds()*24), "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, int(time.Hour.Seconds()*168), "/", "localhost", false, true)
}
