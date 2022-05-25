package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println(" Endpoint Hit: All Articles Endpoint")
	responseWriter := json.NewEncoder(w)
	responseWriter.Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
