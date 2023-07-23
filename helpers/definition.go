package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Definition struct {
	Content string `json:"definition"`
	Word    string `json:"word"`
	Valid   bool   `json:"valid"`
}

func GetDefinition(word string) (string, error) {

	url := "https://api.api-ninjas.com/v1/dictionary?word=" + word
	apiKey := os.Getenv("JOKE_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return "", err
	}

	req.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading response failed:", err)
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return "", err
	}
	var definition Definition
	err = json.Unmarshal(body, &definition)
	if err != nil {
		log.Fatal("Something went wrong while decoding assets data")
	}
	fmt.Println(definition)

	if !definition.Valid {
		return "Not a word", nil
	}

	meaning := strings.Split(definition.Content, "2.")[0]

	return meaning, nil
}
