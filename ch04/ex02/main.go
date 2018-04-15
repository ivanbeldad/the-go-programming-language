// Write a program that prints the SHA256 hash of its standard input by default but
// supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

type sha int

const (
	s256 sha = 256
	s384 sha = 384
	s512 sha = 512
)

func main() {
	f := flag.Int("base", 256, "Set to 256, 384 or 512 to print SHA in that base.")
	flag.Parse()
	printSha(os.Stdout, generateSha(), sha(*f))
}

func printSha(w io.Writer, s []byte, base sha) {
	switch base {
	case s256:
		fmt.Fprintf(w, "%x\n", sha256.Sum256(s))
	case s384:
		fmt.Fprintf(w, "%x\n", sha512.Sum384(s))
	case s512:
		fmt.Fprintf(w, "%x\n", sha512.Sum512(s))
	default:
		fmt.Fprintf(w, "Invalid base: %d\n", base)
	}
}

func generateSha() []byte {
	s := make([]byte, 0, 256)
	sha256.New().Write(s)
	return s
}
