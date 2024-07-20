package main

import (
	"fmt"
	"errors"
	"os"
	"encoding/csv"
)



func callbackSave(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no file save name provided")
	}

	fileSaveName := "SaveData/" + args[0] + ".csv"


	file, err := os.OpenFile(fileSaveName, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to  open or create file: %s", err)
			}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		headers := []string{
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
	err = writer.Write(headers)
	if err != nil {
		return fmt.Errorf("failed to write headers to CSV file: %s", err)
	}

	for _, pokemon := range cnfg.caughtPokemon {
		types := []string{}
		for _, t := range pokemon.Types {
			types = append(types, t.Type.Name)
		}

		abilities := []string{}
		for _, a := range pokemon.Abilities {
			abilities = append(abilities, a.Ability.Name)
		}

		record := []string{
			pokemon.Name,
			fmt.Sprintf("%d", pokemon.ID),
			fmt.Sprintf("%d", pokemon.BaseExperience),
			fmt.Sprintf("%d", pokemon.Height),
			fmt.Sprintf("%d", pokemon.Weight),
			fmt.Sprintf("%t", pokemon.IsDefault),
			pokemon.LocationAreaEncounters,
			fmt.Sprintf("%s", types),
			fmt.Sprintf("%s", abilities),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record to CSV file: %s", err)
		}
	}
	fmt.Println("Data successfully written to CSV file!")

	return nil
	}
