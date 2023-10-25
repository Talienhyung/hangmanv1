package hangman

import (
	"fmt"
	"os"
	"os/exec"
)

// Display a manual for the utilisation of argument
func Help() {
	filePath := "Ressources/doc.txt"

	pagerCommand := "less"

	cmd := exec.Command(pagerCommand, filePath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error when executing %s: %v\n", pagerCommand, err)
		os.Exit(1)
	}
}

// Display the rule of the game
func Rules() {
	filePath := "Ressources/rules.txt"

	pagerCommand := "less"

	cmd := exec.Command(pagerCommand, filePath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error when executing %s: %v\n", pagerCommand, err)
		os.Exit(1)
	}
}
