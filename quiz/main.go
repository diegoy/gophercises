package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("file", "problems.csv", "file name")
	flag.Parse()

	csvReader := openCsv(*fileName)
	inputReader := bufio.NewReader(os.Stdin)

	rightAnswers, total := 0, 0

	for {
		line, err := csvReader.Read()
		if err != nil && err.Error() == "EOF" {
			fmt.Printf("-----------------------\nFinished!\n-----------------------\n")
			fmt.Printf("Total questions: %d\n", total)
			fmt.Printf("Correct answers: %d\n", rightAnswers)
			fmt.Printf("Accuracy: %.2f%%\n", float64(rightAnswers)/float64(total)*100)
			fmt.Printf("-----------------------\n")
			break
		}
		question, answer := line[0], line[1]
		fmt.Printf("%s = ", question)

		userAnswer, _ := inputReader.ReadString('\n')
		userAnswer = strings.TrimSuffix(userAnswer, "\n")

		if strings.Compare(answer, userAnswer) == 0 {
			rightAnswers++
		}
		total++
	}
}

func openCsv(fileName string) *csv.Reader {
	file, err := os.Open(fileName)

	if err != nil {
		println("Oops alguma coisa deu errada.")
		panic(err.Error())
	}

	return csv.NewReader(file)
}
