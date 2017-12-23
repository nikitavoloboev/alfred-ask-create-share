# Alfred Ask Create Share [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> [Alfred](https://www.alfredapp.com/) workflow for creating various Web Submissions (Stack exchange, Reddit and more)

<img src="media/demo.gif" width="400" alt="img">

You can filter submissions by using various prefixes.

|  Prefix |  Description |
|---|---|
| s:  |  Filter for asking questions on any one of stack exchange sites |
|  r: | Filter for creating new threads on various reddit subreddits  |
|  f: |  Filter for asking questions on various forums like [Alfred Forum](https://www.alfredforum.com/) |
|  g: |  Currently allows you to quickly create new repository or new gist |
|  w: | All other websites like creating new hacker news thread or new codepen|
| d:  |  Create google docs, sheets, slide or form |

<img src="https://i.imgur.com/hZe2AUY.png" width="400" alt="img">

## Installing
Download the workflow from [GitHub releases](https://github.com/nikitavoloboev/alfred-ask-create-share/releases/latest).

## Contributing
[Suggestions](https://github.com/nikitavoloboev/alfred-ask-create-share/issues) and pull requests are highly encouraged!

You can [edit the CSV file](https://github.com/nikitavoloboev/alfred-ask-create-share/edit/master/workflow/ask-create-share.csv) and add more web submissions to add to the workflow.

It has a simple structure of argument, followed by comma and then what website is going to be opened.

## Developing
If you want to add features and things to the workflow. It is best to use [this Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then run `alfred link` inside this repo you cloned. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make your changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then use from inside Alfred [script filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

## Credits
The workflow is built using [Deanishe](https://github.com/deanishe)'s amazing [AwGo](https://github.com/deanishe/awgo) library.

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look into [other repositories](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) I shared.

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)