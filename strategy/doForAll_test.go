package strategy

import (
	"log"
	"os"
	"testing"

	"github.com/file-duplicate-search/search/fileSign"
)

// You can use testing.T, if you want to test the code without benchmarking
func teardown() {
	os.RemoveAll("./testDir")
	os.Remove("./file1")
	os.Remove("./file2")
	os.Remove("./file3")
}

func TestRemoveAllAndCopyOneOccureneToNewCatalog(t *testing.T) {
	defer teardown()

	content := "ala ma kota"
	catalogName := "testDir"
	fileSignCollection := createExampleFileSignCollection(content)

	removeAllAndCopyOneOccurenceToNewCatalog(
		fileSignCollection,
		catalogName,
	)

	_, error := os.Stat("./testDir/file1")
	if os.IsNotExist(error) {
		t.Fatalf("File does not exits")
	}

	data, _ := os.ReadFile("./testDir/file1")
	if string(data) != content {
		t.Fatalf("Content does not match")
	}
}

func TestLeaveFirstOccurenceAndRemoveRest(t *testing.T) {
	defer teardown()

	content := "marek ma psa"
	fileSignCollection := createExampleFileSignCollection(content)
	leaveFirstOccurenceAndRemoveRest(fileSignCollection)

	_, error := os.Stat("./file1")
	if os.IsNotExist(error) {
		t.Fatalf("File 1 does exits")
	}

	data, _ := os.ReadFile("./file1")
	if string(data) != content {
		t.Fatalf("Content does not match")
	}

	_, error = os.Stat("./file2")
	if os.IsExist(error) {
		t.Fatalf("File does not exits")
	}
	_, error = os.Stat("./file3")
	if os.IsExist(error) {
		t.Fatalf("File does not exits")
	}
}

func createExampleFileSignCollection(content string) fileSign.FileSignCollection {
	createFileWithContent("./file1", content)
	createFileWithContent("./file2", content)
	createFileWithContent("./file3", content)

	fileSignCollection := fileSign.FileSignCollection{
		FileSigns: []fileSign.FileSign{
			{
				Path: []string{
					"./file1",
					"./file2",
					"./file3",
				},
				Size:        10,
				Hash:        []byte{},
				Occurencies: 3,
			},
		},
	}

	return fileSignCollection
}

func createFileWithContent(filePath string, content string) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
