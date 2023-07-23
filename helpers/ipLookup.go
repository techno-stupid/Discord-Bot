package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type GeoLocation struct {
	IsValid     bool    `json:"is_valid"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	RegionCode  string  `json:"region_code"`
	Region      string  `json:"region"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Address     string  `json:"address"`
}

func IpLookup(ip string) (GeoLocation, error) {

	url := "https://api.api-ninjas.com/v1/iplookup?address=" + ip
	apiKey := os.Getenv("JOKE_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return GeoLocation{}, err
	}

	req.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return GeoLocation{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading response failed:", err)
		return GeoLocation{}, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return GeoLocation{}, err
	}
	var geoLocation GeoLocation
	err = json.Unmarshal(body, &geoLocation)
	if err != nil {
		log.Fatal("Something went wrong while decoding assets data")
	}
	fmt.Println(geoLocation)

	if !geoLocation.IsValid {
		return GeoLocation{}, fmt.Errorf("Invalid IP")
	}

	return geoLocation, nil
}

func IsIP(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil
}
