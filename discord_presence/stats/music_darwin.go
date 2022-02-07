package stats

import (
	"os/exec"
	"strings"

	"computer-presence/embed"
)

var musicPlayers = []string{
	"Music",
	"Spotify",
}

func GetMusic() string {

	for _, player := range musicPlayers {
		osa := strings.ReplaceAll(embed.MusicOsaScript, "{{PLAYER}}", player)
		result, err := exec.Command("osascript", "-e", osa).Output()

		if err != nil || len(result) == 0 {
			continue
		}

		return string(result)
	}

	return "N/A"

}
