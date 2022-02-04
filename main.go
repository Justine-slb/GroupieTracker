package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)
//Structure artiste
type Artists struct {
	Id           int
	Name         string
	Members      []string
	CreationDate int
	Image        string
	Locations    string
	FirstAlbum   string
}

// Artist page
func artistpage(w http.ResponseWriter, r *http.Request) {
	//recupe donné artiste via URL API
	var artist []Artists
	tmpl := template.Must(template.ParseFiles("template/artist.html"))
	request, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	json.Unmarshal(byteValue, &artist)
	tmpl.Execute(w, artist)
}

func main() {
	//Url Page
	http.HandleFunc("/artist", artistpage)
	//CSS + Picture integration
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	//Server
	if err := http.ListenAndServe(":1993", nil); err != nil {
		log.Fatal(err)
	}
}
