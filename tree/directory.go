// package tree handles deciding where underneath the users $HOME a directory
// should be located.
package tree

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

type VCS int

const (
	FS VCS = iota
	Git
)

type Directory struct {
	Path string // The absolute path to this directory.
	req  DirectoryRequest
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
	switch d.req.vcs {
	case FS:
		err = os.MkdirAll(d.Path, os.ModePerm)
	case Git:
		_, err = git.PlainClone(d.Path, false, &git.CloneOptions{
			URL:      d.req.raw,
			Progress: os.Stderr,
		})
	}
	return err
}

func NewDirectory(req DirectoryRequest) (out Directory) {
	parts := []string{}
	out.req = req
	switch req.vcs {
	case Git:
		parts = append(parts, []string{
			os.Getenv("HOME"),
			"git",
			req.url.Host,
			req.url.Path,
		}...)
	default:
		if len(strings.Split(req.raw, string(filepath.Separator))) > 0 {
			parts = append(parts, req.raw)
			break
		}
		dict := &FuzzyDictionary{}
		filepath.Walk(os.Getenv("HOME"), dict.Wanderer)
		found := dict.Search(req.raw, 5)
		if found == "" {
			found = req.raw
		}
		parts = append(parts, found)
	}
	out.Path = filepath.Join(parts...)
	return out
}
