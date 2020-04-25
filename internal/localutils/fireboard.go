package localutils

import (
	"fmt"
	"github.com/ocular-data/copy-passed-go/pkg/fireboardtools"
	"io/ioutil"
	"os"
	"runtime"
)

func GenerateNewTokenOnError(onDone func()) {
	if r := recover(); r != nil {
		err := r.(error)
		if err.Error() == "Bad Token" {
			RetrieveToken(true)
			onDone()
		}
	}
}

func RetrieveToken(newKey ...bool) string {
	var token, tempFile string
	switch uos := runtime.GOOS; uos {
	case "darwin", "linux", "freebsd", "openbsd", "android", "netbsd", "solaris", "plan9", "nacl":
		tempFile = "/tmp/.CopyPasteFireToken"
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
