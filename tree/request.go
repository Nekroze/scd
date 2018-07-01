package tree

import (
	"net/url"
	"regexp"
)

type DirectoryRequest struct {
	url *url.URL
}

func NewDirectoryRequest(userInput string) (new DirectoryRequest, err error) {
	for _, generator := range drGenerators {
		if ok, dr := generator(userInput); ok {
			return dr, nil
		}
	}
	new.url, err = url.Parse(userInput)
	return new, err
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i, _ := range n {
		r[n[i]] = m[i]
	}
	return r
}

func checkForMatch(input string, regexes []*regexp.Regexp) (bool, map[string]string) {
	for _, r := range regexes {
		if m := r.FindStringSubmatch(input); len(m) > 0 {
			return true, mapSubexpNames(m, r.SubexpNames())
		}
	}
	return false, nil
}

type directoryRequestGenerator func(string) (bool, DirectoryRequest)

var drGenerators = []directoryRequestGenerator{
	attemptGitDirectoryRequest,
}
