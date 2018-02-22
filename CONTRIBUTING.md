# Contributing
Thank you for wanting to make Alfred Ask Create Share better. You can:
- [Edit the CSV file](../../edit/master/workflow/submissions.csv) and add more web submissions to the workflow. It has a simple structure of argument, followed by comma and then what website is going to be opened.
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features with [Pull Requests](../../pulls/).

## Dependencies
The workflow is written in [Go language](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all the Alfred related things.

## Developing the workflow
If you want to add features and things to the workflow.

It is best to use [Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repository and run: `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then call from inside Alfred [Script Filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

I also wrote about my own process in making Alfred workflows [here](https://wiki.nikitavoloboev.xyz/macOS/apps/alfred/making-workflows.html).

## Submitting a Pull Request
Please go through [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

#### Thank you for taking the time to contribute! 💜