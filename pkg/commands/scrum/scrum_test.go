package scrum

import (
	"scrumkin/pkg/messages"
	"testing"
)

var matchTests = []struct {
	text     string
	expected bool
}{
	{"today", true},
	{"yesterday", true},
	{"scrum", true},
	{" scrum", true},
	{"today ", true},
	{"SCRUM", true},
	{"foo", false},
	{"ping me", false},
}

func TestMatch(t *testing.T) {
	ping, _ := New()

	for _, testcase := range matchTests {
		msg := &messages.Message{
			Text: testcase.text,
		}
		if ping.Match(msg) != testcase.expected {
			t.Error("Expected scrum to match", testcase.text)
		}
	}
}

func TestName(t *testing.T) {
	ping, _ := New()

	expected := "scrum"
	actual := ping.Name()
	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestHelpSummary(t *testing.T) {
	ping, _ := New()

	if ping.HelpSummary() == "" {
		t.Error("Expected non-empty help summary")
	}
}
