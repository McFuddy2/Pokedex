package main


import (
	"fmt"
	"errors"
)

func callbackExplore(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cnfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Woah look at all the Pokemon you can find at %v:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters{
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}