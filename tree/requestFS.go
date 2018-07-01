package tree

import (
	"net/url"
)

func generateFSDirectoryRequest(input string) (match bool, new DirectoryRequest) {
	match = true
	new.vcs = FS
	new.raw = input
	var err error
	new.url, err = url.Parse(input)
	if err != nil {
		new.url = &url.URL{
			Path: input,
		}
	}
	return match, new
}
