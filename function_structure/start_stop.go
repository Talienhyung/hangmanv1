package hangman

import (
	"encoding/json"
	"os"
)

func (data HangManData) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func Load(filename string) (HangManData, error) {
	var data HangManData

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
