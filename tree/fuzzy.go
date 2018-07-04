package tree

import (
	"os"
	"sort"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

type FileInfo struct {
	os.FileInfo
	Path string
}

// FuzzyCollector functions are used when using path.filepath.Walk to build a
// dictionary for fuzzy matching.
type FuzzyCollector func(FileInfo) []entry

var fuzzyCollectors = []FuzzyCollector{}

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
	if closest.Distance < limit {
		return ""
	}
	return rev[closest.Target]
}
