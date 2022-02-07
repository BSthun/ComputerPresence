package stats

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/host"
)

func GetSystem() string {
	hostnameStat, err := os.Hostname()
	if err != nil {
		println("UNABLE TO GET HOSTNAME")
		hostnameStat = "N/A"
	}

	platformStat, _, versionStat, err := host.PlatformInformation()
	if err != nil {
		println("UNABLE TO GET OS")
		platformStat = "N/A"
		versionStat = "N/A"
	}

	stat := fmt.Sprintf("%s %s (%s)", platformStat, versionStat, hostnameStat)

	return stat
}
