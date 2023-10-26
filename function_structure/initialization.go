package hangman

import "math/rand"

// Set HangMan's first Data
func (hangman *HangManData) setData() {
	hangman.Word = []rune{}
	hangman.ToFind = ""
	hangman.Attempts = 10
	hangman.HangmanPositions = -1
	hangman.ListUsed = []rune{}
	hangman.ListWord = []string{}
}

// Set Word and ToFind for HangManData
func (hang *HangManData) SetWord(dico []string) {
	// Find a random word
	randomIndex := rand.Intn(len(dico) - 1)
	hang.ToFind = dico[randomIndex]

	nbVisibleLetter := len(hang.ToFind)/2 - 1 // Set the number of letters that will be visible

	for range hang.ToFind { // Set Word
		hang.Word = append(hang.Word, '_')
	}

	again := false
	var place []int
	for nbVisibleLetter > 0 { // Finds randomly different nbVisibleLetter in ToFind
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
	for _, index := range place { // Add the different letter into Word
		hang.Word[index] = WordRune[index]
		hang.LetterInWord(WordRune[index])
		hang.UsedLetter(WordRune[index])
	}
}
