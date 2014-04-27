package build

import (
	"ironman/callbacks"
	"ironman/config"
	"log"
)

// Build waits for new notifications and runs the build process after
// receiving one.
func Build(not chan callbacks.Notification) {
	for {
		n := <-not
		log.Println(n.Branch())
		log.Println(n.URL())
	}
}
