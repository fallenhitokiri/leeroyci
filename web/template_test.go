package web

import (
	"testing"
)

func TestGetTemplates(t *testing.T) {
	template := getTemplates("job/list.html")

	if template.Name() != "job/list.html" {
		t.Error("Got wrong template", template.Name())
	}

	if len(template.Templates()) != 6 {
		t.Error("Wrong template count", len(template.Templates()))
	}
}
