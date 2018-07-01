package cli

import (
	"fmt"
	"os"

	"github.com/Nekroze/scd/tree"
)

func Entrypoint() {
	if len(os.Args) != 2 {
		fmt.Println("You must provide an argument")
		os.Exit(1)
	}
	pathInput := os.Args[1]
	req, err := tree.NewDirectoryRequest(pathInput)
	if err != nil {
		panic(err)
	}
	dir := tree.DetermineDirectory(req)
	if e := dir.Ensure(); e != nil {
		panic(e)
	}
	if e := dir.ChangeTo(); e != nil {
		panic(e)
	}
}
