package main

import (
	"os"
	"fmt"
)

func callbackExit(cnfg *config, args ...string) error {
	fmt.Println("Bye Bye for now! I'll CATCH you later!")
	os.Exit(0)
	return nil
}