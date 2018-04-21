// Write a function to populate a mapping from element names—p, div, span, and
// so on—to the number of elements with that name in an HTML document tree.

package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"golang.org/x/net/html"
)

// NodeCount ...
type NodeCount map[string]int

// Fprint ...
func (nc NodeCount) Fprint(w io.Writer) (err error) {
	tpl, err := template.ParseFiles("nodeCountTpl.txt")
	if err != nil {
		return
	}
	return tpl.Execute(w, nc)
}

// Print ...
func (nc NodeCount) Print() error {
	return nc.Fprint(os.Stdout)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	nodeCount := NodeCount{}
	fill(&nodeCount, doc)
	nodeCount.Print()
}

func fill(nc *NodeCount, cn *html.Node) {
	if cn.Type == html.ElementNode {
		(*nc)[cn.Data]++
	}
	for next := cn.FirstChild; next != nil; next = next.NextSibling {
		fill(nc, next)
	}
}
