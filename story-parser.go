package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type storyArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func parseStory(filename string) map[string]storyArc {
	var byteStory map[string]json.RawMessage
	parsedStory := map[string]storyArc{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("unable to read file: ", err)
	}

	err = json.Unmarshal(data, &byteStory)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range byteStory {
		arc := new(storyArc)
		err := json.Unmarshal(v, arc)
		if err != nil {
			fmt.Println(err)
			break
		}
		parsedStory[k] = *arc
	}

	return parsedStory
}

func muxGenerator(story map[string]storyArc) *http.ServeMux {
	mux := http.NewServeMux()
	for k, v := range story {
		s := v
		if k == "intro" {
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				tpl.ExecuteTemplate(w, "index.gohtml", s)
			})
			continue
		}
		mux.HandleFunc(fmt.Sprintf("/%v", k), func(w http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(w, "index.gohtml", s)
		})
	}
	return mux
}
