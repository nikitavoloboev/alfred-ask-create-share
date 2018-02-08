package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// doSearch makes a search for web submissions and returns results to Alfred.
func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	links := parseCSV()

	for key, value := range links {
		if strings.Contains(key, "r: ") {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key).Icon(redditIcon)
		} else if strings.Contains(key, "s: ") {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key).Icon(stackIcon)
		} else if strings.Contains(key, "g: ") {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key).Icon(githubIcon)
		} else if strings.Contains(key, "f: ") {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key).Icon(forumsIcon)
		} else if strings.Contains(key, "d: ") {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key).Icon(docIcon)
		} else {
			wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key)
		}
	}

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
	return nil
}

// parseCSV parses CSV for links and arguments.
func parseCSV() map[string]string {
	var err error

	// Load values from file to a hash map
	f, err := os.Open("ask-create-share.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Holds user's search arguments and an appropriate search URL
	links := make(map[string]string)

	for _, record := range records {
		links[record[0]] = record[1]
	}

	return links

}
