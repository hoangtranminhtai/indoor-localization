package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

//func SimpleIndexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
//}

type Position struct {
	Lat float32
	Lng float32 //exported field since it begins with a capital letter
}

type Error struct {
	Err string //exported field since it begins with a capital letter
}

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {

	temp := template.New("index.html")      //create a new template with some name
	temp, _ = temp.ParseFiles("index.html") //parse some content and generate a template, which is an internal representation

	rawQuery := request.URL.RawQuery

	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		http.ServeFile(response, request, "error.html")
		return
	}

	_, isHasLat := params["lat"]
	if !isHasLat {
		http.ServeFile(response, request, "error.html")
		return
	}

	_, isHasLng := params["lng"]
	if !isHasLng {
		http.ServeFile(response, request, "error.html")
		return
	}

	lat, err := strconv.Atoi(params["lat"][0])
	if err != nil {
		http.ServeFile(response, request, "error.html")
		return
	}

	lng, err := strconv.Atoi(params["lng"][0])
	if err != nil {
		http.ServeFile(response, request, "error.html")
		return
	}

	p := Position{float32(lat), float32(lng)} //define an instance with required field

	temp.Execute(response, p) //merge template ‘t’ with content of ‘p’
}

func main() {

	fmt.Println("Server listening at port 8080...................")
	http.HandleFunc("/localization", HttpFileHandler)

	http.ListenAndServe(":8080", nil)
}
