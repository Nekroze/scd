package cli

import (
	"fmt"
	"os"

	"github.com/Nekroze/scd/tree"
)

func Entrypoint() {
	if len(os.Args) != 2 {
		fmt.Println(os.Getenv("HOME"))
		return
	}
	pathInput := os.Args[1]
	req := tree.NewDirectoryRequest(pathInput)
	dir := tree.NewDirectory(req)
	if e := dir.Ensure(); e != nil {
		panic(e)
	}
	fmt.Println(dir.Path)
}
