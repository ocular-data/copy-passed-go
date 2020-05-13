package fireboardtools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//RevokeToken revokes the given token
func RevokeToken(token string) {
	formData := map[string]string{"id": token, "revoke": "true"}

	jsonValue, _ := json.Marshal(formData)

	resp, err := http.Post("https://us-central1-copy-passed.cloudfunctions.net/authenticator",
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		//return getToken()
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Token Revoked")
	//if resp.StatusCode == 200 {
	//	fmt.Println("Token Revoked")
	//}
}
