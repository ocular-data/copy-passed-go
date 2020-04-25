package fireboardtools

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func SetClipboard(token string, data string) {
	formData := map[string]string{"id": token, "method": "post", "data": data}

	jsonValue, _ := json.Marshal(formData)

	resp, err := http.Post("https://us-central1-copy-passed.cloudfunctions.net/access",
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		//return getToken()
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		panic(errors.New("Bad Token"))
	} else if resp.StatusCode != 200 {
		panic(errors.New(resp.Status))
	}
}
