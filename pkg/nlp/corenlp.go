package nlp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client describes a interface of an NLP parser
type Client interface {
	Parse(string) (*Text, error)
}

// CoreNLPClient represents a client for the CoreNLP server
type CoreNLPClient struct {
	config   *CoreNLPConfig
	queryURL string
}

// CoreNLPConfig holds configuration for CoreNLPClient
type CoreNLPConfig struct {
	BaseURL    string
	Annotators []string
}

// NewCoreNLPClient builds and returns a new CoreNLP client
func NewCoreNLPClient(config *CoreNLPConfig) (Client, error) {
	url, err := buildURL(config)
	if err != nil {
		return nil, err
	}

	client := &CoreNLPClient{
		config:   config,
		queryURL: url,
	}

	return client, nil
}

// Parse queries CoreNLP server and returns text parsing results
func (c *CoreNLPClient) Parse(text string) (*Text, error) {
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

func buildURL(config *CoreNLPConfig) (string, error) {
	p := properties{
		Annotators: strings.Join(config.Annotators, ","),
	}
	encoded, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Set("properties", string(encoded[:]))

	return fmt.Sprintf("%s/?%s", config.BaseURL, v.Encode()), nil
}
