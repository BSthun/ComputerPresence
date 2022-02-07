package utils

import "github.com/hugolgst/rich-go/client"

func SetActivity(stat string, sys string, image string, window string) {
	err := client.SetActivity(client.Activity{
		State:      stat,
		Details:    window,
		LargeImage: image,
		LargeText:  window,
		SmallImage: "macos_bigsur",
		SmallText:  sys,
	})

	if err != nil {
		panic(err)
	}
}
