package deployment

import (
	"leeroy/database"
	"testing"
)

func TestContinouosTrue(t *testing.T) {
	r := database.AddRepository("foo", "http://test.tld", "", false, false, false)
	d := database.AddDeploy(r, "", "foo", "", "")

	j := logging.Job{
		URL:    "http://test.tld",
		Branch: "foo",
	}

	deployed := ContinuousDeploy(&j)

	if deployed == false {
		t.Error("Not deployed, but should")
	}
}

func TestContinouosNoRepo(t *testing.T) {
	r := database.AddRepository("foo", "http://test.tld", "", false, false, false)
	d := database.AddDeploy(r, "", "foo", "", "")

	j := logging.Job{
		URL:    "url2",
		Branch: "foo",
	}

	deployed := ContinuousDeploy(&j)

	if deployed == true {
		t.Error("Deployed, but should not")
	}
}

func TestContinouosFalse(t *testing.T) {
	r := database.AddRepository("foo", "http://test.tld", "", false, false, false)
	d := database.AddDeploy(r, "", "foo", "", "")

	j := logging.Job{
		URL:    "http://test.tld",
		Branch: "baz",
	}

	deployed := ContinuousDeploy(&j)

	if deployed == true {
		t.Error("Deployed, but should not")
	}
}
