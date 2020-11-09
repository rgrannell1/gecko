// Implements an `echo` clone with support for ansi colour-highlighting.
package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/logrusorgru/aurora"
)

// A map of CLI flags to non-colour transformations, like underlines
// and bold.
var formatFlags = map[string]func(str interface{}) aurora.Value{
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

// A map of CLI flags to text-colour highlights.
var foregroundFlags = map[string]func(str interface{}) aurora.Value{
	"--black":   aurora.Black,
	"--blue":    aurora.Blue,
	"--cyan":    aurora.Cyan,
	"--green":   aurora.Green,
	"--magenta": aurora.Magenta,
	"--red":     aurora.Red,
	"--white":   aurora.White,
	"--yellow":  aurora.Yellow,
}

// A map of CLI flags to text background-colour highlights.
var backgroundFlags = map[string]func(str interface{}) aurora.Value{
	"--bg-black":   aurora.BgBlack,
	"--bg-blue":    aurora.BgBlue,
	"--bg-cyan":    aurora.BgCyan,
	"--bg-green":   aurora.BgGreen,
	"--bg-magenta": aurora.BgMagenta,
	"--bg-red":     aurora.BgRed,
	"--bg-white":   aurora.BgWhite,
	"--bg-yellow":  aurora.BgYellow,
}

// Given a string to highlight, the CLI arguments that were passed to this program
// and a map of flag-names to string transformations. Applies all requested transformations
// (e.g bold + make read) to the string, and returns the serially updated string.
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

// Given a string to highlight, the CLI arguments that were passed to this program.
// Applies all requested transformations
// (e.g bold + make read) to the string, and returns the serially updated string.
func highlightInput(str interface{}, opts docopt.Opts) (highlighted interface{}, err error) {
	out := str

	out = highlightFlags(out, opts, foregroundFlags)
	out = highlightFlags(out, opts, backgroundFlags)
	out = highlightFlags(out, opts, formatFlags)

	return out, nil
}

// Run the Docopt CLI, and provided the parsed CLI arguments to the highlighting function. Finally,
// display the highlighted string with or without a terminating newline.
func main() {
	usage := `Gecko - display a line of text, with colours

Usage:
	gecko <string> [--bold] [--faint] [--doubly-underline] [--fraktur] [--italic] [--underline] [--slow-blink] [--rapid-blink] [--reverse-video] [--conceal] [--crossed-out] [--framed] [--encircled] [--overlined] [--black] [--red] [--green] [--yellow] [--blue] [--magenta] [--cyan] [--white] [--bg-black] [--bg-red] [--bg-green] [--bg-yellow] [--bg-blue] [--bg-magenta] [--bg-cyan] [--bg-white] [-n]

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
	// cast to string.
	input := fmt.Sprintf("%v", opts["<string>"])

	highlighted, err := highlightInput(input, opts)

	// shoulda newline be added at the end?
	isPresent, _ := opts.Bool("-n")

	if isPresent {
		fmt.Print(highlighted)
	} else {
		fmt.Println(highlighted)
	}

	// log to stderr and exit with status one.
	if err != nil {
		log.Fatal(err)
	}
}
