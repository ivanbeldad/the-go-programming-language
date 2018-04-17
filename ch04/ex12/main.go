package main

import (
	"log"
)

func main() {
	err := execCommandAction()
	if err != nil {
		log.Fatal(err)
	}
}
