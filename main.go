package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type storyArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {

	var byteStory map[string]json.RawMessage
	parsedStory := map[string]storyArc{}

	data, err := ioutil.ReadFile("story.json")
	if err != nil {
		log.Fatalln("unable to read file: ", err)
	}

	err = json.Unmarshal(data, &byteStory)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range byteStory {
		arc := storyArc{}
		err := json.Unmarshal(v, &arc)
		if err != nil {
			fmt.Println(err)
			break
		}
		parsedStory[k] = arc
	}

	for k, v := range parsedStory {
		fmt.Println(k, v, "\n", "\n", "\n", "")
	}

}
