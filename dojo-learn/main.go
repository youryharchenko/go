package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8087", http.FileServer(http.Dir(".")))
	if err != nil {
		log.Fatalln(err)
	}

}
