package tree

import (
	"net/url"
	"regexp"
)

var gitRegexes = []*regexp.Regexp{
	regexp.MustCompile(`^\w+@(?P<Domain>[a-zA-Z0-9\.@-_]+):(?P<Path>.+.git)$`),                    // https://regex101.com/r/kLMnfa/4
	regexp.MustCompile(`^(?:https?:\/\/)?(?P<Domain>[a-zA-Z0-9\.@-_]+)\/(?P<Path>(\/?.)+\.git)$`), // https://regex101.com/r/v2nXgX/7
}

func attemptGitDirectoryRequest(input string) (match bool, new DirectoryRequest) {
	match, data := checkForMatch(input, gitRegexes)
	new.vcs = Git
	new.raw = input
	new.url = &url.URL{
		Host: data["Domain"],
		Path: data["Path"],
	}
	return match, new
}
