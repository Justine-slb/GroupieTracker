package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Research(research string, artists ApiHerokuapp) (APIFullData, bool) {
	var answer_research APIFullData
	for _, value := range artists {
		if strings.EqualFold(value.Name, research) == true || strings.EqualFold(value.Name, research) == true || strings.EqualFold(value.Name, research) == true || strings.EqualFold(value.FirstAlbum, research) == true {
			answer_research = value
			return answer_research, true
		} else {
			for _, members := range value.Members {
				if strings.EqualFold(members, research) == true {
					answer_research = value
					return answer_research, true
				}
			}
		}
	}
	return answer_research, false
}

/*search*/
func ChangetUrlForSearch(find bool) string {
	var url string
	if find == true {
		url = "/ArtistPage"
	} else {
		url = "/"
	}
	return url
}

/*function filters*/
func Filters(filters FiltersPost, artists []APIFullData) (ApiHerokuapp, string) {
	errorMessage := ""
	if filters.End != 0 || filters.Start != 0 || filters.FirstAlbumEnd != 0 || filters.FirstAlbumStart != 0 || filters.Artist != "" || filters.Band != "" {
		artistsFilters := Filter_Band_Art(filters, artists)
		artistsFilters = Filters_Creation(filters, artistsFilters)
		artistsFilters = Filter_first_Album(filters, artistsFilters)
		if len(artistsFilters) >= 1 {
			fmt.Println("Artists filtrés")
			artists = artistsFilters
		} else {
			errorMessage = "no matching, Sorry"
		}
	}
	return artists, errorMessage
}

/*function filters */
func Filters_Creation(filters FiltersPost, allArtists ApiHerokuapp) ApiHerokuapp {
	new_artists := []APIFullData{}
	for _, value := range allArtists {
		if filters.End != 0 && filters.Start != 0 {
			if filters.Start <= value.CreationDate && filters.End >= value.CreationDate {
				new_artists = append(new_artists, value)
			}
		} else if filters.End != 0 && filters.Start == 0 {
			if filters.End >= value.CreationDate {
				new_artists = append(new_artists, value)
			}
		} else if filters.End == 0 && filters.Start != 0 {
			if filters.Start <= value.CreationDate {
				new_artists = append(new_artists, value)
			}
		} else {
			new_artists = append(new_artists, value)
		}
	}
	fmt.Println("Le resultat band ou artist est :")
	fmt.Println(new_artists)
	return new_artists
}

func Filter_Band_Art(filters FiltersPost, allArtists []APIFullData) []APIFullData {
	new_artists := []APIFullData{}
	fmt.Println(filters.Band)
	fmt.Println(filters.Artist)
	if filters.Artist == "Artist" && filters.Band == "Band" || (filters.Artist == "" && filters.Band == "") {
		for _, value := range allArtists {
			new_artists = append(new_artists, value)
		}
	} else if filters.Artist == "Artist" && filters.Band == "" {
		for _, value := range allArtists {
			fmt.Print("le nombre de membre est : ")
			fmt.Println(len(value.Members))
			if len(value.Members) == 1 {
				new_artists = append(new_artists, value)
			}
		}
	} else if filters.Artist == "" && filters.Band == "Band" {
		for _, value := range allArtists {
			if len(value.Members) > 1 {
				if filters.Members != 0 {
					fmt.Println(filters.Members)
					fmt.Println(".....................")
					if len(value.Members) >= filters.Members {
						new_artists = append(new_artists, value)
					}
				} else {
					new_artists = append(new_artists, value)
				}
			}
		}
	}
	fmt.Println("Le resultat band ou artist est :")
	for _, value := range new_artists {
		fmt.Println(len(value.Members))
	}

	return new_artists
}

func Filter_first_Album(filters FiltersPost, allArtists []APIFullData) []APIFullData {
	new_artists := []APIFullData{}
	if filters.FirstAlbumStart != 0 || filters.FirstAlbumEnd != 0 {
		for _, value := range allArtists {
			s := strings.Split(value.FirstAlbum, "-")
			firstAlbumDate, _ := strconv.Atoi(s[len(s)-1])
			if filters.FirstAlbumStart != 0 && filters.FirstAlbumEnd != 0 {
				if filters.FirstAlbumStart <= firstAlbumDate && filters.End >= firstAlbumDate {
					new_artists = append(new_artists, value)
				}
			} else if filters.FirstAlbumStart != 0 && filters.FirstAlbumEnd == 0 {
				if filters.FirstAlbumStart <= firstAlbumDate {
					new_artists = append(new_artists, value)
				}
			} else if filters.FirstAlbumStart == 0 && filters.FirstAlbumEnd != 0 {
				if filters.FirstAlbumEnd >= firstAlbumDate {
					new_artists = append(new_artists, value)
				}
			}
		}
		fmt.Println("Le resultat firstAlbum est :")
		fmt.Println(new_artists)
		return new_artists
	} else {
		return allArtists
	}
}

//----------------------Constants---------------------------//

/*Fontion Get DATA from API Vanessa*/

func GetArtistsData() error {
	resp, err := http.Get(APIHerokuappURL + "/artists")
	if err != nil {
		return errors.New("Error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &Artists)
	fmt.Println()
	return nil
}

func GetLocations() error {
	resp, err := http.Get(APIHerokuappURL + "/locations")
	if err != nil {
		return errors.New("Error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &ConcertCountry)
	fmt.Println(ConcertCountry)
	return nil
}

func GetDates() error {
	resp, err := http.Get(APIHerokuappURL + "/dates")
	if err != nil {
		return errors.New("Error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &ConcertDates)
	fmt.Println(ConcertDates)
	return nil
}

func GetRelations() error {
	resp, err := http.Get(APIHerokuappURL + "/relation")
	if err != nil {
		return errors.New("Error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Error by ReadAll")
	}
	json.Unmarshal(bytes, &ConcertLocalisation)
	return nil
}

func GetData() {
	GetArtistsData()
	GetLocations()
	GetDates()
	GetRelations()
	var template APIFullData
	for i := range Artists {
		template.ID = i + 1
		template.Image = Artists[i].Image
		template.Name = Artists[i].Name
		template.Members = Artists[i].Members
		template.CreationDate = Artists[i].CreationDate
		template.FirstAlbum = Artists[i].FirstAlbum
		template.Locations = ConcertCountry.Index[i].Locations
		template.ConcertDates = ConcertDates.Index[i].Dates
		template.Relations = ConcertLocalisation.Index[i].DatesLocation
		APIHerokuapp = append(APIHerokuapp, template)
	}
	fmt.Println(APIHerokuapp)
	return
}

func GetDatabyId(id int) APIFullData {
	var data APIFullData
	for i := range Artists {
		if i == id {
			data.ID = Artists[i].ID
			data.Image = Artists[i].Image
			data.Name = Artists[i].Name
			data.Members = Artists[i].Members
			data.CreationDate = Artists[i].CreationDate
			data.FirstAlbum = Artists[i].FirstAlbum
			data.Locations = ConcertCountry.Index[i].Locations
			data.ConcertDates = ConcertDates.Index[i].Dates
			data.Relations = ConcertLocalisation.Index[i].DatesLocation
			break
		}
	}
	return data
}
