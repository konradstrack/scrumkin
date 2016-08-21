package nlp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBuildURL(t *testing.T) {
	config := &CoreNLPConfig{
		BaseURL:    "http://corenlp:5000",
		Annotators: []string{"tokenize", "pos"},
	}

	expectedURL := `http://corenlp:5000/?properties=%7B%22annotators%22%3A%22tokenize%2Cpos%22%7D`
	actualURL, _ := buildURL(config)
	if actualURL != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, actualURL)
	}
}

func TestParse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(coreNLPHandler))
	defer s.Close()

	config := &CoreNLPConfig{
		BaseURL: s.URL,
	}

	client, _ := NewCoreNLPClient(config)
	text, _ := client.Parse("I like pie")
	if actual, expected := len(text.Sentences), 1; actual != expected {
		t.Errorf("Expected %d sentences, got %d", expected, actual)
	}

	tokens := text.Sentences[0].Tokens
	actualPosTags := make([]string, 3)
	for i, t := range tokens {
		actualPosTags[i] = t.POS
	}

	expectedPosTags := []string{"PRP", "VBP", "NN"}
	if !reflect.DeepEqual(actualPosTags, expectedPosTags) {
		t.Errorf("Expected %v, got %v", expectedPosTags, actualPosTags)
	}
}

var parserResponse = `
{
  "sentences": [
    {
      "index": 0,
      "tokens": [
        {
          "index": 1,
          "word": "I",
          "originalText": "I",
          "lemma": "I",
          "characterOffsetBegin": 0,
          "characterOffsetEnd": 1,
          "pos": "PRP",
          "ner": "O",
          "before": "",
          "after": " "
        },
        {
          "index": 2,
          "word": "like",
          "originalText": "like",
          "lemma": "like",
          "characterOffsetBegin": 2,
          "characterOffsetEnd": 6,
          "pos": "VBP",
          "ner": "O",
          "before": " ",
          "after": " "
        },
        {
          "index": 3,
          "word": "pie",
          "originalText": "pie",
          "lemma": "pie",
          "characterOffsetBegin": 7,
          "characterOffsetEnd": 10,
          "pos": "NN",
          "ner": "O",
          "before": " ",
          "after": ""
        }
      ]
    }
  ]
}
`

func coreNLPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, parserResponse)
}
