package data

// Branch represents one branch in the git repository.
type Branch struct {
	Name        string
	ProjectName string

	Results []*Result
}
