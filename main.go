package main

import (
	"SortinGopher2/unzipper"
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println("Please enter the path where the zip file is located ... ")

	args := os.Args
	var paths []string

	for _, arg := range args {
		p := path.Clean(arg)
		if path.IsAbs(p) {
			paths = append(paths, p)
		}
	}

	exts, err := unzipper.Extractor(paths)
	if err != nil {
		fmt.Errorf("zip extract was failed by : %w", err)
	}

	fmt.Println(exts)
}
