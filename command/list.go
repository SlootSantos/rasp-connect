package command

import (
	"fmt"

	"github.com/SlootSantos/rasp-connect/config"
	"github.com/SlootSantos/rasp-connect/device"
)

func list(pi *device.Pi) error {
	for id, dev := range config.Pis {
		fmt.Println("\n", id)
		fmt.Println("----------")
		fmt.Println("Nick:", dev.Nickname)
		fmt.Println("Host:", dev.Hostname)
		fmt.Println("User", dev.User)
	}

	return nil
}
