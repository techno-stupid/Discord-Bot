package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Joke struct {
	Content string `json:"joke"`
}

func GetJoke() (string, error) {

	url := "https://api.api-ninjas.com/v1/jokes?limit=1"
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
	var jokes []Joke
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		log.Fatal("Something went wrong while decoding assets data")
	}
	return jokes[0].Content, nil
}
