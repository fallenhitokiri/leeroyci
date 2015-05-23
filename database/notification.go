// Package database provides a wrapper between the database and stucts
package database

const (
	// Notification type for email notifications.
	NotificationServiceEmail = "email"

	// Notification type for slack notifications.
	NotificationServiceSlack = "slack"
)

// Notify stores the configuration needed for a notification plugin to work. All
// optiones required by the services are stored as map - it is the job of the
// notification plugin to access them correctly and handle missing ones.
type Notification struct {
	ID        int64
	Service   string
	Arguments string

	Repository   Repository
	RepositoryID int64
}

// CreateNotify create a new notification for a repository.
func CreateNotification(service, arguments string, repo *Repository) *Notification {
	not := Notification{
		Service:    service,
		Arguments:  arguments,
		Repository: *repo,
	}

	db.Save(&not)

	return &not
}

// GetNotification returns a notification.
func GetNotification(id int64) *Notification {
	not := &Notification{}
	db.Where("id = ?", id).First(&not)
	return not
}

// UpdateNotification updates a notification.
func UpdateNotification(id int64, service, arguments string) *Notification {
	not := GetNotification(id)
	not.Service = service
	not.Arguments = arguments
	db.Save(not)
	return not
}

// DeleteNotification deletes a notification.
func DeleteNotification(id int64) {
	not := GetNotification(id)
	db.Delete(not)
}
