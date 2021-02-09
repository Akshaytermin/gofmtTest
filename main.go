package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

func AttachUtm(productURL string, utmValue string) {
	url, err := url.Parse(productURL)
	if err != nil {
		log.Fatal(err)
	}
	q := url.Query()

	utmMap := make(map[string]string)
	err = json.Unmarshal([]byte(utmValue), &utmMap)

	fmt.Println(utmMap)

	for key, value := range utmMap {
		q.Add(key, value)
	}

	url.RawQuery = q.Encode()

	fmt.Println(url.String())
}

func main() {
	url := "https//local/asd/"
	AttachUtm(url, "{\"utm_source\":\"email\",\"utm_sourdce\":\"sd\"}")
}
