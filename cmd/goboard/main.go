package main

import (
	"errors"
	"github.com/atotto/clipboard"
	"github.com/bobziuchkovski/writ"
	"github.com/ocular-data/copy-passed-go/internal/localutils"
	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
	"io/ioutil"
	"os"
)

type CopyPassed struct {
	HelpFlag  bool `flag:"h, help" description:"Display this help message and exit"`
	Clipboard bool `flag:"c, clipboard" description:"Copy from the computers clipboard instead of stdin"`
	Revoke    bool `flag:"r, re-authenticate" description:"Revoke token registered to computer and register new one"`
}

func main() {
	copyPassed := &CopyPassed{}
	cmd := writ.New("greeter", copyPassed)
	cmd.Help.Usage = "Usage: [PIPE IN] goboard [OPTION]"
	cmd.Help.Header = "Copy Paste across multiple interfaces"
	cmd.Help.Footer = "By default, goboard copys from stdin and pastes to stdout. Use the -c option to override."

	_, positional, err := cmd.Decode(os.Args[1:])
	if err != nil || copyPassed.HelpFlag {
		cmd.ExitHelp(err)
	}

	if len(positional) > 0 {
		cmd.ExitHelp(errors.New("goboard does not accept positional arguments"))
	}

	if copyPassed.Clipboard {
		text, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}
		localutils.Copy(text)
		return
	}

	if copyPassed.Revoke {
		token := localutils.RetrieveToken()
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				if err.Error() == "Bad Token" {
					localutils.RetrieveToken(true)
				}
			}
		}()
		fireboardtools.RevokeToken(token)
		localutils.RetrieveToken(true)
		return
	}

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		out, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		localutils.Copy(string(out))
	} else {
		localutils.Paste()
	}
}
