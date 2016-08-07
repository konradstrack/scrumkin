package ping

import (
	"reflect"
	"scrumkin/pkg/messages"
	"testing"
)

func TestNew(t *testing.T) {
	ping, _ := New()
	if ping.Name != "ping" {
		t.Error("Expected name to be 'ping', got", ping.Name)
	}
}

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
