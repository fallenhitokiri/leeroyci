package database

import (
	"testing"
)

func TestIdentifier(t *testing.T) {
	r := Repository{
		URL: "foo",
	}

	i := r.Identifier()

	if i != "foo" {
		t.Error("Wrong identifier", i)
	}

	r.Name = "bar"

	i = r.Identifier()

	if i != "bar" {
		t.Error("Wrong identifier", i)
	}
}

func TestDeploymentTarget(t *testing.T) {
	r := Repository{}

	d1 := Deploy{
		Branch:  "foo",
		Execute: "123",
	}

	d2 := Deploy{
		Branch: "bar",
	}

	r.Deploy = append(r.Deploy, d1, d2)

	dep, err := r.DeployTarget("foo")

	if err != nil {
		t.Error(err)
	}

	if dep.Execute != "123" {
		t.Error("Wrong execute", dep.Execute)
	}
}
