package fileSign

import "bytes"

type FileSignCollection struct {
	FileSigns []FileSign
}

func (FileSignCollection *FileSignCollection) DiplayResult() {
	for _, fileSign := range FileSignCollection.FileSigns {
		if fileSign.Occurencies > 1 {
			fileSign.DisplayInfo()
		}
	}
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
