package hangman

type HangManData struct {
	Word             []rune   // Word composed of '_', ex: H_ll_
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int      // Number of attempts left
	HangmanPositions int      // Positions parsed in "hangman.txt" are stored
	ListWord         []string // List of words suggested by the user
	ListUsed         []rune   // List of letter sugested by the user
}

type Game struct {
	save       bool   // True if the --startWith (-sw) argument is given
	classic    bool   // True if the --classic (-c) argument is given
	ascii      bool   // True if the --ascii (-a) argument is given
	letter     bool   // True if the --letter (-l) argument is given
	saveFile   string // Name of the file given after --startWith (-sw) where the backup is stored
	letterFile string // Name of the file given after --letter (-l) where the ascii art is stored
	dico       string // First argument given, contains the name of the file containing the desired dictionary
}
