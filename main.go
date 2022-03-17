package main

import (
	"Groupie-Tracker/function"
	//"encoding/json"
	//"errors"
	"fmt"
	"html/template"
	"math/rand"
	//"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//----------------------Structure---------------------------//
type PageIndex struct {
	Artists    function.ApiHerokuapp
	Featured   function.ApiHerokuapp
	Artist     function.APIFullData
	Location   function.LocationData
	WeekArtist function.APIFullData
	LinkUrl    string
	Error      string
	Random     int
}

type AllArtistsPage struct {
	Artists  function.ApiHerokuapp
	Artist   function.APIFullData
	Location function.LocationData
	LinkUrl  string
	Error    string
	Members  string
	Random   int
}

//------------------------Index-----------------------------------//
func Index(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	location := function.ConcertCountry
	artists := function.APIHerokuapp
	featured := function.Featured_Artist()
	relocation := function.Get_Locations("https://groupietrackers.herokuapp.com/api/locations")
	function.Filters_Location(relocation)
	var errorMessage string
	var find bool
	var url string

	search := r.FormValue("search")
	research := function.APIFullData{}

	if search != "" {
		research, find, url = function.Research(search, artists)
		if find == true {
			_, errorMessage = function.ChangetUrlForSearch(find)
		}
	}
	random := rand.Intn(52)
	_, week := time.Now().ISOWeek()
	weekArtist := artists[week]

	index := PageIndex{artists, featured, research, location, weekArtist, url, errorMessage, random}

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	err := tmpl.ExecuteTemplate(w, "index", index)
	if err != nil {
		fmt.Println(err)
	}
}

//-------------------------FullArtistPage--------------------------//
func FullArtistPage(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	location := function.ConcertCountry
	data := function.APIHerokuapp
	research := function.APIFullData{}
	var find bool
	var url string
	var errorMessage string

	search := r.FormValue("search")
	if search != "" {
		research, find, url = function.Research(search, data)
		_, errorMessage = function.ChangetUrlForSearch(find)
	}
	members, _ := strconv.Atoi(r.FormValue("members"))
	start, _ := strconv.Atoi(r.FormValue("CreationDateStart"))
	end, _ := strconv.Atoi(r.FormValue("CreationDateEnd"))
	firstAlbumStart, _ := strconv.Atoi(r.FormValue("FirstAlbumStart"))
	firstAlbumEnd, _ := strconv.Atoi(r.FormValue("FirstAlbumEnd"))
	filters := function.FiltersPost{r.FormValue("Artist"), r.FormValue("Band"), members, start, end, firstAlbumStart, firstAlbumEnd}

	random := rand.Intn(52)

	data, errorMessage = function.Filters(filters, data)
	artists_Page := AllArtistsPage{data, research, location, url, errorMessage, r.FormValue("members"), random}

	t, _ := template.ParseFiles("template/FullArtist.html")
	err := t.Execute(w, artists_Page)
	if err != nil {
		return
	}
}

//-------------------MoreArtistInfoPage--------------------------//
func MoreinfoArtistPage(w http.ResponseWriter, r *http.Request) {
	var id int
	data := function.APIFullData{}
	Artists := function.APIHerokuapp
	if r.URL.Query().Get("search") != "" {
		ids := r.URL.Query().Get("search")
		search, find, _ := function.Research(ids, Artists)
		fmt.Println("search is ", search)
		if find == true {
			data = function.GetDatabyId(search.ID - 1)
			t, _ := template.ParseFiles("template/moreinfo-Copie.html")
			err := t.Execute(w, data)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			t, _ := template.ParseFiles("template/error.html")
			err := t.Execute(w, Artists)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		id, _ = strconv.Atoi(r.URL.Path[12:])
		data = function.GetDatabyId(id - 1)
		t, _ := template.ParseFiles("template/moreinfo-Copie.html")
		err := t.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	}

	//t, _ := template.ParseFiles("template/MoreSoloInfoPage.html")
	//	t, _ := template.ParseFiles("template/MoreGroupInfoPage.html")

}

//------------------------Filter--------------------------------//
func FilterForSoloOrGroup(Solo bool) []function.APIFullData {
	var data []function.APIFullData
	var num function.APIFullData

	for i := range function.Artists {
		num.Members = function.Artists[i].Members
		if Solo && len(num.Members) == 1 {
			data = append(data, function.GetDatabyId(i))
		} else if !(Solo) && len(num.Members) > 1 {
			data = append(data, function.GetDatabyId(i))
		}
	}
	return data
}

//------------------------Server-------------------------------//
func main() {
	function.GetData()
	server1 := http.NewServeMux()
	server1.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	server1.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	//Url Page
	server1.HandleFunc("/", Index)
	server1.HandleFunc("/Artist", FullArtistPage)
	server1.HandleFunc("/ArtistPage/", MoreinfoArtistPage)
	err := http.ListenAndServe(":8000", server1)
	if err != nil {
		fmt.Println(err)
		return
	}
}
