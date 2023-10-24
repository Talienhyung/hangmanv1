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
	fmt.Print(string(tab))
}

// return user input
func input(s string, inputs string) string {
	fmt.Println(s)
	fmt.Scanln(&inputs)
	return inputs
}
