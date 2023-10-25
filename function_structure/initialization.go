package hangman

import "math/rand"

func (hangman *HangManData) setData() {
	hangman.Word = []rune{}
	hangman.ToFind = ""
	hangman.Attempts = 10
	hangman.HangmanPositions = -1
	hangman.ListUsed = []rune{}
	hangman.ListWord = []string{}
}

func (hang *HangManData) SetWord(dico []string) {
	randomIndex := rand.Intn(len(dico) - 1)
	hang.ToFind = dico[randomIndex]
	nbVisibleLetter := len(hang.ToFind)/2 - 1
	for range hang.ToFind {
		hang.Word = append(hang.Word, '_')
	}
	again := false
	var place []int

	for nbVisibleLetter > 0 {
		randomIndex = rand.Intn(len(hang.ToFind))
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
