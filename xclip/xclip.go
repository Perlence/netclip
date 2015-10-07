package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/Perlence/netclip/config"
)

const (
	outputUsage    = "prints the selection to standard out (generally for piping to a file or program)"
	inputUsage     = "read text into X selection from standard input or files (default)"
	filterUsage    = "when xclip is invoked in the in mode with output level set to silent (the defaults), the filter option will cause xclip to print the text piped to standard in back to standard out unmodified"
	loopsUsage     = "number of X selection requests (pastes into X applications) to wait for before exiting, with a value of 0 (default) causing xclip to wait for an unlimited number of requests until another application (possibly another invocation of xclip) takes ownership of the selection"
	displayUsage   = `X display to use (e.g. "localhost:0"), xclip defaults to the value in $DISPLAY if this option is omitted`
	selectionUsage = `specify which X selection to use, options are "primary" to use XA_PRIMARY (default), "secondary" for XA_SECONDARY or "clipboard" for XA_CLIPBOARD`
	versionUsage   = "show version information"
	silentUsage    = "forks into the background to wait for requests, no informational output, errors only (default)"
	quiteUsage     = "show informational messages on the terminal and run in the foreground"
	verboseUsage   = "provide a running commentary of what xclip is doing"
)

var (
	output    bool
	input     bool
	filter    bool
	loops     int
	display   string
	selection string
	version   string
	silent    bool
	quite     bool
	verbose   bool
)

func init() {
	flag.BoolVar(&output, "o", false, "")
	flag.BoolVar(&output, "out", false, outputUsage)
	flag.BoolVar(&input, "i", true, "")
	flag.BoolVar(&input, "in", true, inputUsage)
	flag.BoolVar(&filter, "f", false, "")
	flag.BoolVar(&filter, "filter", false, filterUsage)
	flag.IntVar(&loops, "l", 0, "")
	flag.IntVar(&loops, "loops", 0, loopsUsage)
	flag.StringVar(&display, "d", "", "")
	flag.StringVar(&display, "display", "", displayUsage)
	flag.StringVar(&selection, "sel", "primary", "")
	flag.StringVar(&selection, "selection", "primary", selectionUsage)
	flag.StringVar(&version, "v", "", "")
	flag.StringVar(&version, "version", "", versionUsage)
	flag.BoolVar(&silent, "silent", false, silentUsage)
	flag.BoolVar(&quite, "quiet", false, quiteUsage)
	flag.BoolVar(&verbose, "verbose", false, verboseUsage)
}

func main() {
	flag.Parse()
	args := flag.Args()

	conn, err := net.Dial("tcp", config.Addr)
	fatalOnError(err)
	defer conn.Close()

	var tee io.Reader
	if output {
		if wc, ok := conn.(WriteCloser); ok {
			err := wc.CloseWrite()
			fatalOnError(err)
		} else {
			log.Fatalln("conn does not implement WriteCloser interface")
		}

		tee = io.TeeReader(conn, os.Stdout)
	} else {
		var r io.Reader
		if len(args) == 0 {
			r = os.Stdin
		} else {
			var readers []io.Reader
			for _, filename := range args {
				file, err := os.Open(filename)
				fatalOnError(err)
				defer file.Close()
				readers = append(readers, file)
			}
			r = io.MultiReader(readers...)
		}

		tee = io.TeeReader(r, conn)
	}
	_, err = ioutil.ReadAll(tee)
	fatalOnError(err)
}

type WriteCloser interface {
	CloseWrite() error
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
