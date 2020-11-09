package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/logrusorgru/aurora"
)

var FORMAT_FLAGS = map[string]func(str interface{}) aurora.Value{
	"--bold":             aurora.Bold,
	"--conceal":          aurora.Conceal,
	"--crossed-out":      aurora.CrossedOut,
	"--doubly-underline": aurora.DoublyUnderline,
	"--encircled":        aurora.Encircled,
	"--faint":            aurora.Faint,
	"--fraktur":          aurora.Fraktur,
	"--framed":           aurora.Framed,
	"--italic":           aurora.Italic,
	"--overlined":        aurora.Overlined,
	"--rapid-blink":      aurora.RapidBlink,
	"--reverse-video":    aurora.Reverse,
	"--slow-blink":       aurora.SlowBlink,
	"--underline":        aurora.Underline,
}

var FOREGROUND_FLAGS = map[string]func(str interface{}) aurora.Value{
	"--black":   aurora.Black,
	"--blue":    aurora.Blue,
	"--cyan":    aurora.Cyan,
	"--green":   aurora.Green,
	"--magenta": aurora.Magenta,
	"--red":     aurora.Red,
	"--white":   aurora.White,
	"--yellow":  aurora.Yellow,
}

var BACKGROUND_FLAGS = map[string]func(str interface{}) aurora.Value{
	"--bg-black":   aurora.BgBlack,
	"--bg-blue":    aurora.BgBlue,
	"--bg-cyan":    aurora.BgCyan,
	"--bg-green":   aurora.BgGreen,
	"--bg-magenta": aurora.BgMagenta,
	"--bg-red":     aurora.BgRed,
	"--bg-white":   aurora.BgWhite,
	"--bg-yellow":  aurora.BgYellow,
}

func highlightFlags(str interface{}, opts docopt.Opts, flags map[string]func(str interface{}) aurora.Value) (highlighted interface{}) {
	out := str

	for key, ansi := range flags {
		isPresent, _ := opts.Bool(key)

		if isPresent {
			out = ansi(out)
		}
	}

	return out
}

func highlightInput(str interface{}, opts docopt.Opts) (highlighted interface{}, err error) {
	out := str

	out = highlightFlags(out, opts, FOREGROUND_FLAGS)
	out = highlightFlags(out, opts, BACKGROUND_FLAGS)
	out = highlightFlags(out, opts, FORMAT_FLAGS)

	return out, nil
}

func main() {
	usage := `Gecko - display a line of text, with colours

Usage:
	gecko <string> [--bold] [--faint] [--doubly-underline] [--fraktur] [--italic] [--underline] [--slow-blink] [--rapid-blink] [--reverse-video] [--conceal] [--crossed-out] [--framed] [--encircled] [--overlined] [--black] [--red] [--green] [--yellow] [--blue] [--magenta] [--cyan] [--white] [--bg-black] [--bg-red] [--bg-green] [--bg-yellow] [--bg-blue] [--bg-magenta] [--bg-cyan] [--bg-white]

Options:
	--bold
	--conceal
	--crossed-out
	--doubly-underline
	--encircled
	--faint
	--fraktur
	--framed
	--italic
	--overlined
	--rapid-blink
	--reverse-video
	--slow-blink
	--underline
	--black
	--blue
	--cyan
	--green
	--magenta
	--red
	--white
	--yellow
	--bg-black
	--bg-blue
	--bg-cyan
	--bg-green
	--bg-magenta
	--bg-red
	--bg-white
	--bg-yellow
	-n           do not output the trailing newline
	-h --help    show this document

Author:
	Róisín Grannell <r.grannell2@gmail.com>
	`

	opts, _ := docopt.ParseDoc(usage)
	// -- cast to string.
	input := fmt.Sprintf("%v", opts["<string>"])

	highlighted, err := highlightInput(input, opts)

	// -- shoulda newline be added at the end?
	isPresent, _ := opts.Bool("-n")

	if isPresent {
		fmt.Print(highlighted)
	} else {
		fmt.Println(highlighted)
	}

	// -- log to stderr and exit with status one.
	if err != nil {
		log.Fatal(err)
	}
}
