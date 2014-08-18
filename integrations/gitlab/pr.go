// Structs and methods used to process a pull request.
package gitlab

import (
	"encoding/json"
	"leeroy/config"
	"leeroy/logging"
	"log"
)

type MergeRequest struct {
	Attributes Attributes `json:"object_attributes"`
}

type Attributes struct {
	Id              int
	TargetBranch    string `json:"target_branch"`
	SourceBranch    string `json:"source_branch"`
	SourceProjectId int    `json:"source_project_id"`
	AuthorId        int    `json:"author_id"`
	AssigneeId      int    `json:"assignee_id"`
	Title           string
	Created         string `json:"created_at"`
	Updated         string `json:"updated_at"`
	MilestoneId     int    `json:"milestone_id"`
	State           string
	MergeStatus     string `json:"merge_status"`
	TargetProjectId int    `json:"target_project_id"`
	Iid             int
	Description     string
	Position        int
}

// Handle GitLab pull requests.
func handlePR(body []byte, blog *logging.Buildlog, c *config.Config) {
	var mr MergeRequest
	err := json.Unmarshal(body, &mr)

	if err != nil {
		log.Println(string(body))
		panic("Could not unmarshal request")
	}

	if mr.Attributes.State != "closed" {
		log.Println("handling pull request", mr.Attributes.Id)
	}
}
