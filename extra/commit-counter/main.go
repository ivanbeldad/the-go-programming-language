package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	reposURL       = "https://api.github.com/users/$$$/repos"
	githubUsername = "ivandelabeldad"
	githubToken    = ""
)

// Contributor ...
type Contributor struct {
	Login   string `json:"login"`
	Commits int    `json:"contributions"`
}

// Repository ...
type Repository struct {
	ContributorsURL string `json:"contributors_url"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("\nYou must specify the github username\n\n")
		os.Exit(1)
	}
	login := os.Args[1]

	user := Contributor{Login: login}
	repos := getRepos(user)
	contribs := make([]Contributor, 0)
	for _, repo := range repos {
		contribs = append(contribs, getContributors(repo)...)
	}
	contribs = filterContributors(contribs, user)
	fmt.Printf("\nUser %q have a total of %d commits in %d repositories\n",
		user.Login, sumCommits(contribs), len(contribs))
}

func getRepos(user Contributor) []Repository {
	u := strings.Replace(reposURL, "$$$", user.Login, 1)

	fmt.Printf("Fetching %s\n", u)

	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(githubUsername, githubToken))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	repos := make([]Repository, 0)
	err = json.Unmarshal(content, &repos)
	if err != nil {
		log.Fatal(err)
	}
	return repos
}

func getContributors(repo Repository) []Contributor {
	fmt.Printf("Fetching %s\n", repo.ContributorsURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", repo.ContributorsURL, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(githubUsername, githubToken))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	contribs := make([]Contributor, 0)
	err = json.Unmarshal(content, &contribs)
	if err != nil {
		log.Fatal(err)
	}
	return contribs
}

func filterContributors(contribs []Contributor, c Contributor) []Contributor {
	for i := 0; i < len(contribs); i++ {
		if contribs[i].Login != c.Login {
			contribs = append(contribs[:i], contribs[i+1:]...)
			i--
		}
	}
	return contribs
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func sumCommits(contribs []Contributor) (sum int) {
	for _, c := range contribs {
		sum += c.Commits
	}
	return
}
