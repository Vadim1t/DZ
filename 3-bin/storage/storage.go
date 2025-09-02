package storage

import (
	"encoding/json"
	"os"

	"github.com/Vadim1t/DZ/3-bin/bins"
)

func SaveBins(filename string, binsList []bins.Bin) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // красивое форматирование
	return encoder.Encode(binsList)
}

func LoadBins(filename string) ([]bins.Bin, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var binsList []bins.Bin
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&binsList); err != nil {
		return nil, err
	}
	return binsList, nil
}
