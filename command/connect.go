package command

import (
	"errors"
	"fmt"

	"github.com/SlootSantos/rasp-connect/device"
)

func connect(pi *device.Pi) error {
	if pi.Hostname == "" {
		err := errors.New("Device not found")

		return err
	}

	fmt.Println("Connection to Pi: ", pi.Nickname)
	return pi.Connect()
}
