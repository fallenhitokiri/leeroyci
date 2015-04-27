package deployment

import (
	"leeroy/config"
	"leeroy/logging"
	"testing"
)

func TestContinouosTrue(t *testing.T) {
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	config.CONFIG.Repositories = append(config.CONFIG.Repositories, r)

	j := logging.Job{
		URL:    "url",
		Branch: "foo",
	}

	deployed := ContinuousDeploy(&j)

	if deployed == false {
		t.Error("Not deployed, but should")
	}
}

func TestContinouosNoRepo(t *testing.T) {
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	config.CONFIG.Repositories = append(config.CONFIG.Repositories, r)

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
	r := config.Repository{
		URL: "url",
	}
	d := config.Deploy{
		Branch: "foo",
	}
	r.Deploy = append(r.Deploy, d)
	config.CONFIG.Repositories = append(config.CONFIG.Repositories, r)

	j := logging.Job{
		URL:    "url",
		Branch: "baz",
	}

	deployed := ContinuousDeploy(&j)

	if deployed == true {
		t.Error("Deployed, but should not")
	}
}
