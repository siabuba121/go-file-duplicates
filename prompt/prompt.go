package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/file-duplicate-search/search/utility"
)

const SHALL_REMOVE_QUESTION = "Do you want to remove duplicates (Y/N)"
const INVALID_ANSWER_MESSAGE = "Provided answer is not valid!"

const LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST = "Leave first occurence and remove rest"
const REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG = "Remove all and copy one occurence to new catalog"

const STRATEGY_GO_THROUGH_ONE_BY_ONE = "Go through one by one"
const STRATEGY_DO_FOR_ALL = "Do action for all"

func getPossibelYesNoAnswers() []string {
	return []string{"Y", "N"}
}

func SelectStrategyQuestion() string {
	chosenStrategy := ""
	prompt := &survey.Select{
		Message: "Which strategy should be used for found file duplicates?",
		Options: []string{STRATEGY_DO_FOR_ALL, STRATEGY_GO_THROUGH_ONE_BY_ONE},
	}
	survey.AskOne(prompt, &chosenStrategy)
	fmt.Println("")

	return chosenStrategy
}

func AskIfDuplicatesShouldBeRemoved() bool {
	return displayYesNoPrompt(SHALL_REMOVE_QUESTION)
}

func displayYesNoPrompt(question string) bool {
	var answer string
	r := bufio.NewReader(os.Stdin)
	for {
		utility.LogQuestion(question + " ")
		answer, _ = r.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if isAnswerValid(answer, getPossibelYesNoAnswers()) {
			break
		}
		utility.LogError(INVALID_ANSWER_MESSAGE)
		fmt.Println("")
	}

	if answer == "Y" {
		return true
	} else {
		return false
	}
}

func isAnswerValid(answer string, validAnswers []string) bool {
	for _, possibleAnswer := range validAnswers {
		if possibleAnswer == answer {
			return true
		}
	}

	return false
}

func AskForCatalogName() string {
	var answer string
	r := bufio.NewReader(os.Stdin)
	for {
		utility.LogQuestion("What name should backup catalog have?")
		answer, _ = r.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer != "" {
			break
		}
		utility.LogError(INVALID_ANSWER_MESSAGE)
		fmt.Println("")
	}

	return answer

}

func SelectDoForAllAction() string {
	chosenAction := ""
	prompt := &survey.Select{
		Message: "What action should be done against all duplicates?",
		Options: []string{LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST, REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG},
	}
	survey.AskOne(prompt, &chosenAction)
	fmt.Println("")

	return chosenAction
}
