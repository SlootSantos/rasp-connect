package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/SlootSantos/rasp-connect/command"
	"github.com/SlootSantos/rasp-connect/config"
	"github.com/SlootSantos/rasp-connect/device"
)

const cmdArgIdx = 1
const deviceArgIdx = 2
const fullArgLength = 3
const minArgLength = 2

func main() {
	command, piDevice, err := parseArguments()
	if err != nil {
		grace(err)
	}

	err = command(piDevice)
	if err != nil {
		grace(err)
	}
}

func parseArguments() (command.Executable, *device.Pi, error) {
	var err error
	var piDevice = &device.Pi{}

	if len(os.Args) < minArgLength {
		err = errors.New("Too few arguments")
		return nil, piDevice, err
	}

	if len(os.Args) >= fullArgLength {
		deviceArg := os.Args[deviceArgIdx]
		piDevice = config.Pis[deviceArg]
	}

	commandArg := os.Args[cmdArgIdx]
	executableCmd := command.CommandMap[commandArg]

	if executableCmd == nil {
		err = errors.New("No such command " + commandArg)
	}

	return executableCmd, piDevice, err
}

func grace(err error) {
	fmt.Println("Error: ", err)
	printUsage()

	os.Exit(1)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("raspcon connect [device] => connects you to a defined Raspberry")
	fmt.Println("raspcon list => lists all of the configured Raspberries")
	fmt.Println("raspcon all [command] => executes given command on all Raspberries")
}
