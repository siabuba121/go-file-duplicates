package fileSign

import (
	"bytes"
	"fmt"

	messages "github.com/file-duplicate-search/search/prompt/messages"
	"github.com/file-duplicate-search/search/utility"
)

type FileSignCollection struct {
	FileSigns []FileSign
}

func (FileSignCollection FileSignCollection) IsThereAnyDuplicate() bool {
	for _, fileSign := range FileSignCollection.FileSigns {
		if fileSign.Occurencies != 1 {
			return true
		}
	}

	return false
}

func (FileSignCollection *FileSignCollection) DiplayResult() {

	utility.LogInfo(fmt.Sprintf(messages.FOUND_X_DUPLICATES, len(FileSignCollection.FileSigns)))
	for _, fileSign := range FileSignCollection.FileSigns {
		if fileSign.Occurencies > 1 {
			for _, path := range fileSign.Path {
				utility.LogInfo(fmt.Sprintf("%s \n", path))
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

func (FileSignCollection *FileSignCollection) DoesFileExistInResult(hash []byte) bool {
	for _, fileSign := range FileSignCollection.FileSigns {
		if bytes.Equal(fileSign.Hash, hash) {
			return true
		}
	}

	return false
}

func (FileSignCollection *FileSignCollection) AddFileSign(fileSign FileSign) {
	FileSignCollection.FileSigns = append(FileSignCollection.FileSigns, fileSign)
}

func (FileSignCollection *FileSignCollection) AddOccurency(hash []byte, path string) {
	for index, element := range FileSignCollection.FileSigns {
		if element.IsHashEqual(hash) {
			FileSignCollection.FileSigns[index].AddPath(path)
		}
	}
}
