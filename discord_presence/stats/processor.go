package stats

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetProcessor() string {
	cpuStat, err := cpu.Percent(100*time.Millisecond, false)
	if err != nil {
		println("UNABLE TO GET CPU INFO")
		cpuStat = []float64{0}
	}

	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		println("UNABLE TO GET MEMORY INFO")
		memoryStat = &mem.VirtualMemoryStat{
			Used: 0,
		}
	}

	stat := fmt.Sprintf("%.2f%% CPU | %.2f GB RAM", cpuStat[0], float64(memoryStat.Used)/1024/1024/1024)

	return stat
}
