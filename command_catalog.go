package main

import (
	"fmt"
)

func callbackCatalog(cnfg *config, args ...string) error {
	
	fmt.Println("Certainly! Here is a list of all the pokemon we have caught!")
	if len(cnfg.caughtPokemon) == 0{
		return fmt.Errorf(" Hmmm, well i looked all over and it doesnt look like we have caught any yet... lets go out and get some! ")	
	} else {	
		for _, pokemon := range cnfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
		}
	}
	return nil


}