package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type FileSign struct {
	path        []string
	size        int
	hash        []byte
	occurencies int
}

func main() {
	dirToSearch := os.Args[1:][0]
	currentMD5Map := []FileSign{}
	err := filepath.WalkDir(dirToSearch, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			hash := md5.New()
			_, _ = io.Copy(hash, file)
			md5hash := hash.Sum(nil)
			fileExists := doesFileExistInResult(currentMD5Map, md5hash)

			if !fileExists {
				stats, _ := file.Stat()
				fileSign := FileSign{
					path:        []string{path},
					size:        int(stats.Size()),
					hash:        hash.Sum(nil),
					occurencies: 1,
				}
				currentMD5Map = append(currentMD5Map, fileSign)
				fmt.Println("created")
			} else {
				addOccurency(currentMD5Map, md5hash, path)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}
	diplayResult(currentMD5Map)
}

func diplayResult(result []FileSign) {
	for _, element := range result {
		fmt.Println(getStringRepresentationOfMD5Sum(element.hash))
		fmt.Println(element.size)
		fmt.Printf("%d times! \n", element.occurencies)
		fmt.Println("In paths:")
		for index, element := range element.path {
			fmt.Printf("\n%d %s", index, element)
		}
		savedBytes := (element.occurencies - 1) * element.size
		fmt.Printf("\nYou can save %d bytes\n", savedBytes)
	}
}

func getStringRepresentationOfMD5Sum(sum []byte) string {
	return hex.EncodeToString(sum[:])
}

func doesFileExistInResult(actualResult []FileSign, hash []byte) bool {
	for _, element := range actualResult {
		if bytes.Equal(element.hash, hash) {
			return true
		}
	}

	return false
}

func addOccurency(actualResult []FileSign, hash []byte, path string) {
	for index, element := range actualResult {
		if bytes.Equal(element.hash, hash) {
			fmt.Println("added")
			actualResult[index].occurencies += 1
			actualResult[index].path = append(actualResult[index].path, path)
		}
	}
}
