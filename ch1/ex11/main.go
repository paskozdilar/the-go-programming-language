package main

// Exercise 1.11: Try fetchall with longer argument lists, such as samples from
// the top million web sites available at alexa.com. How does the program
// behave if a web site just doesn't respond?

// Exercise log (no code modification): The alexa.com list doesn't exist
// anymore. An alternative list is available on:
// https://www.domcop.com/top-10-million-websites
// I'm using 1000 websites from this list as arguments. They are also located
// in the file: `top1000websites.txt`
// When website just doesn't respond, one of the following errors are printed:
//	- net/http: TLS handshake timeout
//	- read tcp [...]: read: connection reset by peer
//	- dial tcp [...]: i/o timeout
//	- stream error: stream ID 1; INTERNAL_ERROR; received from peer
// After many errors, the program started to hang as well. This is probably
// what happens when the TCP connection stays open, but server does not send
// any data.

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
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", secs, nbytes, url)
}
