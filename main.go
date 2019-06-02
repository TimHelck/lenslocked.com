package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "<h1>Welcome to my awesome site</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)    // default serve mux
	http.ListenAndServe(":4000", nil)
}
