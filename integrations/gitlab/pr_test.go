package gitlab

import (
	"encoding/json"
	"testing"
)

var mergeRequest = []byte(`{"object_kind":"merge_request","object_attributes":{"id":1,"target_branch":"master","source_branch":"test","source_project_id":1,"author_id":1,"assignee_id":null,"title":"Test","created_at":"2014-08-17 20:11:35 UTC","updated_at":"2014-08-17 20:12:05 UTC","milestone_id":null,"state":"closed","merge_status":"can_be_merged","target_project_id":1,"iid":1,"description":"test","position":0}}`)

func TestMergeUnmarshal(t *testing.T) {
	var mr MergeRequest
	err := json.Unmarshal(mergeRequest, &mr)

	if err != nil {
		t.Error("Error while unmarschal", err)
	}

	if mr.Attributes.SourceProjectId != 1 {
		t.Error("Wrong source project id", mr.Attributes.SourceProjectId)
	}
}
