package strategy

import (
	"fmt"

	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt/messages"
)

type OneByOneStrategy struct {
	name string
}

func CreateDoOneByOneStrategy() OneByOneStrategy {
	return OneByOneStrategy{
		name: messages.STRATEGY_GO_THROUGH_ONE_BY_ONE,
	}
}

func (concreteStrategy OneByOneStrategy) Run(fileSign.FileSignCollection) {
	//TODO
	fmt.Println("onebyone")
}
