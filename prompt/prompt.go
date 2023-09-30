package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/file-duplicate-search/search/fileSign"
	"github.com/file-duplicate-search/search/prompt/messages"
	"github.com/file-duplicate-search/search/utility"
)

func getPossibelYesNoAnswers() []string {
	return []string{"Y", "N"}
}

func SelectStrategyQuestion() string {
	chosenStrategy := ""
	prompt := &survey.Select{
		Message: messages.QUESTION_CHOOSE_STRATEGY,
		Options: []string{
			messages.STRATEGY_DO_FOR_ALL,
			messages.STRATEGY_GO_THROUGH_ONE_BY_ONE,
		},
	}
	survey.AskOne(prompt, &chosenStrategy)
	fmt.Println("")

	return chosenStrategy
}

func AskIfDuplicatesShouldBeRemoved() bool {
	return displayYesNoPrompt(messages.SHALL_REMOVE_QUESTION)
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
		utility.LogError(messages.INVALID_ANSWER_MESSAGE)
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
		utility.LogQuestion(messages.QUESTION_NAME_OF_BACKUP_CATALOG)
		answer, _ = r.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer != "" {
			break
		}
		utility.LogError(messages.INVALID_ANSWER_MESSAGE)
		fmt.Println("")
	}

	return answer

}

func SelectDoForAllAction() string {
	chosenAction := ""
	prompt := &survey.Select{
		Message: messages.QUESTION_WHAT_TO_DO_WITH_DUPLICATES,
		Options: []string{
			messages.LEAVE_FIRST_OCCURENCE_AND_REMOVE_REST,
			messages.REMOVE_ALL_AND_COPY_ONE_OCCURENCE_TO_NEW_CATALOG,
		},
	}
	survey.AskOne(prompt, &chosenAction)
	fmt.Println("")

	return chosenAction
}

func PrintSummaryBasedOnFileSignCollection(fileSignCollection fileSign.FileSignCollection) {
	deletedFiles := 0
	spaceSaved := 0

	for _, fileSign := range fileSignCollection.FileSigns {
		deletedFiles += fileSign.Occurencies - 1
		spaceSaved += (fileSign.Occurencies - 1) * fileSign.Size
	}

	utility.LogSuccess(
		fmt.Sprintf(
			"%d files were deleted,\n%d kb of space saved.",
			deletedFiles,
			spaceSaved,
		),
	)
}
