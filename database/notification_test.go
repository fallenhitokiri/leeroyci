package database

import (
	"testing"
)

func TestNotificationCRUD(t *testing.T) {
	r := CreateRepository("name", "url", "accessKey", false, false, false)
	n1 := CreateNotification("service", "arguments", r)
	n2 := GetNotification(n1.ID)
	n2.Update("service", "arguments2")
	n2.Delete()
	n3 := GetNotification(n1.ID)

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
