package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Results []Datas `json:"results"`
}

type Datas struct {
	Basics    Basics      `json:"basic"`
	Addresses []Addresses `json:"addresses"`
}
type Basics struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}
type Addresses struct {
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	City            string `json:"city"`
	TelephoneNumber string `json:"telephone_number"`
}

type Details struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Address   Addresses `json:"addresses"`
}

func main() {

	// Api
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://npiregistry.cms.hhs.gov/api/?version=2.1&number=1558444216")
	if err != nil {
		log.Fatalln(err)
	}

	// data
	data, _ := ioutil.ReadAll(resp.Body)

	// parse
	bodyString := string(data)
	fmt.Println("API Response as String:\n" + bodyString)

	fmt.Println("-------------------------------------")

	var res Response
	var doctor Details
	json.Unmarshal(data, &res)

	doctor.FirstName = res.Results[0].Basics.FirstName
	doctor.LastName = res.Results[0].Basics.LastName
	doctor.Gender = res.Results[0].Basics.Gender

	doctor.Address.CountryCode = res.Results[0].Addresses[0].CountryCode
	doctor.Address.CountryName = res.Results[0].Addresses[0].CountryName
	doctor.Address.City = res.Results[0].Addresses[0].City
	doctor.Address.TelephoneNumber = res.Results[0].Addresses[0].TelephoneNumber

	fmt.Println("rse 1", res.Results[0])
	fmt.Println(doctor.FirstName)
	fmt.Println(doctor.LastName)
	fmt.Println(doctor.Gender)
	fmt.Println(doctor.Address.CountryCode)
	fmt.Println(doctor.Address.CountryName)
	fmt.Println(doctor.Address.City)
	fmt.Println(doctor.Address.TelephoneNumber)

}
