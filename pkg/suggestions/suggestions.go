package suggestions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var ParseURLError = fmt.Errorf("cannot parse url")
var RequestError = fmt.Errorf("request error")
var UnmarshalError = fmt.Errorf("unmarshal error")

type Suggestions interface {
	GetSuggestions(text string, lon float32, lat float32, results int) ([]Suggestion, error)
}

type suggestionsProxy struct {
	url string
	key string
}

func NewSuggestions(apiKey, apiUrl string) Suggestions {
	return &suggestionsProxy{
		url: apiUrl,
		key: apiKey,
	}
}

func (p *suggestionsProxy) GetSuggestions(text string, lon float32, lat float32, results int) ([]Suggestion, error) {
	v := url.Values{}
	v.Add("apikey", p.key)
	v.Add("text", strings.TrimSpace(text))
	v.Add("ll", fmt.Sprintf("%f,%f", lon, lat))
	v.Add("results", strconv.Itoa(results))
	res, err := http.Get(fmt.Sprintf("%s?%s", p.url, v.Encode()))
	if err != nil {
		fmt.Println(err)
		return nil, RequestError
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var resStruct struct {
		Results []Suggestion `json:"results"`
	}
	if err := json.Unmarshal(body, &resStruct); err != nil {
		fmt.Println(string(body), err)
		return nil, UnmarshalError
	}

	return resStruct.Results, nil
}
