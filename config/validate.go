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

	return nil
}
