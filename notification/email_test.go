package notification

import (
	"leeroy/logging"
	"testing"
)

func TestAddHeaders(t *testing.T) {
	message := addHeaders("foo", "bar", "bla", "baz")

	if len(message) != 137 {
		t.Error("Message got the wrong length", len(message))
	}
}

func TestSubject(t *testing.T) {
	ta := logging.Task{
		Command: "foo",
	}
	j := logging.Job{
		Tasks:  []logging.Task{ta},
		Branch: "foo",
	}

	s := subject(&j)

	if s != "Build for foo finished successfully" {
		t.Error("Wrong subject", s)
	}

	ta.Return = "123"
	j.Tasks = []logging.Task{ta}
	s = subject(&j)

	if s != "Build for foo finished with errors" {
		t.Error("Wrong subject", s)
	}
}

func TestBody(t *testing.T) {
	ta := logging.Task{
		Command: "foo",
	}
	j := logging.Job{
		Tasks:  []logging.Task{ta},
		Branch: "foo",
	}

	b := body(&j, "foo")

	if len(b) != 114 {
		t.Error("Wrong body", b)
	}
}
