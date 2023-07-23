package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Fact struct {
	Content string `json:"fact"`
}

func GetFact() (string, error) {

	url := "https://api.api-ninjas.com/v1/facts?limit=1"
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
	var facts []Fact
	err = json.Unmarshal(body, &facts)
	if err != nil {
		log.Fatal("Something went wrong while decoding assets data")
	}
	return facts[0].Content, nil
}
