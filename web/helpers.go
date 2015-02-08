// Provide helpers for requests.
package web

import (
	"encoding/hex"
	"leeroy/logging"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	paginateBy = 10
)

// Splits a request path and returns the first part after the endpoint.
// This is usually the hex string of the repository.
func splitFirst(path string) string {
	s := strings.Split(path, "/")[3]
	r, err := hex.DecodeString(s)

	if err != nil {
		log.Println(err)
	}

	return string(r)
}

// Splits a request path and returns the first part after the endpoint.
// This is likely the branch name or commit sha1.
func splitSecond(path string) string {
	return strings.Split(path, "/")[4]
}

// paginatedJobs returns a part of the jobs slices starting at 'start', the next
// start element and the previous element.
func paginatedJobs(jobs []*logging.Job, start string) ([]*logging.Job, string, string) {
	c := len(jobs) - 1
	f := paginateGetFirst(start, c)
	l := paginateGetLast(f, c)

	return jobs[f:l], paginateGetNext(f, c), paginateGetPrevious(f)
}

// paginateGetPrevious returns the previous index for a paginated job list.
func paginateGetPrevious(first int) string {
	p := first - paginateBy + 1 // slice index vs. count

	if p-1 == paginateBy*-1 {
		return ""
	}

	if p < 0 {
		return "0"
	}

	return strconv.Itoa(p)
}

// paginateGetNext returns the next index for a paginated job list.
func paginateGetNext(first, count int) string {
	log.Println(first)
	n := first + paginateBy

	if first != 0 {
		n += 1
	}

	log.Println(n)

	if n >= count {
		return ""
	}

	return strconv.Itoa(n)
}

// paginateGetFirst returns the first element of the job slice to return.
func paginateGetFirst(start string, count int) int {
	f, err := strconv.Atoi(start)

	if err != nil {
		log.Fatalln("Could not convert start to int: ", start)
	}

	if f > 0 {
		f -= 1
	}

	if f > count {
		log.Fatalln("Start index out of range: ", f)
	}

	return f
}

// paginateGetLast returns the last element of the job slice to return.
func paginateGetLast(first, count int) int {
	l := first + paginateBy

	if l > count {
		l = count
	}

	return l
}

// getParameter tries to get key form the URL parameters and returns it. If this
// is not possible 'def' will be returned.
func getParameter(req *http.Request, key, def string) string {
	params, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		log.Fatalln(err)
	}

	if val, ok := params[key]; ok == true {
		return val[0]
	}

	return def
}
