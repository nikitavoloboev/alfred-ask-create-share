package main

import (
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}
	redditIcon      = &aw.Icon{Value: "icons/reddit.png"}
	githubIcon      = &aw.Icon{Value: "icons/github.png"}
	forumsIcon      = &aw.Icon{Value: "icons/forums.png"}
	stackIcon       = &aw.Icon{Value: "icons/stack.png"}
	docIcon         = &aw.Icon{Value: "icons/doc.png"}

	query string

	repo = "nikitavoloboev/alfred-ask-create-share"

	// Workflow stuff
	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	searchSubmissions()
}

func main() {
	wf.Run(run)
}
