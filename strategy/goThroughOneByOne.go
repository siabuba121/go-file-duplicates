package strategy

import (
	"fmt"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt"
)

type OneByOneStrategy struct {
	name string
}

func CreateDoOneByOneStrategy() OneByOneStrategy {
	return OneByOneStrategy{
		name: prompt.STRATEGY_GO_THROUGH_ONE_BY_ONE,
	}
}

func (concreteStrategy OneByOneStrategy) Run(fileSign.FileSignCollection) {
	fmt.Println("onebyone")
}
