package main

import (
	"io"
	"net/http"
)

func main() {

	var i homepage
	var w aboutpage

	mux := http.NewServeMux()
	mux.Handle("/homepage", i)
	mux.Handle("/aboutpage", w)

	http.ListenAndServe(":8080",mux)

}


type homepage int

func (x homepage) ServeHTTP(res http.ResponseWriter,r *http.Request){
	io.WriteString(res , "Hi  HomePage !")
}

type aboutpage int

func (x aboutpage) ServeHTTP(res http.ResponseWriter,r *http.Request){
	io.WriteString(res , "Hi  AboutPage!")
}
