package config

import "github.com/SlootSantos/rasp-connect/device"

// Pis is a map of Raspberry Pi devices
var Pis = map[string]*device.Pi{
	"A": &device.Pi{
		Nickname: "A",
		Hostname: "raspberrypi4.local",
		User:     "A",
		Password: "raspberryA",
		Port:     "22",
	},
	"B": &device.Pi{
		Nickname: "B",
		Hostname: "raspberrypi.local",
		User:     "B",
		Password: "raspberryB",
		Port:     "22",
	},
	"C": &device.Pi{
		Nickname: "C",
		Hostname: "raspberrypi4.local",
		User:     "C",
		Password: "raspberryC",
		Port:     "22",
	},
	"D": &device.Pi{
		Nickname: "D",
		Hostname: "raspberrypi.local",
		User:     "D",
		Password: "raspberryD",
		Port:     "22",
	},
	"E": &device.Pi{
		Nickname: "E",
		Hostname: "raspberrypi4.local",
		User:     "E",
		Password: "raspberryE",
		Port:     "22",
	},
	"F": &device.Pi{
		Nickname: "F",
		Hostname: "raspberrypi.local",
		User:     "F",
		Password: "raspberryF",
		Port:     "22",
	},
}
