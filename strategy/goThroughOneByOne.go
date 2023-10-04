package strategy

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt"
	"github.com/file-duplicate-search/search/prompt/messages"
	"github.com/file-duplicate-search/search/utility"
)

type OneByOneStrategy struct {
	name string
}

func CreateDoOneByOneStrategy() OneByOneStrategy {
	return OneByOneStrategy{
		name: messages.STRATEGY_GO_THROUGH_ONE_BY_ONE,
	}
}

func (concreteStrategy OneByOneStrategy) Run(fileSignCollection fileSign.FileSignCollection) {
	for _, fileSign := range fileSignCollection.FileSigns {
		fileSign.DisplayInfo()
		action := prompt.SelectActionAgainstDuplicate()
		if action == messages.LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST {
			leaveFirstOccurenceAndRemoveRestForSingleDuplicate(fileSign)
		} else if action == messages.REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG {
			removeAllAndCopyOneOccurenceToNewCatalogForSingleDuplicate(
				fileSign,
				getCatalogName(),
			)
		}

		utility.LogSuccess("Duplicate solved!")
	}
}

func leaveFirstOccurenceAndRemoveRestForSingleDuplicate(fileSign fileSign.FileSign) {
	filePaths := fileSign.Path
	filePaths = filePaths[1:]
	for _, filePath := range filePaths {
		os.Remove(filePath)
	}
}

func removeAllAndCopyOneOccurenceToNewCatalogForSingleDuplicate(
	fileSign fileSign.FileSign,
	catalogName string,
) {
	if err := os.Mkdir(catalogName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if fileSign.Occurencies > 1 {
		firstPath := fileSign.Path[0]
		filename := filepath.Base(firstPath)
		utility.CopyFile(firstPath, fmt.Sprintf("%s/%s", catalogName, filename))
	}
	for _, filePath := range fileSign.Path {
		os.Remove(filePath)
	}
}
