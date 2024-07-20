package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/McFuddy2/Pokedex/internal/pokeapi"
	"strconv"
)

func callbackLoad(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no file name provided")
	}

	fileName := "SaveData/" + args[0] + ".csv"

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %s", err)
	}

	if len(records) < 2 {
		return fmt.Errorf("CSV file does not contain enough data")
	}

	// Extract headers
	headers := records[0]

	// Check if headers are valid
	expectedHeaders := []string{
		"Name",
		"ID",
		"BaseExperience",
		"Height",
		"Weight",
		"IsDefault",
		"LocationAreaEncounters",
		"Types",
		"Abilities",
	}
	for i, header := range expectedHeaders {
		if headers[i] != header {
			return fmt.Errorf("unexpected header: got %s, want %s", headers[i], header)
		}
	}

	// Process each record
	for _, record := range records[1:] {
		if len(record) != len(expectedHeaders) {
			return fmt.Errorf("record length does not match header length")
		}

		id, err := strconv.Atoi(record[1])
		if err != nil {
			return fmt.Errorf("invalid ID value: %s", record[1])
		}
		baseExp, err := strconv.Atoi(record[2])
		if err != nil {
			return fmt.Errorf("invalid BaseExperience value: %s", record[2])
		}
		height, err := strconv.Atoi(record[3])
		if err != nil {
			return fmt.Errorf("invalid Height value: %s", record[3])
		}
		weight, err := strconv.Atoi(record[4])
		if err != nil {
			return fmt.Errorf("invalid Weight value: %s", record[4])
		}
		isDefault, err := strconv.ParseBool(record[5])
		if err != nil {
			return fmt.Errorf("invalid IsDefault value: %s", record[5])
		}

		cnfg.caughtPokemon[record[0]] = pokeapi.Pokemon{
			Name:                  record[0],
			ID:                    id,
			BaseExperience:        baseExp,
			Height:                height,
			Weight:                weight,
			IsDefault:             isDefault,
			LocationAreaEncounters: record[6],
			Types: []struct {
				Slot int `json:"slot"`
				Type struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"type"`
			}{
				// Populate types from the types slice
			},
			Abilities: []struct {
				Ability struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"ability"`
				IsHidden bool `json:"is_hidden"`
				Slot     int  `json:"slot"`
			}{
				// Populate abilities from the abilities slice
			},
		}
	}

	fmt.Println("Data successfully loaded from CSV file!")

	return nil
}
