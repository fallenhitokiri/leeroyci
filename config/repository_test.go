package config

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	r1 := Repository{
		Name: "foobar",
		URL:  "baz",
	}
	r2 := Repository{
		URL: "baz",
	}

	if r1.Identifier() != "foobar" {
		t.Error("Wrong identifier", r1.Identifier())
	}

	if r2.Identifier() != "baz" {
		t.Error("Wrong identifier", r2.Identifier())
	}
}

func TestDeployTarget(t *testing.T) {
	r := Repository{}
	d1 := Deploy{
		Branch: "foo",
	}
	d2 := Deploy{
		Branch: "bar",
	}

	r.Deploy = append(r.Deploy, d1)

	_, err := r.DeployTarget("foo")

	if err != nil {
		t.Error("Got an error, should not:", err)
	}

	_, err = r.DeployTarget("bar")

	if err == nil {
		t.Error("Got no error, but should!")
	}

	r.Deploy = append(r.Deploy, d2)

	_, err = r.DeployTarget("bar")

	if err != nil {
		t.Error("Got an error, should not:", err)
	}
}
