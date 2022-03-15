package function

type Todo struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

/*
//Artist
// structs API https://groupietrackers.herokuapp.com/api/artists
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
*/
/*type Artists []Artist*/

//Locations
// structs API https://groupietrackers.herokuapp.com/api/locations
type Location_Artist struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type LocationsIndex struct {
	Index LocationsArray `json:"index"`
}

type LocationsArray []Location_Artist

//Dates
// structs API https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/dates
/*type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesArray []Dates

type DatesIndex struct {
	Index DatesArray `json:"index"`
}

//Relation
// structs API https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type RelationsArray []Relation

type RelationsIndex struct {
	Index RelationsArray `json:"index"`
}
*/

//Filters
type FiltersPost struct {
	Artist          string
	Band            string
	Members         int
	Start           int
	End             int
	FirstAlbumStart int
	FirstAlbumEnd   int
}

type APIFullData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type MIArtist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type MILocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type LocationData struct {
	Index []MILocation `json:"index"`
}

type MIConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ConcertDateData struct {
	Index []MIConcertDate `json:"index"`
}

type MIRelationDate struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
type RelationData struct {
	Index []MIRelationDate `json:"index"`
}

//----------------------Variable---------------------------//

var Artists []MIArtist
var ConcertCountry LocationData
var ConcertDates ConcertDateData
var ConcertLocalisation RelationData

var APIHerokuapp []APIFullData

type ApiHerokuapp []APIFullData

const APIHerokuappURL string = "https://groupietrackers.herokuapp.com/api"
