package unzipper

import (
	"SortinGopher2/cells"
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

	for _, zf := range zfs {
		err := extractor(zf)
		if err != nil {
			return nil, fmt.Errorf("zip extracting was failed by: %w", err)
		}
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

func extractor(zf cells.ZipFolder) error {
	fmt.Println("Extracting is started")

	for _, p := range zf.Zips {
		fmt.Println(p)

		reader, err := zip.OpenReader(p)
		if err != nil {
			return err
		}

		for i, img := range reader.File {

			if img.FileInfo().IsDir() {
				ph := filepath.Join(zf.FolderPath, "/", img.Name)

				err := os.MkdirAll(ph, img.Mode())
				if err != nil && !os.IsExist(err) {
					return fmt.Errorf("failed to make dir : %w", err)
				}
			} else {
				buf := make([]byte, img.UncompressedSize64)
				raw, err := img.OpenRaw()
				if err != nil {
					return fmt.Errorf("failed to open the img file : %w", err)
				}

				_, readErr := io.ReadFull(raw, buf)
				if readErr != nil {
					return fmt.Errorf("failed to open the img file : %w", readErr)
				}

				t := time.Now().Nanosecond()

				strs := strings.Split(img.Name, "-")
				imgName := strs[0] + "-" + strconv.Itoa(i) + strconv.Itoa(t) + filepath.Ext(img.Name)

				path := filepath.Join(zf.FolderPath, "/", imgName)
				out, createErr := os.Create(path)
				if createErr != nil {
					return fmt.Errorf("failed to open the img file : %w", createErr)

				}

				writeErr := os.WriteFile(path, buf, img.Mode())
				if writeErr != nil {
					return fmt.Errorf("failed to write file : %w", writeErr)
				}

				closeErr := out.Close()
				if closeErr != nil {
					return fmt.Errorf("failed to close file : %w", closeErr)
				}

			}

		}

		readErr := reader.Close()
		if readErr != nil {
			return fmt.Errorf("reader closing failed : %w", readErr)
		}

	}

	fmt.Println("Extracted the zip.")

	return nil

}
