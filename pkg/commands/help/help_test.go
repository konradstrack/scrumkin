package help

import (
	"scrumkin/pkg/messages"
	"testing"
)

var matchTests = []struct {
	text     string
	expected bool
}{
	{"help", true},
	{"Help", true},
	{"Help ", true},
	{" Help", true},
	{"HELP", true},
	{"foo", false},
	{"help me", false},
}

func TestMatch(t *testing.T) {
	help, _ := New()

	for _, testcase := range matchTests {
		msg := &messages.Message{
			Text: testcase.text,
		}
		if help.Match(msg) != testcase.expected {
			t.Error("Expected help to match", testcase.text)
		}
	}
}
