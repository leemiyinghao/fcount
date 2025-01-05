package main

import (
	"flag"
	"fmt"

	humanize "github.com/dustin/go-humanize"

	"github.com/leemiyinghao/fcount/pkg/walk_directory"
)

func main() {
	dir := flag.String("dir", "", "directory to walk")
	flag.Parse()

	if *dir == "" {
		fmt.Println("Please provide a directory to walk")
		fmt.Println("Usage: fcount -dir <directory>")
		return
	}

	totalSize, totalFiles, err := walk_directory.WalkDirectory(*dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total size: %s\n", humanize.Bytes(totalSize))
	fmt.Printf("Total files: %s\n", humanize.Comma(int64(totalFiles)))
}
