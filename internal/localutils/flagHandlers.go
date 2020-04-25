package localutils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
)

func ReissueKey() {
	token := RetrieveToken()
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			if err.Error() == "Bad Token" {
				RetrieveToken(true)
			}
		}
	}()
	fireboardtools.RevokeToken(token)
	RetrieveToken(true)
}

func CopyLocal() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	Copy(text)
}

func ToLocalPaste() {
	defer GenerateNewTokenOnError(Paste)
	token := RetrieveToken()
	text := fireboardtools.GetClipboard(token).Last
	if err := clipboard.WriteAll(text); err != nil {
		fmt.Println("There was a problem copying to the computers local clipboard")
	}
}
