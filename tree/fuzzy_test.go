package tree

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeFileInfo(dir bool, name string) os.FileInfo {
	// make and configure a mocked osFileInfo
	return &osFileInfoMock{
		IsDirFunc: func() bool {
			return dir
		},
		ModTimeFunc: func() time.Time {
			panic("TODO: mock out the ModTime method")
		},
		ModeFunc: func() os.FileMode {
			panic("TODO: mock out the Mode method")
		},
		NameFunc: func() string {
			return name
		},
		SizeFunc: func() int64 {
			panic("TODO: mock out the Size method")
		},
		SysFunc: func() interface{} {
			panic("TODO: mock out the Sys method")
		},
	}
}

func Test_fuzzyGitCollector(t *testing.T) {
	type args struct {
		info FileInfo
	}
	tests := []struct {
		name string
		args args
		want []entry
	}{
		{
			name: "empty result for file",
			args: args{
				info: FileInfo{
					FileInfo: makeFileInfo(false, "thing"),
					Path:     "~/thing",
				},
			},
			want: []entry{},
		},
		{
			name: "empty result for non .git dir",
			args: args{
				info: FileInfo{
					FileInfo: makeFileInfo(true, "thing"),
					Path:     "~/thing",
				},
			},
			want: []entry{},
		},
		{
			name: "parent result for .git dir",
			args: args{
				info: FileInfo{
					FileInfo: makeFileInfo(true, ".git"),
					Path:     "/thing/.git",
				},
			},
			want: []entry{{
				path: "/thing/",
				hay:  "thing",
			}},
		},
		{
			name: "parent result for realistic .git dir",
			args: args{
				info: FileInfo{
					FileInfo: makeFileInfo(true, ".git"),
					Path:     "/home/nekroze/git/github.com/Nekroze/scd/.git",
				},
			},
			want: []entry{{
				path: "/home/nekroze/git/github.com/Nekroze/scd/",
				hay:  "scd",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fuzzyGitCollector(tt.args.info))
		})
	}
}
