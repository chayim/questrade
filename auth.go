package questrade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	APIServer    string `json:"api_server"`
	Expiry       int32  `json:"expires_in"`
}

func RefreshToken(refresh_token string) (Auth, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://login.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token=%s", refresh_token)
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Auth{}, fmt.Errorf("Questrade token refresh failure.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Auth{}, err
	}

	qauth := Auth{}
	err = json.Unmarshal(body, &qauth)
	if err != nil {
		return Auth{}, err
	}

	return qauth, nil

}
