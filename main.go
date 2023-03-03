package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/", ArtistsPage)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type Artists struct {
	ID            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Locations     string              `json:"locations"`
	ConcertDates  string              `json:"concertDates"`
	DatesLocation map[string][]string `json:"datesLocations"`
	Relations     string              `json:"relations"`
}

type Relations struct {
	Index []struct {
		ID            int                 `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var (
	Artist   []Artists
	Relation Relations
)

func UnmarshallArtists() error {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	jsonErr := json.Unmarshal(body, &Artist)
	if jsonErr != nil {
		log.Println("jsonErr")
	}
	return nil
}

func UnmarshallRelations() error {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	jsonErr := json.Unmarshal(body, &Relation)
	if jsonErr != nil {
		log.Println("jsonErr")
	}
	return nil
}
