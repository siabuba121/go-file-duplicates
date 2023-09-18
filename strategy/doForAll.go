package strategy

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

const STRATEGY_DO_FOR_ALL = "Do action for all"

const LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST = "Leave first occurence and remove rest"
const REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG = "Remove all and copy one occurence to new catalog"

type DoForAllStrategy struct {
	name string
}

func CreateDoForAllStrategy() DoForAllStrategy {
	return DoForAllStrategy{
		name: STRATEGY_DO_FOR_ALL,
	}
}

func (concreteStrategy DoForAllStrategy) Run() {
	action := selectDoForAllAction()
	//TODO HERE ENDED
	fmt.Println(action)
}

func selectDoForAllAction() string {
	chosenAction := ""
	prompt := &survey.Select{
		Message: "What action should be done against all duplicates?",
		Options: []string{LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST, REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG},
	}
	survey.AskOne(prompt, &chosenAction)

	return chosenAction
}
