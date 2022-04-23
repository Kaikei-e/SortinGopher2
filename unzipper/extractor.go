package unzipper

import (
	"SortinGopher2/cells"
	"fmt"
	"io/fs"
	"path/filepath"
)

func Extractor(ps []string) ([]cells.ZipFolder, error) {
	var zfs []cells.ZipFolder

	for _, p := range ps {
		var zf cells.ZipFolder
		z, err := zipSearcher(p)
		if err != nil {
			return nil, fmt.Errorf("zip searching was failed by: %w", err)
		}

		zf.FolderPath = p
		zf.Zips = z

		zfs = append(zfs, zf)

	}

	return zfs, nil
}

func zipSearcher(folderPath string) ([]string, error) {
	var zipFiles []string

	fmt.Println("search will start at :", folderPath)

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

func extractor() {

}
