package nlp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	BaseURL    string
	Annotators []string
}

type Client interface {
	Query(string) (*Text, error)
}

type CoreNLPClient struct {
	config   *Config
	queryURL string
}

func NewCoreNLPClient(config *Config) Client {
	url := buildURL(config)

	log.Print(url)
	return &CoreNLPClient{
		config:   config,
		queryURL: url,
	}
}

func (c *CoreNLPClient) Query(text string) (*Text, error) {
	response, err := http.Post(c.queryURL, "text/plain", strings.NewReader(text))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("Received status %s, %s", response.Status, b)
	}

	var textObj Text
	json.NewDecoder(response.Body).Decode(&textObj)
	return &textObj, nil
}

type properties struct {
	Annotators string `json:"annotators"`
}

func buildURL(config *Config) string {
	p := properties{
		Annotators: strings.Join(config.Annotators, ","),
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	v := url.Values{}
	v.Set("properties", b.String())

	return fmt.Sprintf("%s/?%s", config.BaseURL, v.Encode())
}
