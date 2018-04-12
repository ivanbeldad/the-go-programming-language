// Modify the echo program to also print os.Args[0], the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var args string
	for _, arg := range os.Args {
		args += arg + " "
	}
	args = args[0 : len(args)-1]
	fmt.Println(args)
}
