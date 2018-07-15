package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Nekroze/scd/tree"
)

var dryRun bool

func parseArgs() {
	flag.BoolVar(&dryRun, "dryrun", false, "If true no files or directories will be created")
	flag.Parse()
}

func Entrypoint() {
	parseArgs()
	if len(flag.Args()) != 1 {
		fmt.Println(os.Getenv("HOME"))
		return
	}
	pathInput := flag.Arg(0)
	req := tree.NewDirectoryRequest(pathInput)
	dir := tree.NewDirectory(req)
	if !dryRun {
		if e := dir.Ensure(); e != nil {
			panic(e)
		}
	}
	fmt.Println(dir.Path)
}
