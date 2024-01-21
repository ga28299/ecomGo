package generate

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserID    string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_TOKEN_KEY")

func TokenGen(email string, firstName string, lastName string, userID string) (string, string, error) {
	fmt.Printf("Secret key is : %v", SECRET_KEY)
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return token, refreshtoken, nil

}

func ValidateToken(token string) (claims *string, msg string) {
	log.Panic("ValidateToken function is not implemented yet")
	return nil, "nil"
}
