package database

import (
	"testing"
)

func TestNotificationCRUD(t *testing.T) {
	r, _ := CreateRepository("name", "url", "accessKey", false, false)
	n1, _ := CreateNotification("service", "arguments", r)
	n2, _ := GetNotification(string(n1.ID))
	n2.Update("service", "arguments2")
	n2.Delete()
	n3, _ := GetNotification(string(n1.ID))

	if n1.Service != n2.Service {
		t.Error("Service does not match")
	}

	if n2.Arguments == "arguments" {
		t.Error("Arguments not updated")
	}

	if n3.ID == n1.ID || n3.ID != 0 {
		t.Error("Notification not deleted")
	}
}

func TestGetNotificationForRepoAndType(t *testing.T) {
	r, _ := CreateRepository("name", "url", "accessKey", false, false)
	not, _ := CreateNotification(NotificationServiceSlack, "arguments", r)

	notGot, _ := GetNotificationForRepoAndType(r, NotificationServiceSlack)

	if not.ID != notGot.ID {
		t.Error("got the wrong notification.")
	}
}

func TestGetConfigValue(t *testing.T) {
	r, _ := CreateRepository("name", "url", "accessKey", false, false)
	not, _ := CreateNotification(NotificationServiceSlack, "", r)

	_, err := not.GetConfigValue("foo")
	if err.Error() != "No Arguments defined." {
		t.Error("Wrong return", err.Error())
	}

	not, _ = CreateNotification(NotificationServiceSlack, "foo:::bar:::::zab:::123", r)

	_, err = not.GetConfigValue("baz")
	if err.Error() != "Not found." {
		t.Error("Wrong return", err.Error())
	}

	val, _ := not.GetConfigValue("zab")
	if val != "123" {
		t.Error("Wrong return", val)
	}
}
