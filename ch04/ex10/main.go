// Modify issues to report the results in age categories, say less than a month old,
// less than a year old, and more than a year old.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	oldYear = iota
	oldMonth
	newMonth
)

func main() {
	issuesReport()
}

func issuesReport() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	iss := make(map[int][]github.Issue)
	iss[oldYear] = make([]github.Issue, 0)
	iss[oldMonth] = make([]github.Issue, 0)
	iss[newMonth] = make([]github.Issue, 0)
	fmt.Printf("%d issues:\n\n", result.TotalCount)
	for _, item := range result.Items {
		t := item.CreatedAt
		plusYear := time.Date(
			t.Year()+1,
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			t.Nanosecond(),
			t.Location())
		plusMonth := time.Date(
			t.Year(),
			t.Month()+1,
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			t.Nanosecond(),
			t.Location())
		switch {
		case time.Now().Unix() > plusYear.Unix():
			iss[oldYear] = append(iss[oldYear], *item)
		case time.Now().Unix() > plusMonth.Unix():
			iss[oldMonth] = append(iss[oldMonth], *item)
		default:
			iss[newMonth] = append(iss[newMonth], *item)
		}
		// fmt.Printf("#%-5d %9.9s %.55s\n",
		// 	item.Number, item.User.Login, item.Title)
	}
	for key := range iss {
		switch key {
		case oldYear:
			fmt.Printf("More than a year old:\n\n")
		case oldMonth:
			fmt.Printf("Less than a year old:\n\n")
		case newMonth:
			fmt.Printf("Less than a month old:\n\n")
		}
		for _, i := range iss[key] {
			fmt.Printf("#%-5d %s %9.9s %.55s\n",
				i.Number, i.CreatedAt, i.User.Login, i.Title)
		}
		fmt.Println()
	}
}
