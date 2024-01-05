package quiz_game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func chooseQuestionFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input the filepath to your question file: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(input, "\n", ""), nil
}

func parseAnswer() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(answer, "\n", ""), nil
}
