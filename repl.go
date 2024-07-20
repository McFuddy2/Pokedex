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
	fmt.Println("Go ahead and type a command! if you want a list of options try 'help'.")

	for {
		fmt.Print("\n \n pokedex > ")
		
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
		"hunt":{
			name: "hunt {location_area}",
			description: "We can try our luck at finding a new friend to add to our collection!",
			callback: callbackHunt,
		},
		"inspect":{
			name: "inspect {pokemon_name}",
			description: "Once we have caught a Pokemon I can tell you all kinds of eciting information about it!",
			callback: callbackInspect,
		},
		"catalog":{
			name: "catalog",
			description: "I can show you a list of Pokemon we have caught and that you can now inspect!",
			callback: callbackCatalog,
		},
		"save":{
			name: "save {file_name}",
			description: "I can save all the pokemon you have caught so that you can return and keep hunting!",
			callback: callbackSave,
		},
		"load":{
			name: "load {file_name}",
			description: "Oh! I might have a list of all the Pokemon you have caught before. We can start from there!",
			callback: callbackLoad,
		},
	}
}

func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)
	words := strings.Fields(loweredString)
	return words
}