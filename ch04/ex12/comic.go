package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	comicFile = "data/comics.json"
	comicURL  = "https://xkcd.com/$$$/info.0.json"
)

var ErrComicNotFound = errors.New("Comic not found")

// Comic ...
type Comic struct {
	Num        int
	Year       int `json:",string"`
	Month      int `json:",string"`
	Day        int `json:",string"`
	Title      string
	Transcript string
	Img        string `json:"ImgUrl"`
}

func getComicURL(i int) (url.URL, error) {
	u := strings.Replace(comicURL, "$$$", strconv.Itoa(i), 1)
	ur, err := url.Parse(u)
	return *ur, err
}

func fetchComic(i int) (c Comic, next bool, e error) {
	fmt.Printf("Fetching comic\t%5d\n", i)
	u, err := getComicURL(i)
	if err != nil {
		return c, true, err
	}
	resp, err := http.Get(u.String())
	if err != nil {
		return c, true, err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Comic not Found\t%5d\n", i)
		return c, false, nil
	}
	data := make([]byte, 0)
	dst := bytes.NewBuffer(data)
	if _, err := io.Copy(dst, resp.Body); err != nil {
		return c, true, err
	}
	if err = json.Unmarshal(dst.Bytes(), &c); err != nil {
		return c, true, err
	}
	return c, true, nil
}

// max -1 = no limit
func fetchFromComic(start int, max int) (comics []Comic, err error) {
	fails := 0
	next := true
	c := Comic{}
	count := 0
	for i := start; fails < 10 && (max == -1 || max > count); i++ {
		c, next, err = fetchComic(i)
		if err != nil {
			return nil, err
		}
		if !next {
			fails++
			continue
		}
		fails = 0
		comics = append(comics, c)
		count++
	}
	fmt.Printf("Total fetched: %d\n", len(comics))
	return comics, nil
}

func readComics(r io.Reader) (comics []Comic, err error) {
	data, err := ioutil.ReadAll(r)
	if len(data) == 0 {
		return
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &comics)
	return
}

func readComic(r io.Reader, i int) (comic Comic, err error) {
	comics, err := readComics(r)
	if err != nil {
		return
	}
	for _, c := range comics {
		if c.Num == i {
			return c, nil
		}
	}
	return comic, ErrComicNotFound
}

func writeComics(w io.Writer, comics []Comic) error {
	writer := bufio.NewWriter(w)
	data, err := json.MarshalIndent(comics, "", "  ")
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Flush()
	return err
}

func updateComics(r io.Reader, w io.Writer, max int) (err error) {
	comics, err := readComics(r)
	i := 0
	if err != nil {
		return err
	}
	if len(comics) > 0 {
		i = comics[len(comics)-1].Num
	}
	newComics, err := fetchFromComic(i+1, max)
	if err != nil {
		return err
	}
	comics = append(comics, newComics...)
	err = writeComics(w, comics)
	return
}
