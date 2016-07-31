package database

import (
	"github.com/satori/go.uuid"
)

// InMemory is a simple in memory database without persistence. This should
// never be used in production.
type InMemory struct {
	config     *Config
	users      []*User
	projects   []*Project
	mailserver *MailServer
}

// Config returns the current configuration.
func (m *InMemory) Config() (*Config, error) {
	return m.config, nil
}

// ConfigUpdate updates the current configuration.
func (m *InMemory) ConfigUpdate(cfg *Config) error {
	m.config = cfg
	return nil
}

// UserList returns a list of all users.
func (m *InMemory) UserList() ([]*User, error) {
	return m.users, nil
}

// UserByEmail returns the user for a given email address.
func (m *InMemory) UserByEmail(email string) (*User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrorUserNotFound
}

// UserBySessionKey returns the user for a given session key.
func (m *InMemory) UserBySessionKey(key string) (*User, error) {
	for _, user := range m.users {
		for _, k := range user.Sessions {
			if k.Key == key {
				return user, nil
			}
		}
	}
	return nil, ErrorUserNotFound
}

// UserByAPIKey returns the user for a given API key.
func (m *InMemory) UserByAPIKey(key string) (*User, error) {
	for _, user := range m.users {
		for _, api := range user.APIKeys {
			if api.Key == key {
				return user, nil
			}
		}
	}
	return nil, ErrorUserNotFound
}

// UserByID returns a user for a given ID.
func (m *InMemory) UserByID(id string) (*User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrorUserNotFound
}

// UserCreate creates a new user.
func (m *InMemory) UserCreate(user *User) error {
	for {
		uid := uuid.NewV4().String()
		_, err := m.UserByID(uid)
		if err != nil {
			user.ID = uid
			break
		}
	}

	if _, err := m.UserByEmail(user.Email); err == nil {
		return ErrorUserEmailExists
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	m.users = append(m.users, user)
	return nil
}

// UserUpdatePassword updates a users password.
func (m *InMemory) UserUpdatePassword(user *User, password string) error {
	hash, err := hashPassword(password)

	if err != nil {
		return err
	}

	user.Password = hash
	return nil
}

// UserDelete deletes a user from the datastore.
func (m *InMemory) UserDelete(user *User) error {
	for index, u := range m.users {
		if u.ID == user.ID {
			m.users = append(m.users[:index], m.users[index+1:]...)
			return nil
		}
	}
	return ErrorUserNotFound
}

// ProjectCreate creates a new project.
func (m *InMemory) ProjectCreate(project *Project) error {
	if _, err := m.ProjectByName(project.Name); err == nil {
		return ErrorProjectNameExists
	}
	m.projects = append(m.projects, project)
	return nil
}

// ProjectList returns all projects.
func (m *InMemory) ProjectList() ([]*Project, error) {
	return m.projects, nil
}

// ProjectByName returns a project for a given name.
func (m *InMemory) ProjectByName(name string) (*Project, error) {
	for _, project := range m.projects {
		if project.Name == name {
			return project, nil
		}
	}
	return nil, ErrorProjectNotFound
}

// ProjectByAccessKey returns a project for a given access key.
func (m *InMemory) ProjectByAccessKey(key string) (*Project, error) {
	for _, project := range m.projects {
		if project.AccessKey == key {
			return project, nil
		}
	}
	return nil, ErrorProjectNotFound
}

// ProjectUpdate updates a project.
func (m *InMemory) ProjectUpdate(project *Project) error {
	for index, p := range m.projects {
		if p.Name == project.Name {
			m.projects = append(m.projects[:index], m.projects[index+1:]...)
			m.projects = append(m.projects, project)
			return nil
		}
	}
	return ErrorProjectNotFound
}

// ProjectDelete removes a project form the db.
func (m *InMemory) ProjectDelete(project *Project) error {
	for index, p := range m.projects {
		if p.Name == project.Name {
			m.projects = append(m.projects[:index], m.projects[index+1:]...)
		}
	}
	return ErrorProjectNotFound
}

// TaskCreate creates a new task for a project.
func (m *InMemory) TaskCreate(task *Task, project *Project) error {
	return nil
}

// func (m *InMemory) TaskUpdate(task *Task, project *Project) error {}

// func (m *InMemory) TaskDelete(task *Task, project *Project) error {}
//
// func (m *InMemory) NotificationCreate(not *Notification, project *Project) error {}

// func (m *InMemory) NotificationUpdate(not *Notification, project *Project) error {}

// func (m *InMemory) NotificationDelete(not *Notification, project *Project) error {}
//

// func (m *InMemory) BranchCreate(branch *Branch, project *Project) error {}

// func (m *InMemory) BranchUpdate(branch *Branch, project *Project) error {}

// func (m *InMemory) BranchDelete(branch *Branch, project *Project) error {}
//

// func (m *InMemory) MailServer() (*MailServer, error)      {}

// func (m *InMemory) MailServerUpdate(ms *MailServer) error {}
