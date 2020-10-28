package main

import (
	"io"
	"net/http"
)

func main() {

	var home homepage
	var about aboutpage

	mux := http.NewServeMux()
	mux.Handle("/homepage", home)
	mux.Handle("/aboutpage", about)

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
