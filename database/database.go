package database

var db *Database

// Database defines all methos which have to be supported by a database.
type Database interface {
	Config() (*Config, error)
	ConfigUpdate(cfg *Config) error

	UserList() ([]*User, error)
	UserByEmail(email string) (*User, error)
	UserBySessionKey(key string) (*User, error)
	UserByAPIKey(key string) (*User, error)
	UserByID(id string) (*User, error)
	UserCreate(user *User) error
	UserUpdatePassword(user *User, password string) error
	UserDelete(user *User) error

	ProjectCreate(project *Project) error
	ProjectList() ([]*Project, error)
	ProjectByName(name string) (*Project, error)
	ProjectByAccessKey(key string) (*Project, error)
	ProjectUpdate(project *Project) error
	ProjectDelete(project *Project) error

	TaskCreate(task *Task, project *Project) error
	TaskUpdate(task *Task, project *Project) error
	TaskDelete(task *Task, project *Project) error

	NotificationCreate(not *Notification, project *Project) error
	NotificationUpdate(not *Notification, project *Project) error
	NotificationDelete(not *Notification, project *Project) error

	BranchCreate(branch *Branch, project *Project) error
	BranchUpdate(branch *Branch, project *Project) error
	BranchDelete(branch *Branch, project *Project) error

	MailServer() (*MailServer, error)
	MailServerUpdate(ms *MailServer) error
}
