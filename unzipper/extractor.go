package unzipper

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Extractor(ps []string) ([][]string, error) {
	var zipPaths [][]string

	for _, p := range ps {
		z, err := zipSearcher(p)
		if err != nil {
			return nil, fmt.Errorf("zip searching was failed by: %w", err)
		}

		zipPaths = append(zipPaths, z)

	}

	return zipPaths, nil
}

func zipSearcher(folderPath string) ([]string, error) {
	var zipFiles []string

	fmt.Printf("search will start at %v", folderPath)

	err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("searching dir made the error: %w", err)
		}
		if d.IsDir() {
			return nil
		}

		if matched, err := filepath.Match("*.zip", filepath.Base(path)); err != nil {
			return fmt.Errorf("zip matching occuered the error : %w", err)
		} else if matched {
			zipFiles = append(zipFiles, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking dir occuered the error : %w", err)
	}

	return zipFiles, nil
}
