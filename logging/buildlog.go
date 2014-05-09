// Buildlog stores all builds that were triggered and the result.
package logging

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	bin, err := json.Marshal(&b)

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

	json.Unmarshal(f, &b)
}
