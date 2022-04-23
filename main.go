package main

import (
	"SortinGopher2/sorter"
	"SortinGopher2/unzipper"
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)

	err := unzipper.Extractor(paths, &wg)
	if err != nil {
		fmt.Errorf("zip extract was failed by : %w", err)
	}

	wg.Wait()

	wg.Add(1)

	cErr := sorter.ImgClassifier(paths, &wg)
	if cErr != nil {
		fmt.Errorf("classifying was failed : %w", cErr)
	}

	wg.Wait()

	//fmt.Println(exts)
}
