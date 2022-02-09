package utils

import (
	"github.com/hugolgst/rich-go/client"

	"computer-presence/stats"
	"computer-presence/types"
)

var stage = types.BatchStageA
var batchNo = 0
var music = "N/A"

func SetActivity(stat string, sys string, image string, window string) {

	// TODO: Make status flippable
	// if stage == types.BatchStageA {
	// 	detail = window
	// } else {
	// 	detail = stats.GetMusic()
	// }
	// stage.Toggle()

	if batchNo%5 == 0 {
		music = stats.GetMusic()
	}
	batchNo++

	err := client.SetActivity(client.Activity{
		State:      stat,
		Details:    music,
		LargeImage: image,
		LargeText:  window,
		SmallImage: "macos_bigsur",
		SmallText:  sys,
	})

	if err != nil {
		panic(err)
	}
}
