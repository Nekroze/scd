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

func TestFuzzyGitCollector(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		arg  FileInfo
		want []entry
	}{
		{
			name: "empty result for file",
			arg: FileInfo{
				FileInfo: makeFileInfo(false, "thing"),
				Path:     "~/thing",
			},
			want: []entry{},
		},
		{
			name: "empty result for non .git dir",
			arg: FileInfo{
				FileInfo: makeFileInfo(true, "thing"),
				Path:     "~/thing",
			},
			want: []entry{},
		},
		{
			name: "parent result for .git dir",
			arg: FileInfo{
				FileInfo: makeFileInfo(true, ".git"),
				Path:     "/thing/.git",
			},
			want: []entry{{
				path: "/thing/",
				hay:  "thing",
			}},
		},
		{
			name: "parent result for realistic .git dir",
			arg: FileInfo{
				FileInfo: makeFileInfo(true, ".git"),
				Path:     "/home/nekroze/git/github.com/Nekroze/scd/.git",
			},
			want: []entry{{
				path: "/home/nekroze/git/github.com/Nekroze/scd/",
				hay:  "scd",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fuzzyGitCollector(tt.arg))
		})
	}
}

func TestFuzzyDictionary_Search(t *testing.T) {
	type args struct {
		input string
		limit int
	}
	tests := []struct {
		name    string
		entries []entry
		args    args
		want    string
	}{
		{
			name:    "no entry gives no output",
			entries: []entry{},
			args: args{
				input: "foo",
				limit: 4,
			},
			want: "",
		},
		{
			name: "no input gives no output",
			entries: []entry{
				{path: "~/thing/", hay: "thing"},
			},
			args: args{
				input: "",
				limit: 4,
			},
			want: "",
		},
		{
			name: "exact input gives same output",
			entries: []entry{
				{path: "~/foo/", hay: "foo"},
				{path: "~/bar/", hay: "bar"},
			},
			args: args{
				input: "foo",
				limit: 4,
			},
			want: "~/foo/",
		},
		{
			name: "exact input gives same output with ambiguous dictionary",
			entries: []entry{
				{path: "~/foo/", hay: "foo"},
				{path: "~/foish/", hay: "foish"},
			},
			args: args{
				input: "foo",
				limit: 4,
			},
			want: "~/foo/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &FuzzyDictionary{
				entries: tt.entries,
			}
			assert.Equal(t, tt.want, d.Search(tt.args.input, tt.args.limit))
		})
	}
}
