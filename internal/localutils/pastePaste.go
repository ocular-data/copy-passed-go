package localutils

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
)

func Paste() {
	defer GenerateNewTokenOnError(Paste)
	token := RetrieveToken()
	text := fireboardtools.GetClipboard(token).Last
	fmt.Print(text)
}

func Copy(data string) {
	if err := clipboard.WriteAll(data); err != nil {
		fmt.Println("There was a problem copying to the computers local clipboard")
	}

	defer GenerateNewTokenOnError(func() {
		Copy(data)
	})
	token := RetrieveToken()
	fireboardtools.SetClipboard(token, data)
}
