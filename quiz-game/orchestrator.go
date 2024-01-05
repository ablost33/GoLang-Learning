package quiz_game

import "fmt"

func PlayGame() {
	score := 0

	filepath, err := chooseQuestionFile()
	if err != nil {
		panic(err)
	}

	questionsAndAnswers, err := readCSV(filepath)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(questionsAndAnswers); i++ {
		fmt.Println("Questions is: ", questionsAndAnswers[i][0])
		answer, err := parseAnswer()
		if err != nil {
			panic(err)
		}
		if answer == questionsAndAnswers[i][1] {
			score++
		}
	}
	fmt.Println("You got ", score, "/", len(questionsAndAnswers), " correct answers.")

}
