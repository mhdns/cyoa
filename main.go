package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	mux := muxGenerator(parseStory("story.json"))
	http.ListenAndServe(":5000", mux)
}
