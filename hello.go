// some code corresponding to CH1 of "The Go Programming Language"
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// func main() {
// 	// s1()
// 	// s2()
// 	// s5()
// 	// s6()
// 	s7()
// }

// echos given name, and prints each call arg with index
func s1() {
	// go run hello.go <name>
	n := os.Args[1]
	fmt.Println("Hello, " + n)
	for i, v := range os.Args {
		fmt.Printf("%d, %s\n", i, v)
	}
}

// reads either args or file, and counts duplicate instances, displays counts >1
func s2() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			// inc el count
			if input.Text() == "" {
				break
			}
			counts[input.Text()]++
		}
	} else {
		// just read one file
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", os.Args[1], err)
		}
		for _, s := range strings.Split(string(data), "\n") {
			counts[s]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

// calls specified URL, copies response body to Stdout, displays size and status of response
func s5() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	r, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tried to fetch %s, err %v\n", url, err)
		os.Exit(1)
	}
	b, err := io.Copy(os.Stdout, r.Body)
	r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tried to read %s, err %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%d bytes copied\nstatus code: %d\n", b, r.StatusCode)
}

// goroutine implementation of fetch, that only prints call latency
func s6() {
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

// fetches URL and returns latency to ch
func fetch(url string, ch chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	n, err := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, n, url)
}

/*
* server impls
 */
var mu sync.Mutex
var count int

// sets up handler for routes
func s7() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/read", reader)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// default handler - echoes back path, counts calls
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// returns call count for default route
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count-> %d\n", count)
	mu.Unlock()
}

// reads in target query param name, gets value, displays
// then returns count from 1->cnt query param, inclusive
func reader(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	tgt := r.Form["tgt"][0]
	fmt.Fprintf(w, "tgt -> %s\n", r.Form[tgt])
	cnt := r.Form["cnt"][0]
	i, err := strconv.Atoi(cnt)
	if err != nil {
		log.Println("error in Atoi")
	}
	for j := 1; j <= i; j++ {
		fmt.Fprintf(w, "cnt -> %d\n", j)
	}
}
