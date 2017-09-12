package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	// kingpin and script
	app *kingpin.Application

	// app commands
	filterSubmissionsCmd *kingpin.CmdClause

	// script options
	query string

	// icons
	// redditIcon    = &aw.Icon{Value: "icons/reddit.png"}
	// docIcon       = &aw.Icon{Value: "icons/doc.png"}
	// gitubIcon     = &aw.Icon{Value: "icons/github.png"}
	// forumsIcon    = &aw.Icon{Value: "icons/forums.png"}
	// translateIcon = &aw.Icon{Value: "icons/translate.png"}
	// stackIcon     = &aw.Icon{Value: "icons/stack.png"}
	// iconAvailable = &aw.Icon{Value: "icons/update-available.png"}

	repo = "nikitavoloboev/alfred-ask-create-share"

	// workflow
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))

	app = kingpin.New("ask-create-share", "make web submissions from Alfred")
	// app.HelpFlag.Short('h')
	app.Version(wf.Version())

	filterSubmissionsCmd = app.Command("filter", "filters submissions")

	for _, cmd := range []*kingpin.CmdClause{filterSubmissionsCmd} {
		cmd.Flag("query", "search query").Short('q').StringVar(&query)
	}

	// list action commands
	app.DefaultEnvars()
}

// _actions

// fills Alfred with hash map values and shows keys
func filterResults(links map[string]string) {

	for key, value := range links {
		wf.NewItem(key).Valid(true).UID(key).Var("URL", value).Var("ARG", key)
	}

	wf.Filter(query)
	wf.SendFeedback()
}

func run() {
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

	// _arg parsing
	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	switch cmd {
	case filterSubmissionsCmd.FullCommand():
		filterResults(links)
	default:
		err = fmt.Errorf("unknown command: %s", cmd)
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	aw.Run(run)
}
