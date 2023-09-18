package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt"
	"github.com/file-duplicate-search/search/strategy"
	"github.com/file-duplicate-search/search/utility"
)

var minSize int
var directoryToSearch string

func init() {
	flag.IntVar(&minSize, "s", -1, "Optional parameter used to specify minimum file size")
	flag.StringVar(&directoryToSearch, "d", "", "Directory to search for duplicates")
}

func main() {
	flag.Parse()

	fileSignCollection := fileSign.FileSignCollection{
		FileSigns: []fileSign.FileSign{},
	}

	err := filepath.WalkDir(directoryToSearch, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			md5Hash, size := createMD5HashFromFile(path)

			if minSize != -1 && size < minSize {
				return nil
			}

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
	chosenStrategy := prompt.SelectStrategyQuestion()

	var chosenStrategyObj strategy.Strategy

	if chosenStrategy == strategy.STRATEGY_DO_FOR_ALL {
		chosenStrategyObj = strategy.CreateDoForAllStrategy()
	} else if chosenStrategy == strategy.STRATEGY_GO_THROUGH_ONE_BY_ONE {
		chosenStrategyObj = strategy.CreateDoOneByOneStrategy()
	} else {
		panic("Not known strategy!")
	}

	chosenStrategyObj.Run()
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
