package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"

	"computer-presence/stats"
	"computer-presence/utils"
)

func update() {
	// TODO: Add get window for Windows OS
	currentWindow := stats.GetWindowMac()
	utils.SetActivity(
		stats.GetProcessor(),
		stats.GetSystem(),
		strings.ReplaceAll(strings.ToLower(currentWindow), " ", "_"),
		currentWindow,
	)
}

func main() {
	err := client.Login("940150400676417567")
	if err != nil {
		panic(err)
	}

	utils.SetInterval(update, 2*time.Second)

	scan, err := fmt.Scanln()
	if err != nil {
		return
	}

	println(scan)
}
