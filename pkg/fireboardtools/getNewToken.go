package fireboardtools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ocular-data/copy-passed-go/pkg/qrcode"
)

//GetToken walk user through the process of generating a new token
func GetToken() string {
	rand.Seed(int64(time.Now().UnixNano()))
	identifier := strconv.Itoa(rand.Intn(1e9) + 1e8 - 1)
	fmt.Println("please go to")
	fmt.Println(fmt.Sprintf("https://copy-passed.web.app/VerifyID.html#%s", identifier))
	fmt.Println("or scan qrcode with app")
	qrcode.PrintCode(identifier)
	fmt.Println("to complete the signin")

	formData := map[string]string{"id": identifier}

	jsonValue, _ := json.Marshal(formData)

	resp, err := http.Post("https://us-central1-copy-passed.cloudfunctions.net/authenticator",
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		//return getToken()
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		fmt.Println(resp.Status)
		fmt.Println("Error try again")
		return GetToken()
	}
	fmt.Println("Completed")

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["id"].(string)
}
