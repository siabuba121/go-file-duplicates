package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/utility"
)

func main() {
	dirToSearch := os.Args[1:][0]
	currentMD5Map := []fileSign.FileSign{}

	err := filepath.WalkDir(dirToSearch, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			md5Hash, size := createMD5HashFromFile(path)
			fileExists := doesFileExistInResult(currentMD5Map, md5Hash)

			if !fileExists {
				processedFile := fileSign.FileSign{
					Path:        []string{path},
					Size:        size,
					Hash:        md5Hash,
					Occurencies: 1,
				}
				currentMD5Map = append(currentMD5Map, processedFile)
			} else {
				addOccurency(currentMD5Map, md5Hash, path)
			}
		}

		return nil
	})
	if err != nil {
		utility.LogError(
			fmt.Sprintf("Error reading directory: %s", err))
	}
	diplayResult(currentMD5Map)
}

func diplayResult(result []fileSign.FileSign) {
	for _, element := range result {
		if element.Occurencies > 1 {
			element.DisplayInfo()
		}
	}
}

func createMD5HashFromFile(filePath string) ([]byte, int) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hash := md5.New()
	_, _ = io.Copy(hash, file)
	fileStats, _ := file.Stat()
	fileSize := fileStats.Size()

	return hash.Sum(nil), int(fileSize)
}

func doesFileExistInResult(actualResult []fileSign.FileSign, hash []byte) bool {
	for _, element := range actualResult {
		if bytes.Equal(element.Hash, hash) {
			return true
		}
	}

	return false
}

func addOccurency(actualResult []fileSign.FileSign, hash []byte, path string) {
	for index, element := range actualResult {
		if element.IsHashEqual(hash) {
			actualResult[index].AddPath(path)
		}
	}
}
