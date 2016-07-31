package database

// Branch represents one branch in the git repository.
type Branch struct {
	ID          string
	Name        string
	ProjectName string

	Results []*Result
}
