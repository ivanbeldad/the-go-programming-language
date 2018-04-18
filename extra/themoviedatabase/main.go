package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// https://api.themoviedb.org/3/movie/550?api_key=

const (
	baseURL     = "https://api.themoviedb.org/3/"
	tokenv3     = ""
	discoverURL = baseURL + "discover/movie" + "?api_key=" + tokenv3
	getURL      = baseURL + "movie/$$$" + "?api_key=" + tokenv3
	searchURL   = baseURL + "search/movie" + "?api_key=" + tokenv3 + "&query=$$$"
)

// Movie ...
type Movie struct {
	ID    int
	Title string
	Vote  float64 `json:"vote_average"`
}

// Result ...
type Result struct {
	Movies []Movie `json:"results"`
}

// MovieCmd ...
type MovieCmd interface {
	execute() []Movie
}

// GetMovieCmd ...
type GetMovieCmd struct{}

func (c GetMovieCmd) execute() []Movie {
	return get()
}

// SearchMovieCmd ...
type SearchMovieCmd struct{}

func (c SearchMovieCmd) execute() []Movie {
	return get()
}

// DiscoverMovieCmd ...
type DiscoverMovieCmd struct{}

func (c DiscoverMovieCmd) execute() []Movie {
	return get()
}

func main() {
	input := os.Args[1] // "get" or "discover" or "search"

	// LO MAS GUARRERO
	if input == "discover" {
		discover()
	} else if input == "get" {
		get()
	} else if input == "search" {
		search()
	}

	// LO SEGUNDO MAS GUARRERO
	switch input {
	case "discover":
		discover()
	case "get":
		get()
	case "search":
		search()
	}

	// MAPS
	actions := map[string]func() []Movie{
		"get":      get,
		"discover": discover,
		"search":   search,
	}
	movies := actions[input]()

	for _, v := range movies {
		fmt.Printf("ID: %d\t Title: %s\tVotes: %.1f\n", v.ID, v.Title, v.Vote)
	}

	// POLYMORPHISM
	cmds := map[string]MovieCmd{
		"get":      GetMovieCmd{},
		"discover": DiscoverMovieCmd{},
		"search":   SearchMovieCmd{},
	}
	movies = cmds[input].execute()
}

func get() []Movie {
	n := os.Args[2] // "number of movie"
	m := Movie{}
	u := strings.Replace(getURL, "$$$", n, 1)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.Fatal(err)
	}
	content, err = json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return []Movie{m}
}

func discover() []Movie {
	resp, err := http.Get(discoverURL)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := Result{}
	err = json.Unmarshal(content, &r)
	if err != nil {
		log.Fatal(err)
	}
	content, err = json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return r.Movies
}

func search() []Movie {
	s := os.Args[2] // "string to search"
	s = url.QueryEscape(s)
	u := strings.Replace(searchURL, "$$$", s, 1)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	r := Result{}
	err = json.Unmarshal(content, &r)
	if err != nil {
		log.Fatal(err)
	}
	content, err = json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if len(r.Movies) == 0 {
		return []Movie{}
	}
	return []Movie{r.Movies[0]}
}
