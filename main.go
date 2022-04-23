package main

import (
	"SortinGopher2/unzipper"
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Println("Please enter the path where the zip file is located ... ")

	var sc = bufio.NewScanner(os.Stdin)

	sc.Scan()

	inputs := sc.Text()
	args := strings.Split(inputs, " ")

	var paths []string

	for _, arg := range args {
		p := path.Clean(arg)
		if path.IsAbs(p) {
			paths = append(paths, p)
		}
	}

	_, err := unzipper.Extractor(paths)
	if err != nil {
		fmt.Errorf("zip extract was failed by : %w", err)
	}

	//fmt.Println(exts)
}
