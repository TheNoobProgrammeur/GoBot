package request

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Lang     string `json:"lang"`
	ID       int    `json:"id"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
	} `json:"flags"`
	Error bool `json:"error"`
}

func GetJock(route string) (string,error) {

	url := "https://sv443.net/jokeapi/v2/joke/"

	if route != "" {
		url += route
	}

	resp, err := http.Get(url)

	if err != nil {
		return "error", err
	}

	defer resp.Body.Close()
	r := Response{}

	json.NewDecoder(resp.Body).Decode(&r)

	if r.Type == "twopart" {
		jokeTwo := "- " + r.Setup + "\n" + "- " + r.Delivery
		return jokeTwo, nil
	}

	return r.Joke, nil
}
