// Find a web site that produces a large amount of data. Investigate caching by running
// fetchall twice in succession to see whether the reported time changes much. Do you get the
// same content each time? Modify fetchall to print its out put to a file so it can be examined.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	var nbytes int64
	times := make([]float64, 2)
	for i := 0; i < 2; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}

		nbytes, err = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close() // don't leak resources
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		times[i] = time.Since(start).Seconds()
	}
	ch <- fmt.Sprintf("1st lap:\t%.2fs  %7d  %s\n2nd lap:\t%.2fs  %7d  %s",
		times[0], nbytes, url, times[1], nbytes, url)
}
