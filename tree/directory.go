// package tree handles deciding where underneath the users $HOME a directory
// should be located.
package tree

import (
	"net/url"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
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

func (d Directory) exists() (exists bool, err error) {
	_, err = os.Stat(d.Path)
	if err == nil {
		exists = true
	} else if os.IsNotExist(err) {
		err = nil
	}
	return exists, err
}

func (d Directory) Ensure() error {
	exists, err := d.exists()
	if exists || err != nil {
		return err
	}
	switch d.Type {
	case FS:
		err = os.MkdirAll(d.Path, os.ModePerm)
	case Git:
		_, err = git.PlainClone(d.Path, false, &git.CloneOptions{
			URL:      d.url.String(),
			Progress: os.Stderr,
		})
	}
	return err
}

func DetermineDirectory(req DirectoryRequest) (out Directory) {
	parts := []string{}
	out.url = req.url
	switch req.url.Scheme {
	case "git":
		out.Type = Git
		parts = append(parts, []string{
			os.Getenv("HOME"),
			"git",
			req.url.Host,
			req.url.Path,
		}...)
	default:
		out.Type = FS
		parts = append(parts, req.url.Path)
	}
	out.Path = filepath.Join(parts...)
	return out
}
