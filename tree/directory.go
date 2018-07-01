// package tree handles deciding where underneath the users $HOME a directory
// should be located.
package tree

import (
	"net/url"
	"os"
	"path"
)

type VCS int

const (
	FS VCS = iota
	Git
)

type Directory struct {
	Path string // The absolute path to this directory.
	Type VCS    // The type of directory either FS or a VCS.
	url  *url.URL
}

func (d Directory) Ensure() (err error) {
	panic("Check Exists Not Implemented")
	switch d.Type {
	case FS:
		panic("Mkdir Not Implemented")
	case Git:
		panic("Git Clone Not Implemented")
	}
	return err
}

func (d Directory) ChangeTo() (err error) {
	panic("Change To Direcotry Not Implemented")
	return err
}

func DetermineDirectory(req DirectoryRequest) (out Directory) {
	parts := []string{
		os.Getenv("HOME"),
	}
	out.url = req.url
	switch req.url.Scheme {
	case "git":
		out.Type = Git
		parts = append(parts, []string{
			"git",
			req.url.Host,
			req.url.Path,
		}...)
	default:
		out.Type = FS
		parts = append(parts, req.url.Path)
	}
	out.Path = path.Join(parts...)
	return out
}
