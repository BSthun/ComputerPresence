package stats

import (
	"log"
	"os/exec"
	"regexp"
)

var windowMacRegex = regexp.MustCompile("=\"(.*?)\"")

func GetWindowMac() string {
	// lsappinfo info -only name `lsappinfo front`
	front, err := exec.Command("lsappinfo", "front").Output()
	if err != nil {
		log.Fatal(err)
	}

	name, err := exec.Command("lsappinfo", "info", "-only", "name", string(front)).Output()
	if err != nil {
		log.Fatal(err)
	}

	rs := windowMacRegex.FindStringSubmatch(string(name))

	return rs[1]
}
