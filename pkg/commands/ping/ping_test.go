package ping

import (
	"reflect"
	"scrumkin/pkg/messages"
	"testing"
)

var matchTests = []struct {
	text     string
	expected bool
}{
	{"ping", true},
	{"Ping", true},
	{"Ping ", true},
	{" Ping", true},
	{"PING", true},
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
			t.Error("Expected ping to match", testcase.text)
		}
	}
}

func TestProcess(t *testing.T) {
	ping, _ := New()
	msg := &messages.Message{
		Text: "ping",
	}

	expected := &messages.Response{"pong"}
	actual := ping.Process(msg)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestName(t *testing.T) {
	ping, _ := New()

	expected := "ping"
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
