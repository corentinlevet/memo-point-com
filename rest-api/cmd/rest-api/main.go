package main

import (
	"memo-point-com/internal/config"
	"memo-point-com/internal/network"
)

func main() {
	config.InitTimezone()

	network.InitServer()
}
