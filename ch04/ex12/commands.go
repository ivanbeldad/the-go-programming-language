package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

const (
	templateFolder      = "templates/"
	commandsTemplate    = templateFolder + "commands.txt"
	comicDetailTemplate = templateFolder + "comicDetail.txt"
	comicListTemplate   = templateFolder + "comicList.txt"
)

type Info struct {
	Commands []Action
}

type Action struct {
	Name        string
	Description string
	Exec        Command
}

type Command interface {
	execute() error
}

type CommandHelp struct{}
type CommandUpdateComics struct{}
type CommandComicDetail struct{}
type CommandComicList struct{}

var update = Action{
	Name:        "update",
	Description: "Update comics database",
	Exec:        CommandUpdateComics{},
}
var detail = Action{
	Name:        "detail",
	Description: "Shows the info about N comic",
	Exec:        CommandComicDetail{},
}
var list = Action{
	Name:        "list",
	Description: "Shows short information about the last N comics",
	Exec:        CommandComicList{},
}
var help = Action{
	Name:        "help",
	Description: "Shows this help message",
	Exec:        CommandHelp{},
}
var cmds = map[string]Action{
	update.Name: update,
	detail.Name: detail,
	list.Name:   list,
	help.Name:   help,
}

func (c CommandUpdateComics) execute() (err error) {
	max := -1
	args := os.Args
	if len(args) > 2 {
		max, err = strconv.Atoi(args[2])
		if err != nil {
			return err
		}
	}
	fw, err := os.OpenFile(comicFile, os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	fr, err := os.OpenFile(comicFile, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	r := bufio.NewReader(fr)
	w := bufio.NewWriter(fw)
	err = updateComics(r, w, max)
	fr.Close()
	fw.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c CommandComicDetail) execute() (err error) {
	args := os.Args
	if len(args) <= 2 {
		fmt.Printf("Required N (index) of the comic")
		return nil
	}
	n, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	f, err := os.OpenFile(comicFile, os.O_RDONLY, 0755)
	if err != nil {
		return
	}
	r := bufio.NewReader(f)
	comic, err := readComic(r, n)
	f.Close()
	if err != nil {
		return
	}
	templ, err := template.ParseFiles(comicDetailTemplate)
	if err != nil {
		return
	}
	return templ.Execute(os.Stdout, comic)
}

func (c CommandComicList) execute() (err error) {
	args := os.Args
	if len(args) <= 2 {
		fmt.Printf("Required N (amount) of comics")
		return nil
	}
	n, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	f, err := os.OpenFile(comicFile, os.O_RDONLY, 0755)
	if err != nil {
		return
	}
	r := bufio.NewReader(f)
	comics, err := readComics(r)
	f.Close()
	if err != nil {
		return
	}
	if n >= len(comics) {
		n = len(comics)
	}
	if n < 0 {
		n = 0
	}
	templ, err := template.ParseFiles(comicListTemplate)
	if err != nil {
		return
	}
	return templ.Execute(os.Stdout, comics[len(comics)-n:])
}

func (c CommandHelp) execute() error {
	t, err := template.ParseFiles(commandsTemplate)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, cmds)
}

func execCommandAction() error {
	return cmds[getConsoleCommand()].Exec.execute()
}

func getConsoleCommand() string {
	args := os.Args
	if len(args) == 1 {
		return "help"
	}
	if _, ok := cmds[args[1]]; !ok {
		return "help"
	}
	return args[1]
}
