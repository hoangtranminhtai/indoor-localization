package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
//}

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	tmplt := template.New("index.html")       //create a new template with some name
	tmplt, _ = tmplt.ParseFiles("index.html") //parse some content and generate a template, which is an internal representation

	p := Student{Id: 1, Name: "Aisha"} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func main() {

	fmt.Println("Server Starting")
	//http.HandleFunc("/", SimpleIndexHandler)
	http.HandleFunc("/index", HttpFileHandler)

	http.ListenAndServe(":8080", nil)
}