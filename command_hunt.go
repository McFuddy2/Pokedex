package main

import (
	"fmt"
	"errors"
	"math/rand"
)



func callbackHunt(cnfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cnfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	if len(locationArea.PokemonEncounters) == 0 {
		return fmt.Errorf(" Looks like there are no pokemon to catch here! ")
	}


	randomNumbr := rand.Intn(len(locationArea.PokemonEncounters) + ((len(locationArea.PokemonEncounters)+4)/5) )
	
	if randomNumbr >= len(locationArea.PokemonEncounters){
		return errors.New(" Shoot! i couldnt find any pokemon lurking about! Lets try again! ")
	}
		fmt.Printf("Oh! I found a %v!\n", locationArea.PokemonEncounters[randomNumbr].Pokemon.Name)
		for _, myPokemon := range cnfg.caughtPokemon {
			if myPokemon.Name == (locationArea.PokemonEncounters[randomNumbr].Pokemon.Name) {
				return fmt.Errorf(" Shoot, %v is already our friend. Lets let him go and try to find a new friend! ", myPokemon.Name)
			}
		} 
		err = CatchPokemon(cnfg, (locationArea.PokemonEncounters[randomNumbr].Pokemon.Name))
		if err != nil{
			return err
		}
	return nil
}





func CatchPokemon(cnfg *config, pokemonName string) error {

	pokemon, err := cnfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randomNumbr := rand.Intn(pokemon.BaseExperience)

	if randomNumbr > threshold{
		if randomNumbr - threshold >= 200 {
			return fmt.Errorf("oof! it wasnt even close! %v got away! ", pokemonName)	
		} else if randomNumbr - threshold >= 100 {
			return fmt.Errorf("... Darn! %v Ran off! ", pokemonName)	
		} else if randomNumbr - threshold >= 50 {
			return fmt.Errorf("... ... Awe man, I thought we had it this time, but %v fled! ", pokemonName)	
		} else if randomNumbr - threshold >= 10 {
			return fmt.Errorf("... ... ... Oh It was so close! %v broke free! ", pokemonName)	
		}
	}

	cnfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("... ... ... %s was caught!\n", pokemonName)
	return nil


}