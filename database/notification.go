package database

const (
	// NotificationServiceEmail type for email notifications.
	NotificationServiceEmail = "email"

	// NotificationServiceSlack type for slack notifications.
	NotificationServiceSlack = "slack"

	// NotificationServiceCampfire type for campfire notifications.
	NotificationServiceCampfire = "campfire"

	// NotificationServiceHipchat type for hipchat notifications.
	NotificationServiceHipchat = "hipchat"
)

// Notification stores the configuration needed for a notification plugin to
// work. All optiones required by the services are stored as map - it is the
// job of the notification plugin to access them correctly and handle missing
// ones.
type Notification struct {
	ID        string
	Service   string
	Arguments string
}
