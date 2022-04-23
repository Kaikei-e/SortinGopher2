package main

import (
	"SortinGopher2/sorter"
	"SortinGopher2/unzipper"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Please enter the path where the zip file is located ... ")

	var sc = bufio.NewScanner(os.Stdin)

	sc.Split(bufio.ScanWords)
	sc.Scan()

	inputs := sc.Text()
	fmt.Println(inputs)

	args := strings.Split(inputs, " ")

	fmt.Println("Args: ")
	fmt.Println(args)

	var paths []string

	for _, arg := range args {
		//p := path.Clean(arg)
		//if path.IsAbs(arg) {
		paths = append(paths, arg)
		//}
	}

	fmt.Println(paths)
	var wg sync.WaitGroup
	wg.Add(1)

	go unzipper.Extractor(paths, &wg)

	wg.Wait()

	wg.Add(1)

	cErr := sorter.ImgClassifier(paths, &wg)
	if cErr != nil {
		fmt.Errorf("classifying was failed : %w", cErr)
	}

	wg.Wait()

	//fmt.Println(exts)
}
