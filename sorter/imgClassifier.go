package sorter

import (
	"SortinGopher2/cells"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ImgClassifier(ps []string, wg *sync.WaitGroup) error {
	defer wg.Done()

	var imgFs []cells.ImgFolder

	fmt.Println("Classifying started ...")
	for _, p := range ps {
		var imgF cells.ImgFolder

		list, searchErr := imgSearcher(p)
		if searchErr != nil {
			return fmt.Errorf("failed to search images : %w", searchErr)
		}

		imgF.ImgPaths = list
		imgF.FolderPath = p

		imgFs = append(imgFs, imgF)
	}

	for _, imgf := range imgFs {
		err := classifier(imgf)
		if err != nil {
			return err
		}
	}

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

func classifier(imgF cells.ImgFolder) error {
	fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv")
	fmt.Println(imgF.ImgPaths)

	for i, ph := range imgF.ImgPaths {
		if !strings.Contains(ph, ".") {
			continue
		}

		fmt.Println(ph)
		p := filepath.Clean(ph)
		fmt.Println(p)
		e := filepath.Ext(p)
		s := strings.Split(ph, "-")
		t := time.Now().Nanosecond()

		dirName := s[0]
		dirPath := filepath.Join(imgF.FolderPath, "/", dirName)
		_, err := os.Stat(dirPath)
		if strings.Contains(dirPath, ".DS_Store") {
			continue
		}

		if os.IsExist(err) {
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

			err := os.Rename(ph, filepath.Join(dirPath, "/", dirName, "_", strconv.Itoa(t), "_", strconv.Itoa(i), e))
			if err != nil {
				return fmt.Errorf("file rename was failed: %w", err)
			}
		} else {
			fmt.Println("////////////////")
			fmt.Println(dirPath)
			createErr := os.MkdirAll(dirPath, 775)
			if createErr != nil {
				return fmt.Errorf("failed to create the directory : %w", createErr)
			}

			fmt.Println(ph)
			fmt.Println(filepath.Join(dirPath, "/"+dirName+"_"+strconv.Itoa(t)+"_"+strconv.Itoa(i)+e))
			err := os.Rename(ph, filepath.Join(dirPath, "/"+dirName+"_"+strconv.Itoa(t)+"_"+strconv.Itoa(i)+e))
			if err != nil {
				return fmt.Errorf("file rename was failed: %w", err)
			}
		}

	}

	fmt.Println("Classified images ...")

	return nil
}
