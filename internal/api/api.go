package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type University struct {
	Name     string   `json:"name"`
	Country  string   `json:"country"`
	WebPages []string `json:"web_pages"`
}

func GetRequest(country string, c chan<- University) {
	resp, err := http.Get("http://universities.hipolabs.com/search?country=" + country)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data []University
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, element := range data {
		c <- element
	}

}
