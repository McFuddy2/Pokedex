package main

import "fmt"

func callbackHelp(cnfg *config, args ...string) error {
	fmt.Println("Looks like you are looking for some help! Let me find a list of some commands you can try!")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands{
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}