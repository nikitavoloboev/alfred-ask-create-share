package main

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var usage = `alfred-ask-create-share [search|check] [<query>]

Open web submissions from Alfred.

Usage:
	alfred-ask-create-share search [<query>]
    alfred-ask-create-share check
    alfred-ask-create-share -h

Options:
    -h, --help    Show this message and exit.
`

var (
	// Icons
	iconAvailable = &aw.Icon{Value: "icons/update.png"}
	redditIcon    = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon    = &aw.Icon{Value: "icons/github.png"}
	forumsIcon    = &aw.Icon{Value: "icons/forums.png"}
	stackIcon     = &aw.Icon{Value: "icons/stack.png"}
	docIcon       = &aw.Icon{Value: "icons/doc.png"}

	repo = "nikitavoloboev/alfred-ask-create-share"
	wf   *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))
}

func run() {
	// Pass wf.Args() to docopt because our update logic relies on
	// AwGo's magic actions.
	args, _ := docopt.Parse(usage, wf.Args(), true, wf.Version(), false, true)

	// Alternate action: get available releases from remote
	if args["check"] != false {
		wf.TextErrors = true
		log.Println("Checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	// Script filter
	var query string
	if args["<query>"] != nil {
		query = args["<query>"].(string)
	}

	log.Printf("query=%s", query)

	// Call self with "check" command if an update is due and a
	// check job isn't already running.
	if wf.UpdateCheckDue() && !aw.IsRunning(updateJobName) {
		log.Println("Running update check in background...")
		cmd := exec.Command("./alfred-ask-create-share", "check")
		if err := aw.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if query == "" { // Only show update status if query is empty
		// Send update status to Alfred
		if wf.UpdateAvailable() {
			wf.NewItem("Update Available!").
				Subtitle("↩ to install").
				Autocomplete("workflow:update").
				Valid(false).
				Icon(iconAvailable)
		}
	}

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
}

// parseCSV parses CSV for links and arguments
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

func main() {
	wf.Run(run)
}
