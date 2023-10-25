package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This function reads the given ascii file and returns a [95][9]string containing the ascii art characters.
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

// This function reads the given hangman file and returns a [10][8]string containing the hangman position.
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

//########### Dictionary function ##################

// The readfile function returns an array of strings containing all the words in a dictionary
func readFile(fichier string) []string {
	var dictio []string

	readFile, err := os.Open(fichier)
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile) // Creates a scanner to read the file.

	fileScanner.Split(bufio.ScanLines) // Divides the file into lines.

	// Browse each line of the file.
	for fileScanner.Scan() {
		// Adds the text of the current line to the FindWord slice.
		dictio = append(dictio, fileScanner.Text())
	}

	readFile.Close()

	return dictio
}

// The listDictio function returns all files in the Dictinonary directory
func listDictio() []string {
	var listDico []string
	entries, err := os.ReadDir("Ressources/Dictionary/")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		listDico = append(listDico, e.Name())
	}
	return listDico
}

// The ReadAllDico function returns an array of strings containing all the words in the various files in the dictionary folder.
func readAllDico() []string {
	listDico := listDictio()
	var dico []string
	for i := 0; i < len(listDico); i++ {
		newDico := readFile("Ressources/Dictionary/" + listDico[i])
		dico = append(dico, newDico...)
	}
	return dico
}

// This function returns an array of words, depending on the file entered as a parameter only if the file exists, otherwise it uses the other files.
func ReadTheDico(file string) []string {
	listDico := listDictio()
	for _, j := range listDico {
		if file == j {
			Dico := readFile("Ressources/Dictionary/" + file)
			return Dico
		}
	}
	return readAllDico()
}
