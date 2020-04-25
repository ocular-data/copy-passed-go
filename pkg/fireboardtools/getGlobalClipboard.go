package fireboardtools

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type History struct {
	Hist1 string `json:"hist_1,omitempty"`
	Hist2 string `json:"hist_2,omitempty"`
	Hist3 string `json:"hist_3,omitempty"`
	Hist4 string `json:"hist_4,omitempty"`
	Hist5 string `json:"hist_5,omitempty"`
	Hist6 string `json:"hist_6,omitempty"`
	Hist7 string `json:"hist_7,omitempty"`
	Hist8 string `json:"hist_8,omitempty"`
	Hist9 string `json:"hist_9,omitempty"`
}

type Clipboard struct {
	Last    string  `json:"last"`
	History History `json:"history,omitempty"`
}

func GetClipboard(token string) Clipboard {
	formData := map[string]string{"id": token, "method": "get"}

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
	} else if resp.StatusCode != 202 {
		panic(errors.New(resp.Status))
	}

	var data Clipboard
	_ = json.NewDecoder(resp.Body).Decode(&data)

	return data
}
