package strategy

import "github.com/file-duplicate-search/search/fileSign"

type Strategy interface {
	Run(fileSignCollection fileSign.FileSignCollection)
}
