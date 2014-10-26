// Validate will return an error if there is any problem with the configuration.
package config

import (
	"errors"
)

// Validate configuration. Returns an error if there is a problem.
func (c *Config) Validate() error {
	for _, r := range c.Repositories {
		err := validateRepository(&r)

		if err != nil {
			return err
		}
	}

	if c.Secret == "" {
		return errors.New("No secret")
	}

	if c.BuildLogPath == "" {
		return errors.New("No path to build log")
	}

	if c.URL == "" {
		return errors.New("No URL")
	}

	if c.Scheme() == "https" {
		if c.Cert == "" {
			return errors.New("SSL configured but no certificate")
		}

		if c.Key == "" {
			return errors.New("SSL configured but no key")
		}
	}

	return nil
}

// Validate a repository. Returns an error if there is a problem.
func validateRepository(r *Repository) error {
	for _, c := range r.Commands {
		err := validateCommand(&c)

		if err != nil {
			return err
		}
	}

	for _, n := range r.Notify {
		err := validateNotify(&n)

		if err != nil {
			return err
		}
	}

	for _, d := range r.Deploy {
		err := validateDeploy(&d)

		if err != nil {
			return err
		}
	}

	if r.URL == "" {
		return errors.New("No URL for repository")
	}

	return nil
}

// Validate a command. Returns an error if there is a problem.
func validateCommand(c *Command) error {
	if c.Name == "" {
		return errors.New("No name for command")
	}

	if c.Execute == "" {
		return errors.New("No execute for command")
	}

	return nil
}

// Validate a notification. Returns an error if there is a problem.
func validateNotify(n *Notify) error {
	if n.Service == "" {
		return errors.New("No service for notification")
	}

	if n.Service == "slack" {
		if _, ok := n.Arguments["channel"]; ok == false {
			return errors.New("No channel configured for Slack")
		}

		if _, ok := n.Arguments["endpoint"]; ok == false {
			return errors.New("No endpoint configured for Slack")
		}
	}

	if n.Service == "hipchat" {
		if _, ok := n.Arguments["channel"]; ok == false {
			return errors.New("No channel configured for HipChat")
		}

		if _, ok := n.Arguments["key"]; ok == false {
			return errors.New("No key configured for HipChat")
		}
	}

	if n.Service == "campfire" {
		if _, ok := n.Arguments["room"]; ok == false {
			return errors.New("No room configured for Campfire")
		}

		if _, ok := n.Arguments["key"]; ok == false {
			return errors.New("No key configured for Campfire")
		}

		if _, ok := n.Arguments["id"]; ok == false {
			return errors.New("No id configured for Campfire")
		}
	}

	return nil
}

// Validate a deployment. Returns an error if there is a problem.
func validateDeploy(d *Deploy) error {
	if d.Name == "" {
		return errors.New("No name")
	}

	if d.Execute == "" {
		return errors.New("No command to execute")
	}

	return nil
}
