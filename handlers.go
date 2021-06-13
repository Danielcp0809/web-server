package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, All it´s ok v2")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api it's ok")
}
