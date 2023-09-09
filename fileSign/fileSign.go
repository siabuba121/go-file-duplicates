package fileSign

import (
	"bytes"
	"fmt"

	"github.com/file-duplicate-search/search/utility"
)

type FileSign struct {
	Path        []string
	Size        int
	Hash        []byte
	Occurencies int
}

func (fileSign FileSign) IsHashEqual(hashToCheck []byte) bool {
	return bytes.Equal(fileSign.Hash, hashToCheck)
}

func (fileSign *FileSign) AddPath(path string) {
	fileSign.Path = append(fileSign.Path, path)
	fileSign.Occurencies++
}

func (fileSign FileSign) DisplayInfo() {
	utility.LogInfo(
		fmt.Sprintf(
			"\nFile: %s \n"+
				"is duplicated %d times \n"+
				"Removing duplicates will reduce space by %.6f MB\n",
			fileSign.Path[0],
			fileSign.Occurencies,
			float64((fileSign.Occurencies-1))*float64(fileSign.Size)/1024.0/1024.0))
}
