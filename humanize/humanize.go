package main

import (
	"log"
)

func main() {
	a := map[string]A{
		"a": {
			Url: "a",
		},
		"b": {
			Url: "b",
		},
		"c": {
			Url: "c",
		},
	}
	log.Printf("v:%#+v\n", a["a"].Url)
	log.Printf("v:%#+v\n", a["b"].Url)
	log.Printf("v:%#+v\n", a["d"].Url)
}

type A struct {
	Url string `json:"url"`
}
