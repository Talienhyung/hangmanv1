package hangman

import "fmt"

func (hang HangManData) displayHangmanClassic() {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	fmt.Println("")
	for i := 0; i <= 7; i++ {
		fmt.Println(hangMan[hang.HangmanPositions][i])
	}
}

func printRune(tab []rune) {
	for _, runes := range tab {
		fmt.Print(string(runes))
	}
	fmt.Print("\n")
}

func input(s string, inputs string) string {
	fmt.Println(s)
	fmt.Scanln(&inputs)
	return inputs
}
