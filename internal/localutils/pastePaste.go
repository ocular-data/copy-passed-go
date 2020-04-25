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
		fmt.Println("There was a problem copying to the computers clipboard")
	}

	defer GenerateNewTokenOnError(func() {
		fmt.Printf("There was an error with you token...\nplease run the command again to copy\n")
	})
	token := RetrieveToken()
	fireboardtools.SetClipboard(token, data)
}
