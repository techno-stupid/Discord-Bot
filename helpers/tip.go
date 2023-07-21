package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type TipData struct {
	Tip string `json:"tip"`
}

var goTips []TipData

func init() {
	jsonFile, err := ioutil.ReadFile("public/json/tips.json")
	if err != nil {
		log.Fatal("Something went wrong while fetching assets data")
	}

	err = json.Unmarshal(jsonFile, &goTips)
	if err != nil {
		log.Fatal("Something went wrong while decoding assets data")
	}
}

func RandomTip() string {

	tips := make([]string, len(goTips))
	for i, tip := range goTips {
		tips[i] = tip.Tip
	}
	if len(tips) == 0 {
		return "No tips available."
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(tips))
	return tips[randomIndex]
}
