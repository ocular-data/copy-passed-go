package localutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
)

//GenerateNewTokenOnError function for if there was an error then generate new tocken
func GenerateNewTokenOnError(onDone func()) {
	if r := recover(); r != nil {
		err := r.(error)
		if err.Error() == "Bad Token" {
			RetrieveToken(true)
			onDone()
		}
	}
}

//RetrieveToken gets the tocken stored in the temp folder on computer
func RetrieveToken(newKey ...bool) string {
	var token, tempFile string
	switch uos := runtime.GOOS; uos {
	case "darwin", "linux", "freebsd", "openbsd", "netbsd", "solaris", "plan9", "nacl":
		tempFile = "/tmp/.CopyPasteFireToken"
	case "android":
		tempFile = os.Getenv("HOME") + "/../usr/tmp/.CopyPasteFireToken"
	case "windows":
		tempFile = os.Getenv("temp") + "\\CopyPasteFireToken"
	default:
		fmt.Println("Not a recognised os generating new key...")
		return fireboardtools.GetToken()
	}
	if len(newKey) != 0 {
		token = fireboardtools.GetToken()
	} else if _, err := os.Stat(tempFile); !os.IsNotExist(err) {
		dat, err := ioutil.ReadFile(tempFile)
		if err != nil {
			token = fireboardtools.GetToken()
		} else {
			token = string(dat)
		}
	} else {
		token = fireboardtools.GetToken()
	}

	err := ioutil.WriteFile(tempFile, []byte(token), 0644)

	if err != nil {
		panic(err)
	}

	return token
}
