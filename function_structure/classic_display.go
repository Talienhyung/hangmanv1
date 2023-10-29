package hangman

import "fmt"

// Displays the hangman in the terminal
func (hang HangManData) displayHangmanClassic() {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	fmt.Println("")
	for i := 0; i <= 7; i++ {
		fmt.Println(hangMan[hang.HangmanPositions][i])
	}
}

// displays the rune array given as a parameter in the terminal
func printRune(tab []rune) {
	for _, runes := range tab {
		fmt.Print(string(runes))
		fmt.Print(" ")
	}
	fmt.Println("")
}

// return user input
func input(s string, inputs string) string {
	fmt.Print(s)
	fmt.Scanln(&inputs)
	return inputs
}
