package Hangman

import (
	"math/rand"
)

type HangManData struct {
	Word             []rune // Word composed of '_', ex: H_ll_
	ToFind           string // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int    // Number of attempts left
	HangmanPositions int    // It can be the array where the positions parsed in "hangman.txt" are stored
	ListWord         []string
	ListUsed         []rune
}

type Game struct {
	save       bool
	classic    bool
	ascii      bool
	letter     bool
	saveFile   string
	letterFile string
	dico       string
}

func (hangman *HangManData) setData() {
	hangman.Word = []rune{}
	hangman.ToFind = ""
	hangman.Attempts = 10
	hangman.HangmanPositions = -1
	hangman.ListUsed = []rune{}
	hangman.ListWord = []string{}
}

func (hang *HangManData) SetWord(dico []string) {
	randomIndex := random(len(dico) - 1)
	hang.ToFind = dico[randomIndex]
	nbVisibleLetter := len(hang.ToFind)/2 - 1
	for range hang.ToFind {
		hang.Word = append(hang.Word, '_')
	}
	again := false
	var place []int

	for nbVisibleLetter > 0 {
		randomIndex = random(len(hang.ToFind))
		for _, j := range place {
			if j == randomIndex {
				again = true
			}
		}
		if !again {
			place = append(place, randomIndex)
			nbVisibleLetter--
		} else {
			again = false
		}
	}
	WordRune := []rune(hang.ToFind)
	for _, index := range place {
		hang.Word[index] = WordRune[index]
		hang.LetterInWord(WordRune[index])
		hang.UsedLetter(WordRune[index])
	}

}

func random(x int) int {
	return rand.Intn(x)
}
