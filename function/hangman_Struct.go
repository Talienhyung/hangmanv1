package Hangman

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
}
