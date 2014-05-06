// Buildlog stores all builds that were triggered and the result.
package logging

type Buildlog struct {
	Jobs []Job
}

// Add adds a new job to the buildlog
func (b *Buildlog) Add(j Job) {
	b.Jobs = append(b.Jobs, j)
}
