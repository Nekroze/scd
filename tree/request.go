package tree

import (
	"net/url"
)

type DirectoryRequest struct {
	url *url.URL
}

func NewDirectoryRequest(userInput string) (new DirectoryRequest, err error) {
	new.url, err = url.Parse(userInput)
	return
}
