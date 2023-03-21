package main

import (
	"fmt"
	"log"
	"net/http"
)

func Himym(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " ted mosby:Bazen verdiğimiz en iyi kararlar, en mantıksız olanlardir  ")
}

func HRobin() {

	http.HandleFunc("/", Himym)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	HRobin()

}
