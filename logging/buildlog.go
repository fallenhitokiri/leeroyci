// Buildlog stores all builds that were triggered and the result.
package logging

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"sync"
)

type Buildlog struct {
	Path  string
	Jobs  []Job
	mutex sync.Mutex
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
	b.mutex.Lock()
	b.Jobs = append(b.Jobs, j)
	b.mutex.Unlock()

	b.persist()
}

// Save buildlog to disk.
func (b *Buildlog) persist() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

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
	b.mutex.Lock()
	defer b.mutex.Unlock()

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
	b.mutex.Lock()
	defer b.mutex.Unlock()

	sort.Sort(sort.Reverse(b))
}

// Returns a slice of Jobs for a specific repo.
func (b *Buildlog) JobsForRepo(repo string) []Job {
	b.Sort()

	b.mutex.Lock()
	defer b.mutex.Unlock()

	var jobs = make([]Job, 0)

	for _, j := range b.Jobs {
		if j.URL == repo {
			jobs = append(jobs, j)
		}
	}

	return jobs
}

// Returns a slice of Jobs for a specific repo and branch.
func (b *Buildlog) JobsForRepoBranch(repo, branch string) []Job {
	b.Sort()

	b.mutex.Lock()
	defer b.mutex.Unlock()

	var jobs = make([]Job, 0)

	for _, j := range b.Jobs {
		if j.URL == repo && j.Branch == branch {
			jobs = append(jobs, j)
		}
	}

	return jobs
}
