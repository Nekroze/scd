package tree

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDirectory(t *testing.T) {
	tests := []struct {
		name     string
		arg      DirectoryRequest
		wantPath string
	}{
		{
			name: "simple dir",
			arg: DirectoryRequest{
				raw: "tree",
				url: &url.URL{
					Host: "",
					Path: "tree",
				},
				vcs: FS,
			},
			wantPath: "tree",
		},
		{
			name: "nested dir",
			arg: DirectoryRequest{
				raw: "things/stuff",
				url: &url.URL{
					Host: "",
					Path: "things/stuff",
				},
				vcs: FS,
			},
			wantPath: "things/stuff",
		},
		{
			name: "http github",
			arg: DirectoryRequest{
				raw: "https://github.com/Nekroze/scd.git",
				url: &url.URL{
					Host: "github.com",
					Path: "Nekroze/scd.git",
				},
				vcs: Git,
			},
			wantPath: os.Getenv("HOME") + "/git/github.com/Nekroze/scd",
		},
		{
			name: "ssh github",
			arg: DirectoryRequest{
				raw: "git@github.com:Nekroze/scd.git",
				url: &url.URL{
					Host: "github.com",
					Path: "Nekroze/scd.git",
				},
				vcs: Git,
			},
			wantPath: os.Getenv("HOME") + "/git/github.com/Nekroze/scd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantPath, NewDirectory(tt.arg).Path)
		})
	}
}
