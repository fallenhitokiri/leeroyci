// Package database provides a wrapper between the database and stucts
package database

const (
	// NotificationServiceEmail type for email notifications.
	NotificationServiceEmail = "email"

	// NotificationServiceSlack type for slack notifications.
	NotificationServiceSlack = "slack"
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
