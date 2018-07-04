package tree

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

//go:generate moq -out fileinfo_moq_test.go . osFileInfo

type osFileInfo interface {
	os.FileInfo
}

type FileInfo struct {
	os.FileInfo
	Path string
}

// FuzzyCollector functions are used when using path.filepath.Walk to build a
// dictionary for fuzzy matching.
type FuzzyCollector func(FileInfo) []entry

var fuzzyCollectors = []FuzzyCollector{
	fuzzyGitCollector,
}

func fuzzyGitCollector(info FileInfo) []entry {
	out := []entry{}
	path, dir := filepath.Split(info.Path)
	if info.IsDir() && dir == ".git" {
		out = append(out, entry{
			path: path,
			hay:  filepath.Base(path),
		})
	}
	return out
}

type entry struct {
	path string // the directory that hay corresponds to
	hay  string // the string to fuzzy match against
}

type FuzzyDictionary struct {
	entries []entry
}

func NewDictionary(info FileInfo) (new *FuzzyDictionary) {
	new.Learn(info)
	return new
}

func (d *FuzzyDictionary) Wanderer(path string, info os.FileInfo, err error) error {
	d.Learn(FileInfo{
		FileInfo: info,
		Path:     path,
	})
	return err
}

func (d *FuzzyDictionary) Learn(info FileInfo) {
	for _, fc := range fuzzyCollectors {
		found := append(d.entries, fc(info)...)
		if len(found) > 0 {
			d.entries = append(d.entries, found...)
		}
	}
}

func (d *FuzzyDictionary) fuzzyData() ([]string, map[string]string) {
	dict := []string{}
	rev := map[string]string{}
	for _, e := range d.entries {
		rev[e.hay] = e.path
		dict = append(dict, e.hay)
	}
	return dict, rev
}

func (d *FuzzyDictionary) Search(input string, limit int) string {
	dict, rev := d.fuzzyData()
	matches := fuzzy.RankFind(input, dict)
	sort.Sort(matches)
	var closest fuzzy.Rank
	for _, match := range matches {
		if closest.Distance == 0 || match.Distance < closest.Distance {
			closest = match
		}
	}
	if closest.Distance > limit {
		return ""
	}
	return rev[closest.Target]
}
