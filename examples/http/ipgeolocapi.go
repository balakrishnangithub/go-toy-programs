package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// IPGeolocationAPI generated with the help of https://mholt.github.io/json-to-go/
type IPGeolocationAPI struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func mainIPGeolocationAPI() {
	var err error
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		panic(err)
	}

	var decoder *json.Decoder = json.NewDecoder(resp.Body)

	// https://stackoverflow.com/questions/12329493/use-of-new-vs-var-in-go
	// var ipGeoLocAPI IPGeolocationAPI
	// decoder.Decode(&ipGeoLocAPI)
	var ipGeoLocAPI *IPGeolocationAPI = new(IPGeolocationAPI)
	err = decoder.Decode(ipGeoLocAPI)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ipGeoLocAPI)
}
