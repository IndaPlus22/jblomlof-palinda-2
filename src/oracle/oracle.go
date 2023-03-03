// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt) // added
	for {
		//fmt.Print(prompt) // sorry i cant stand it printing prompt everytime, redoing
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	dumpChannel := make(chan string)
	questions := make(chan string)
	go recieveQuestion(questions, dumpChannel)
	// TODO: Make prophecies.
	go createPropechy(dumpChannel)
	// TODO: Print answers.
	go dumpToSTDOUT(dumpChannel)

	return questions
}

func createPropechy(dumpChannel chan<- string) {
	// from below
	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"Java is great",
		"Rust is a game",
	}

	for {
		time.Sleep(time.Second * time.Duration(5+rand.Intn(6)))
		dumpChannel <- nonsense[rand.Intn(len(nonsense))]
	}
}

func dumpToSTDOUT(rec <-chan string) {
	for dump := range rec {
		//converting to array of chars
		// using method in https://www.tutorialkart.com/golang-tutorial/golang-convert-string-into-array-of-characters/
		chars := []byte(dump)
		for _, char := range chars {
			// humm a work around to print chars
			fmt.Printf("%c", char)
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Print("\n", prompt)
	}
}

// The task demands me to do this although it's (almost) useless.
func recieveQuestion(recCH <-chan string, dumpChannel chan<- string) {
	for question := range recCH {
		go answerQuestion(question, dumpChannel)
	}
}

// handle questions and dump out answers
func answerQuestion(question string, dumpChannel chan<- string) {
	// from the wiki https://en.wikipedia.org/wiki/Magic_8_Ball
	// but also modified
	// dont know why im using this as a reference
	eight_ball_answers := []string{
		"Yes",
		"Most likley",
		"As I see it, yes",
		"Maybe",
		"Cannot predict now",
		"Reply hazy, try again",
		"No",
		"My sources say no",
		"Very doubtful",
	}

	type standardPromptsAnswers struct {
		prompt string
		answer string
	}
	standardPrompts := []standardPromptsAnswers{
		//make prompt in lower and answer in normal
		{"what to do with lemons?", "When life gives you lemons. Make lemonade."},
		{"will i find love?", "You've already done that"},
		{"java or rust?", "... *disappointed*"},
	}
	foundStandard := false
	var ans string

	time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond) // 500 ms + x ms, 0 =< x < 500
	lowerCaseQuestion := strings.ToLower(question)
	for _, prompt := range standardPrompts {
		if prompt.prompt == lowerCaseQuestion {
			foundStandard = true
			ans = prompt.answer
			break
		}
	}
	if !foundStandard {
		ans = eight_ball_answers[rand.Intn(len(eight_ball_answers))]
	}
	dumpChannel <- ans

}

/*
Im not allowed to touch this and its not used smh.....

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.

	func prophecy(question string, answer chan<- string) {
		// Keep them waiting. Pythia, the original oracle at Delphi,
		// only gave prophecies on the seventh day of each month.
		time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

		// Find the longest word.
		longestWord := ""
		words := strings.Fields(question) // Fields extracts the words into a slice.
		for _, w := range words {
			if len(w) > len(longestWord) {
				longestWord = w
			}
		}

		// Cook up some pointless nonsense.
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
		}
		answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
	}
*/
func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
