// Buildlog stores all builds that were triggered and the result.
package logging

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

type Buildlog struct {
	Path string
	Jobs []Job
}

func New(path string) *Buildlog {
	b := Buildlog{
		Path: path,
	}

	b.load()

	return &b
}

// Add adds a new job to the buildlog.
func (b *Buildlog) Add(j Job) {
	b.Jobs = append(b.Jobs, j)
	b.persist()
}

// Save buildlog to disk.
func (b *Buildlog) persist() {
	bin, err := json.Marshal(&b.Jobs)

	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(b.Path, bin, 0700)

	if err != nil {
		log.Println(err)
	}
}

// Load buildlog from disk.
func (b *Buildlog) load() {
	f, err := ioutil.ReadFile(b.Path)
	if err != nil {
		log.Println(err)
		return
	}

	json.Unmarshal(f, &b.Jobs)
}

// Len => sort.Interface.
func (b *Buildlog) Len() int {
	return len(b.Jobs)
}

// Swap => sort.Interface.
func (b *Buildlog) Swap(i, j int) {
	b.Jobs[i], b.Jobs[j] = b.Jobs[j], b.Jobs[i]
}

// Less => sort.Interface.
func (b *Buildlog) Less(i, j int) bool {
	return b.Jobs[i].Timestamp.UnixNano() < b.Jobs[j].Timestamp.UnixNano()
}

// Sorts the jobs based on Timestamp, newest first.
func (b *Buildlog) Sort() {
	sort.Sort(sort.Reverse(b))
}

// Returns a slice of Jobs for a specific repo.
func (b *Buildlog) JobsForRepo(repo string) []Job {
	b.Sort()
	return b.Jobs
}

// Returns a slice of Jobs for a specific repo and branch.
func (b *Buildlog) JobsForRepoBranch(repo, branch string) []Job {
	b.Sort()
	return b.Jobs
}
