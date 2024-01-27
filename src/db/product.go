package db

import (
	"bytes"
	"log"

	"encoding/json"
)

func ElasticSearch(query []byte) (string, error) {
	return "", nil
}

func GetSuggestions(query string) ([]string, error) {
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"productName.keyword": "*" + query + "*",
			},
		},
		"size": 5,
	}
	jQuery, err := json.Marshal(esQuery)
	if err != nil {
		log.Println("Failed to marshal query")
		return nil, err
	}
	res, err := esCon.Search(
		esCon.Search.WithIndex("products"),
		esCon.Search.WithBody(bytes.NewReader(jQuery)),
	)
	if err != nil {
		log.Println("Failed to execute search query")
		return nil, err
	}
	defer res.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("Failed to parse search results")
		return nil, err
	}

	var suggestions []string
	hits, foundHits := result["hits"].(map[string]interface{})
	if !foundHits {
		log.Println("Failed to find hits")
		return nil, err
	}

	hitsArray, foundHitsArray := hits["hits"].([]interface{})
	if !foundHitsArray {
		log.Println("Failed to find hits array")
		return nil, err
	}

	for _, hit := range hitsArray {
		hitMap, foundHitMap := hit.(map[string]interface{})
		if !foundHitMap {
			log.Println("Error extracting hit map from Elasticsearch response")
			continue
		}

		source, foundSource := hitMap["_source"].(map[string]interface{})
		if !foundSource {
			log.Println("Error extracting source from Elasticsearch hit")
			continue
		}

		// Extract the product name or relevant field from the source
		productName, foundProductName := source["productName"].(string)
		if !foundProductName {
			log.Println("Error extracting product name from Elasticsearch hit")
			continue
		}

		suggestions = append(suggestions, productName)
	}
	return suggestions, nil
}
