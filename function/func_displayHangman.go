package Hangman

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func DisplayHangman(x, y, version int, borderColor termbox.Attribute) {
	hangMan := readHang("Ressources/HangMan_Position/hangman.txt")
	for i := 0; i <= 7; i++ {
		runes := []rune(hangMan[version][i])
		for index, j := range runes {
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
	}
}

func readHang(fichier string) [10][8]string {
	var hangman [10][8]string

	readFile, err := os.Open(fichier)
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile) // Creates a scanner to read the file.

	fileScanner.Split(bufio.ScanLines) // Divides the file into lines.

	i, j := 0, 0
	// Browse each line of the file.
	for fileScanner.Scan() {
		// Adds the text of the current line to the FindWord slice.
		hangman[i][j] = fileScanner.Text()
		j++
		if j == 8 {
			i++
			j = 0
		}
	}

	readFile.Close()

	return hangman
}
