package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ApiUrl = "https://groupietrackers.herokuapp.com/api"

// function to return Url contained in the API :https://groupietrackers.herokuapp.com/api
func Get() Todo {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(ApiUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	return todoStruct
}

// function Get to read the API send in params and return Location data struct about the id choosen
func GetLocation(url string) LocationsArray {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var LocationsData LocationsIndex
	json.Unmarshal(bodyBytes, &LocationsData)
	return LocationsData.Index
}
