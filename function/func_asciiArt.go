package Hangman

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func (data *Game) DisplayAscii(x, y, version int, borderColor termbox.Attribute) {
	var ascii [95][9]string
	switch data.letterFile {
	case "shadow":
		ascii = readAscii("Ressources/Ascii_Letter/shadow.txt")
	case "standard":
		ascii = readAscii("Ressources/Ascii_Letter/standard.txt")
	case "thinkertoy":
		ascii = readAscii("Ressources/Ascii_Letter/thinkertoy.txt")
	default:
		os.Exit(3)
	}

	for i := 0; i <= 8; i++ {
		runes := []rune(ascii[version-32][i])
		for index, j := range runes {
			termbox.SetCell(x+index, y+i, j, borderColor, termbox.ColorDefault)
		}
	}
}

func readAscii(fichier string) [95][9]string {
	var ascii [95][9]string

	readFile, err := os.Open(fichier)
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile) // Creates a scanner to read the file.

	fileScanner.Split(bufio.ScanLines) // Divides the file into lines.

	i, j := 0, 0
	// Browse each line of the file.
	for fileScanner.Scan() {
		ascii[i][j] = fileScanner.Text()
		j++
		if j == 9 {
			i++
			j = 0
		}
	}

	readFile.Close()

	return ascii
}
