package main

import (
	"fmt"
	"errors"
)

func callbackInspect(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cnfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf(" I dont have any info on %v yet. Lets go catch one then ill be able to tell you everything about them! ", pokemonName)
	}


	fmt.Printf("Oh boy! Let me tell you all about %s\n", pokemon.Name)
	fmt.Printf("Pokemon ID: %v\n", pokemon.ID)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats{
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Type(s):\n")
	for _, typ := range pokemon.Types{
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	fmt.Printf("Abilities:\n")
	for _, ability := range pokemon.Abilities{
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	// comebcak and add this later
	//fmt.Printf("Here are the places we can find more of them:\n")
	//for _, location := range pokemon.LocationAreaEncounters{
	//	fmt.Printf(" - %s\n", ability.Ability.Name)
	//}

	return nil


}