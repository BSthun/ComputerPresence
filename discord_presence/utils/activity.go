package utils

import (
	"github.com/hugolgst/rich-go/client"

	"computer-presence/stats"
	"computer-presence/types"
)

var stage = types.BatchStageA

func SetActivity(stat string, sys string, image string, window string) {
	var detail string

	// TODO: Make status flippable
	// if stage == types.BatchStageA {
	// 	detail = window
	// } else {
	// 	detail = stats.GetMusic()
	// }
	// stage.Toggle()

	detail = stats.GetMusic()
	
	err := client.SetActivity(client.Activity{
		State:      stat,
		Details:    detail,
		LargeImage: image,
		LargeText:  window,
		SmallImage: "macos_bigsur",
		SmallText:  sys,
	})

	if err != nil {
		panic(err)
	}
}
