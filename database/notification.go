// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"
	"strings"
)

const (
	// NotificationServiceEmail type for email notifications.
	NotificationServiceEmail = "email"

	// NotificationServiceSlack type for slack notifications.
	NotificationServiceSlack = "slack"

	// NotificationServiceCampfire type for campfire notifications.
	NotificationServiceCampfire = "campfire"

	// NotificationServiceHipChat type for hipchat notifications.
	NotificationServiceHipchat = "hipchat"
)

// Notification stores the configuration needed for a notification plugin to
// work. All optiones required by the services are stored as map - it is the
// job of the notification plugin to access them correctly and handle missing
// ones.
type Notification struct {
	ID        int64
	Service   string
	Arguments string

	Repository   Repository
	RepositoryID int64 `sql:"index"`
}

// CreateNotification create a new notification for a repository.
func CreateNotification(service, arguments string, repo *Repository) (*Notification, error) {
	not := Notification{
		Service:    service,
		Arguments:  arguments,
		Repository: *repo,
	}

	db.Save(&not)

	return &not, nil
}

// GetNotification returns a notification.
func GetNotification(id string) (*Notification, error) {
	not := &Notification{}
	db.Where("id = ?", id).First(&not)
	return not, nil
}

// GetNotificationForRepoAndType returns a specific notification for a repository.
func GetNotificationForRepoAndType(repo *Repository, service string) (*Notification, error) {
	not := &Notification{}
	db.Where("repository_id = ? AND service = ?", repo.ID, service).First(&not)
	return not, nil
}

// UpdateNotification updates a notification.
func (n *Notification) Update(service, arguments string) error {
	n.Service = service
	n.Arguments = arguments
	db.Save(n)
	return nil
}

// DeleteNotification deletes a notification.
func (n *Notification) Delete() {
	db.Delete(n)
}

// GetConfigValue returns a configuration value for a given key that is stored in Arguments.
func (n *Notification) GetConfigValue(key string) (string, error) {
	if n.Arguments == "" {
		return "", errors.New("No Arguments defined.")
	}

	for _, pair := range strings.Split(n.Arguments, ":::::") {
		split := strings.Split(pair, ":::")

		if split[0] == key {
			return split[1], nil
		}
	}

	return "", errors.New("Not found.")
}
