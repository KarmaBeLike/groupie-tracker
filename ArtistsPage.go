package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ArtistsPage(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url[2])
	if r.URL.Path != "/artistspage/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/index.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = UnmarshallArtists()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = ts.Execute(w, Artist[id-1])
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func ID(id int) error {
	if id > 52 || id < 1 {
		return errors.New("out of id")
	}
	return nil
}
