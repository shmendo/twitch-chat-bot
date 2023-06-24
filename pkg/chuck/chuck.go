package chuck

import (
	"encoding/json"
	"io"
	"net/http"
)

type ChuckFact struct {
	Id    string `json:id`
	Url   string `json:url`
	Value string `json:value`
}

func RandomChuckFact() (string, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var fact ChuckFact
	err = json.Unmarshal(body, &fact)
	if err != nil {
		return "", err
	}

	return fact.Value, err
}
