package command

import (
	"fmt"
	"github.com/SlootSantos/rasp-connect/config"
	"github.com/SlootSantos/rasp-connect/device"
	"sync"
)

func all(pi *device.Pi) error {
	var wg sync.WaitGroup

	for id := range config.Pis {
		fmt.Println("-----------------")
		fmt.Println("EXEC \"hostname && ls\" FOR: ", id)

		wg.Add(1)
		go config.Pis[id].ExecAsync("", &wg)
	}

	wg.Wait()

	return nil
}
