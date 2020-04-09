// program fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func fetch(uri string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(uri, "http://") {
		uri = "http://" + uri
	}
	resp, err := http.Get(uri)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	f, err := os.Create(url.QueryEscape(uri[7:]))
	if err != nil {
		ch <- err.Error()
	}
	nbytes, err := io.Copy(f, resp.Body)
	fmt.Fprintf(f, "\nStatus code: %s", resp.Status)
	resp.Body.Close() // don't leak resources

	if closeErr := f.Close(); err == nil { // example from chapter 5.8, gopl.io/ch5/fetch
		err = closeErr
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", uri, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, uri)
}
