package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var correctCount int
var questions = []string{"apple", "banana"}

func runQuestions(ch chan string){
	stdin := bufio.NewScanner(os.Stdin)
	questionIndex := 0
	for {
		if questionIndex == len(questions){break}
		expect := questions[questionIndex]
		fmt.Println(expect)
		stdin.Scan()
		if stdin.Text() == expect{
			fmt.Println("correct!")
			correctCount++
			questionIndex++
		} else {
			fmt.Println("Not correct! try again!")
		}
	}
	ch <- "Congrats! Perfect!"
}

func showResult(){
	fmt.Printf("Your score is %d\n", correctCount)
}
func main(){
	c := make(chan string)
	fmt.Println("start!")
	go runQuestions(c)
	select {
		case m:= <- c:
			fmt.Println(m)
			case <- time.After(10 * time.Second):
				fmt.Println("time up!")
	}
	showResult()
}
