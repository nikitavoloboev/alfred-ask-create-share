package main

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var usage = `alfred-ask-create-share [search|check] [<query>]

Open web submissions from Alfred

Usage:
	alfred-ask-create-share search [<query>]
    alfred-ask-create-share check
    alfred-ask-create-share -h

Options:
    -h, --help    Show this message and exit.
`

var (
	iconAvailable = &aw.Icon{Value: "icons/update.png"}
	repo          = "nikitavoloboev/alfred-ask-create-share"
	wf            *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))
}

func run() {
	// log.Println(wf.Args())

	// Pass wf.Args() to docopt because our update logic relies on
	// AwGo's magic actions.
	args, _ := docopt.Parse(usage, wf.Args(), true, wf.Version(), false, true)

	// alternate action: get available releases from remote
	if args["check"] != false {
		wf.TextErrors = true
		log.Println("checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	// _script filter
	var query string
	if args["<query>"] != nil {
		query = args["<query>"].(string)
	}

	log.Printf("query=%s", query)

	// call self with "check" command if an update is due and a
	// check job isn't already running.
	if wf.UpdateCheckDue() && !aw.IsRunning(updateJobName) {
		log.Println("running update check in background...")
		cmd := exec.Command("./alfred-ask-create-share", "check")
		if err := aw.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("error starting update check: %s", err)
		}
	}

	if query == "" { // Only show update status if query is empty
		// Send update status to Alfred
		if wf.UpdateAvailable() {
			wf.NewItem("update available!").
				Subtitle("↩ to install").
				Autocomplete("workflow:update").
				Valid(false).
				Icon(iconAvailable)
		}
	}

	links := parseCSV()

	for key, value := range links {
		wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key)
	}

	// script filter results
	// for i := 1; i <= 20; i++ {
	// 	t := fmt.Sprintf("Item #%d", i)
	// 	wf.NewItem(t).
	// 		Icon(aw.IconFavourite).
	// 		Arg(t).
	// 		Valid(true)
	// }

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("no matching items", "try a different query?")
	wf.SendFeedback()
}

func parseCSV() map[string]string {
	var err error

	// load values from file to a hash map
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

	// holds user's search arguments and an appropriate search URL
	links := make(map[string]string)

	for _, record := range records {
		links[record[0]] = record[1]
	}

	return links

}

// fills Alfred with hash map values and shows keys
func filterResults(links map[string]string) {

	// wf.Filter(query)
	// wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
