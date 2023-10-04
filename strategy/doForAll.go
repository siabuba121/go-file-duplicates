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

type DoForAllStrategy struct {
	name string
}

func CreateDoForAllStrategy() DoForAllStrategy {
	return DoForAllStrategy{
		name: messages.STRATEGY_DO_FOR_ALL,
	}
}

func (concreteStrategy DoForAllStrategy) Run(fileSignCollection fileSign.FileSignCollection) {
	action := prompt.SelectActionAgainstDuplicate()
	if action == messages.LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST {
		leaveFirstOccurenceAndRemoveRest(fileSignCollection)
	} else if action == messages.REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG {

		removeAllAndCopyOneOccurenceToNewCatalog(
			fileSignCollection,
			getCatalogName(),
		)
	}
}

func leaveFirstOccurenceAndRemoveRest(fileSignCollection fileSign.FileSignCollection) {
	for _, fileSign := range fileSignCollection.FileSigns {
		filePaths := fileSign.Path
		filePaths = filePaths[1:]
		for _, filePath := range filePaths {
			os.Remove(filePath)
		}
	}
}

func getCatalogName() string {
	catalogName := prompt.AskForCatalogName()
	return catalogName
}

func removeAllAndCopyOneOccurenceToNewCatalog(
	fileSignCollection fileSign.FileSignCollection,
	catalogName string,
) {
	if err := os.Mkdir(catalogName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	for _, fileSign := range fileSignCollection.FileSigns {
		if fileSign.Occurencies > 1 {
			firstPath := fileSign.Path[0]
			filename := filepath.Base(firstPath)
			utility.CopyFile(firstPath, fmt.Sprintf("%s/%s", catalogName, filename))
		}
		for _, filePath := range fileSign.Path {
			os.Remove(filePath)
		}
	}
}
