package main

import (
	"fmt"
	"net/http"
)

type News struct {
	Time  int
	Title string
}

const PORT = 4000

func (news News) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, news.Time, " --- ", news.Title)
}

func main() {
	// your http.Handle calls here
	http.Handle("/news", News{2019, "a new Go server online"})
	fmt.Println("listen on ",PORT)
	http.ListenAndServe("localhost:4000", nil)
}
