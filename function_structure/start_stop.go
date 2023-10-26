package hangman

import (
	"encoding/json"
	"os"
)

// Saves the party's progress, which is stored in the HangManData structure
func (data HangManData) Save(filename string) error {
	file, err := os.Create(filename) // Create a file
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil { // Save HangManData
		return err
	}

	return nil
}

// Load the party's progress, which is stored in filename, return a HangManData struct
func Load(filename string) (HangManData, error) {
	var data HangManData

	file, err := os.Open(filename) // Open filename
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil { // Load HangManData
		return data, err
	}

	return data, nil
}
