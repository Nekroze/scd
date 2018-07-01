package cli

import (
	"github.com/Nekroze/scd/tree"
)

func Entrypoint() {
	panic("Argument Parsing Not Implemented")
	pathInput := ""
	req := tree.NewDirectoryRequest(pathInput)
	dir := tree.DetermineDirectory(req)
	if e := dir.Ensure(); e != nil {
		panic(e)
	}
	if e := dir.ChangeTo(); e != nil {
		panic(e)
	}
}
