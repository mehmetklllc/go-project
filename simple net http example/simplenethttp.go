package main

import (
	"net/http"
)



func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}


func aboutHandler(writer http.ResponseWriter, request *http.Request) {

	data :=  request.URL.Path[1:]
	writer.Write([]byte("hello about path! " + data))
	writer.WriteHeader(http.StatusOK)


}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello root path!"))
	writer.WriteHeader(http.StatusOK)
}
