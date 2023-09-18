package prompt

import (
	"bufio"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/file-duplicate-search/search/strategy"
	"github.com/file-duplicate-search/search/utility"
)

const SHALL_REMOVE_QUESTION = "Do you want to remove duplicates (Y/N)"
const INVALID_ANSWER_MESSAGE = "Provided answer is not valid!"

func getPossibelYesNoAnswers() []string {
	return []string{"Y", "N"}
}

func SelectStrategyQuestion() string {
	chosenStrategy := ""
	prompt := &survey.Select{
		Message: "Which strategy should be used for found file duplicates?",
		Options: []string{strategy.STRATEGY_DO_FOR_ALL, strategy.STRATEGY_GO_THROUGH_ONE_BY_ONE},
	}
	survey.AskOne(prompt, &chosenStrategy)

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
