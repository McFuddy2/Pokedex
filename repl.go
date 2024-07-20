package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func startRepl(cnfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, and welcome to your Pokedex! My name is Puffluff your AI assistant! but you can call me Puff!")

	for {
		fmt.Print("pokedex > ")
		
		scanner.Scan()
		
		input := scanner.Text()

		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1{
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Huh?? i dont know what you want... Maybe try 'help'    (you typed: %v)\n", input)
			continue
		}

		err := command.callback(cnfg, args...)
		if err != nil{
			fmt.Println(err)
		}
		
		/*
		case "map":
			locations, newConfig, err := fetchLocations(config, "next")
			if err != nil {
				fmt.Println("Error fetching locations:", err)
			} else {
				displayLocations(locations)
				config = newConfig
			}
		case "mapb":
			locations, newConfig, err := fetchLocations(config, "previous")
			if err != nil {
				fmt.Println("Error fetching locations:", err)
			} else {
				displayLocations(locations)
				config = newConfig
			}
			/*
		case "explore":
			fmt.Print("Enter location area name: ")
			if scanner.Scan() {
				areaName := scanner.Text()
				exploreCommand(areaName)
			}
			*/ 

		
	

	}
}



type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays this 'help' message",
			callback: callbackHelp,
		},
		"exit":{
			name: "exit",
			description: "I go to sleep and shut down your Pokedex!",
			callback: callbackExit,
		},
		"q":{
			name: "q",
			description: "I go to sleep and shut down your Pokedex!",
			callback: callbackExit,
		},
		"joke":{
			name: "joke",
			description: "I will tell you a pokemon joke that's sure to leave you Weezing",
			callback: callbackJoke,
		},
		"map":{
			name: "map",
			description: "I can pull up a list of the next 20 locations on our map!",
			callback: callbackMap,
		},
		"mapb":{
			name: "mapb",
			description: "I can pull up a list of the previous 20 locations on our map!",
			callback: callbackMapb,
		},
		"explore":{
			name: "explore {location_area}",
			description: "I will check out the area and tell you what kinds of pokemon you can find here!",
			callback: callbackExplore,
		},
	}
}

func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)
	words := strings.Fields(loweredString)
	return words
}