package questrade

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (qt Questrade) request(url string) ([]byte, error) {
	client := &http.Client{}
	complete_url := fmt.Sprintf("%s%s", qt.apiserver, url)
	req, _ := http.NewRequest("GET", complete_url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", qt.token))
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		return body, nil
	}

	req.Header.Del("Authorization")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", qt.token))
	resp, _ = client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
