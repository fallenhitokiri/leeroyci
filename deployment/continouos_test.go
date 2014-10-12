package deployment

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestContinouosTrue(t *testing.T) {
	c := config.Config{}
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	c.Repositories = append(c.Repositories, r)

	j := logging.Job{
		URL:    "url",
		Branch: "foo",
	}

	deployed := ContinuousDeploy(&j, &c)

	if deployed == false {
		t.Error("Not deployed, but should")
	}
}

func TestContinouosNoRepo(t *testing.T) {
	c := config.Config{}
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	c.Repositories = append(c.Repositories, r)

	j := logging.Job{
		URL:    "url2",
		Branch: "foo",
	}

	deployed := ContinuousDeploy(&j, &c)

	if deployed == true {
		t.Error("Deployed, but should not")
	}
}

func TestContinouosFalse(t *testing.T) {
	c := config.Config{}
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	c.Repositories = append(c.Repositories, r)

	j := logging.Job{
		URL:    "url",
		Branch: "baz",
	}

	deployed := ContinuousDeploy(&j, &c)

	if deployed == true {
		t.Error("Deployed, but should not")
	}
}
