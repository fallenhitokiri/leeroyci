package logging

import (
	"errors"
	"testing"
)

func TestStatus(t *testing.T) {
	j := Job{}

	if j.Status() != "0" {
		t.Error("Wrong status", j.Status())
	}

	j.ReturnCode = errors.New("foo")

	if j.Status() == "0" {
		t.Error("Wrong status", j.Status())
	}
}

func TestSuccess(t *testing.T) {
	j := Job{}

	if j.Success() == false {
		t.Error("Returned error for successful build")
	}

	j.ReturnCode = errors.New("foo")

	if j.Success() == true {
		t.Error("Returned no error for failed build")
	}
}
