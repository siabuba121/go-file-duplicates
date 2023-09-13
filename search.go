package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt"
	"github.com/file-duplicate-search/search/utility"
)

func main() {
	dirToSearch := os.Args[1:][0]
	fileSignCollection := fileSign.FileSignCollection{
		FileSigns: []fileSign.FileSign{},
	}

	err := filepath.WalkDir(dirToSearch, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			md5Hash, size := createMD5HashFromFile(path)
			fileExists := fileSignCollection.DoesFileExistInResult(md5Hash)

			if !fileExists {
				processedFile := fileSign.FileSign{
					Path:        []string{path},
					Size:        size,
					Hash:        md5Hash,
					Occurencies: 1,
				}
				fileSignCollection.AddFileSign(processedFile)
			} else {
				fileSignCollection.AddOccurency(md5Hash, path)
			}
		}

		return nil
	})
	if err != nil {
		utility.LogError(
			fmt.Sprintf("Error reading directory: %s", err))
	}

	fileSignCollection.DiplayResult()
	prompt.AskIfDuplicatesShouldBeRemoved()
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
