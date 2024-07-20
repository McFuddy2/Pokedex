package main


import (
	"fmt"
)

func callbackMap(cnfg *config, args ...string) error {
	resp, err := cnfg.pokeapiClient.ListLocationAreas(cnfg.NextLocationAreaURL)
	if err != nil{
		return err
	}

	fmt.Println("Locations Area:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cnfg.NextLocationAreaURL = resp.Next
	cnfg.PreviousLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(cnfg *config, args ...string) error {
	if cnfg.PreviousLocationAreaURL == nil {
		fmt.Println("Silly Swanna! There is no previous page! try again!")
		return nil
	}
	resp, err := cnfg.pokeapiClient.ListLocationAreas(cnfg.PreviousLocationAreaURL)
	if err != nil{
		return err
	}

	fmt.Println("Locations Area:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cnfg.NextLocationAreaURL = resp.Next
	cnfg.PreviousLocationAreaURL = resp.Previous

	return nil
}