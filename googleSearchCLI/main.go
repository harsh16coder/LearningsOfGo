package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <Search Query>")
		return
	}
	queryKeyWord := strings.Join(os.Args[1:], " ")
	bodystring, err := Search(queryKeyWord)
	fmt.Println(bodystring)
	if err != nil {
		fmt.Println("Program stopped with error: ", err)
		os.Exit(1)
	}
	if strings.Contains(bodystring, "getstream.io") {
		fmt.Println("Query does have getstream.io in response body")
		os.Exit(0)
	} else {
		fmt.Println("Query does not have getstream.io in response")
		os.Exit(1)
	}
}

func Search(queryKeyWord string) (string, error) {
	searchQuery := url.QueryEscape(queryKeyWord)
	query := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", searchQuery)
	// req, err := http.NewRequest("GET", query, nil)
	// req.Header.Set("User-Agent", "Mozilla/5.0")
	// res, err := http.DefaultClient.Do(req)
	res, err := http.Get(query)
	if err != nil {
		log.Printf("%s error occured while processing get request", err)
	}
	defer res.Body.Close()
	bodybytes, err := io.ReadAll(res.Body)
	return string(bodybytes), err
}
