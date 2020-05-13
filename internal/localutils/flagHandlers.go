package localutils

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
)

//ReissueKey revokes the current key(if aviable) and issue a new one
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

//CopyLocal copys from the computers clipboard to the global clipboard
func CopyLocal() {
	text, err := clipboard.ReadAll()
	if err != nil {
		// panic(err)
		fmt.Println("There was a problem copying from the computers local clipboard")
	}
	Copy(text)
}

//ToLocalPaste takes data from global clipboard and puts it in the computers local clipboard
func ToLocalPaste() {
	defer GenerateNewTokenOnError(Paste)
	token := RetrieveToken()
	text := fireboardtools.GetClipboard(token).Last
	if err := clipboard.WriteAll(text); err != nil {
		fmt.Println("There was a problem copying to the computers local clipboard")
	}
}
