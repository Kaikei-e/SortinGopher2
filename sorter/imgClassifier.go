package sorter

import (
	"SortinGopher2/cells"
	"fmt"
	"io/ioutil"
	"sync"
)

func ImgClassifier(ps []string, wg *sync.WaitGroup) error {
	defer wg.Done()

	var imgFs []cells.ImgFolder

	fmt.Println("Classifying started ...")
	for _, p := range ps {
		var imgF cells.ImgFolder

		list, searchErr := imgSearcher(p)
		if searchErr != nil {
			fmt.Errorf("failed to search images : %w", searchErr)
		}

		imgF.ImgPaths = list
		imgF.FolderPath = p

		imgFs = append(imgFs, imgF)
	}

	fmt.Println(imgFs)

	return nil
}

func imgSearcher(fp string) ([]string, error) {
	dir, readErr := ioutil.ReadDir(fp)
	if readErr != nil {
		return nil, fmt.Errorf("failed to search images : %w", readErr)
	}

	var imgList []string
	for _, fileInfo := range dir {
		img := fileInfo.Name()

		imgList = append(imgList, img)
	}

	return imgList, readErr
}
