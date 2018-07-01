package tree

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDirectoryRequest(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want DirectoryRequest
	}{

		{
			name: "simple local fs dir",
			arg:  "tree",
			want: DirectoryRequest{
				raw: "tree",
				url: &url.URL{
					Path: "tree",
				},
				vcs: FS,
			},
		},

		{
			name: "absolute local fs dir",
			arg:  "/tmp/storage",
			want: DirectoryRequest{
				raw: "/tmp/storage",
				url: &url.URL{
					Path: "/tmp/storage",
				},
				vcs: FS,
			},
		},

		{
			name: "https git repository",
			arg:  "https://github.com/Nekroze/scd.git",
			want: DirectoryRequest{
				raw: "https://github.com/Nekroze/scd.git",
				url: &url.URL{
					Host: "github.com",
					Path: "Nekroze/scd.git",
				},
				vcs: Git,
			},
		},

		{
			name: "ssh git repository",
			arg:  "git@github.com:Nekroze/scd.git",
			want: DirectoryRequest{
				raw: "git@github.com:Nekroze/scd.git",
				url: &url.URL{
					Host: "github.com",
					Path: "Nekroze/scd.git",
				},
				vcs: Git,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, NewDirectoryRequest(tt.arg))
		})
	}
}

func BenchmarkNewDirectoryRequest(b *testing.B) {
	for _, input := range []string{
		"tree",
		"/tmp/storage",
		"https://github.com/Nekroze/scd.git",
		"git@github.com:Nekroze/scd.git",
	} {
		b.Run(input, func(sb *testing.B) {
			for n := 0; n < sb.N; n++ {
				NewDirectoryRequest(input)
			}
		})
	}
}
