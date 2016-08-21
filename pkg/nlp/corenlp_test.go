package nlp

import "testing"

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
