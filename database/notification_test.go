package database

import (
	"testing"
)

func TestNotificationCRUD(t *testing.T) {
	r := CreateRepository("name", "url", "accessKey", false, false, false)
	n1 := CreateNotification("service", "arguments", r)
	n2 := GetNotification(n1.ID)
	n3 := UpdateNotification(n1.ID, "service", "arguments2")
	DeleteNotification(n1.ID)
	n4 := GetNotification(n1.ID)

	if n1.Service != n2.Service {
		t.Error("Service does not match")
	}

	if n3.Arguments == n2.Arguments {
		t.Error("Arguments not updated")
	}

	if n4.ID == n1.ID || n4.ID != 0 {
		t.Error("Notification not deleted")
	}
}
