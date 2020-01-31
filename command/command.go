package command

import "github.com/SlootSantos/rasp-connect/device"

// An Executable is a function
// that can be triggered via a CLI argument
type Executable func(*device.Pi) error

// CommandMap contains all CLI executables
var CommandMap = map[string]Executable{
	"connect": connect,
	"list":    list,
	"all":     all,
}
