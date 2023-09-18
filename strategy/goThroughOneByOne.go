package strategy

import "fmt"

const STRATEGY_GO_THROUGH_ONE_BY_ONE = "Go through one by one"

type OneByOneStrategy struct {
	name string
}

func CreateDoOneByOneStrategy() OneByOneStrategy {
	return OneByOneStrategy{
		name: STRATEGY_GO_THROUGH_ONE_BY_ONE,
	}
}

func (concreteStrategy OneByOneStrategy) Run() {
	fmt.Println("onebyone")
}
