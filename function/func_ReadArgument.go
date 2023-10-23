package Hangman

import "os"

func Arg() Game {
	var game Game
	arguments := os.Args[1:]
	needFile := true
	for index, arg := range arguments {
		switch arg {
		case "--startWith", "-sw":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = true
			}
			game.save = true
			if len(arguments) > index+1 {
				game.saveFile = arguments[index+1]
			} else {
				os.Exit(3)
			}

		case "--classic", "-c":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = false
			}
			if game.ascii {
				os.Exit(4)
			} else {
				game.classic = true
			}
		case "--ascii", "-a":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = false
			}
			if game.classic {
				os.Exit(4)
			} else {
				game.ascii = true
			}
		case "--letterFile", "-lf":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = true
			}
			if game.classic {
				os.Exit(4)
			} else {
				game.letter = true
			}
			if len(arguments) > index+1 {
				game.letterFile = arguments[index+1]
			} else {
				os.Exit(3)
			}
		case "--rules", "-r":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = false
			}
			Rules()
			os.Exit(0)
		case "--help", "-h":
			if needFile && index != 0 {
				os.Exit(3)
			} else {
				needFile = false
			}
			Help()
			os.Exit(0)
		default:
			if needFile && index == 0 {
				game.dico = arguments[0]
			}
			if needFile {
				needFile = false
			} else {
				os.Exit(3)
			}
		}
	}
	return game
}
