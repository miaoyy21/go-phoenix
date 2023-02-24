package base

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"io"
	"net/http"
	"strings"
	"time"
)

func System(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)

	// Host
	hState, err := host.Info()
	if err != nil {
		buf.WriteString("Host State:  --\n")
	} else {
		buf.WriteString(fmt.Sprintf("Host ID:  %s\n", hState.HostID))
		buf.WriteString(fmt.Sprintf("Host Name:  %s\n", hState.Hostname))
		buf.WriteByte('\n')
		buf.WriteString(fmt.Sprintf("Operating System:  %s\n", hState.OS))
		buf.WriteString(fmt.Sprintf("Platform:  %s\n", hState.Platform))
		buf.WriteString(fmt.Sprintf("Platform Family:  %s\n", hState.PlatformFamily))
		buf.WriteString(fmt.Sprintf("Platform Version:  %s\n", hState.PlatformVersion))
		buf.WriteByte('\n')
		buf.WriteString(fmt.Sprintf("Kernel Architecture:  %s\n", hState.KernelArch))
		buf.WriteString(fmt.Sprintf("Kernel Version:  %s\n", hState.KernelVersion))
		buf.WriteByte('\n')
		buf.WriteString(fmt.Sprintf("Boot Time:  %s\n", time.Unix(int64(hState.BootTime), 0).Format("2006-01-02 15:04:05")))
		buf.WriteString(fmt.Sprintf("Now Time:  %s\n", time.Now().Local().Format("2006-01-02 15:04:05")))
		buf.WriteString(fmt.Sprintf("Up Time:  %s\n", time.Duration(hState.Uptime*1e9)))
	}
	buf.WriteByte('\n')

	// CPU
	iStates, err := cpu.Info()
	if err != nil {
		buf.WriteString("CPU States :  --\n")
	} else {
		buf.WriteString(fmt.Sprintf("CPU %d States :\n", len(iStates)))
		for index, iState := range iStates {
			buf.WriteString(fmt.Sprintf("\tCPU %d :  %s %d Cores\n", index+1, iState.ModelName, iState.Cores))
		}
	}

	// CPU Percents
	cpuPercents, err := cpu.Percent(0, true)
	if err != nil {
		buf.WriteString("CPU Percents :  --\n")
	} else {
		percents := make([]string, 0, len(cpuPercents))
		var total float64
		for _, cpuPercent := range cpuPercents {
			total = total + cpuPercent
			percents = append(percents, fmt.Sprintf("%.2f%%", cpuPercent))
		}

		buf.WriteString(fmt.Sprintf("CPU Average Percent :  %.2f%%\n", total/float64(len(cpuPercents))))
		buf.WriteString(fmt.Sprintf("CPU Percents :  %s\n", strings.Join(percents, ", ")))
	}
	buf.WriteByte('\n')

	dUsage, err := disk.Usage("/")
	if err != nil {
		buf.WriteString("Disk Usage :  --\n")
	} else {
		buf.WriteString(fmt.Sprintf("Disk Used Percent :  %.2f%%\n", dUsage.UsedPercent))
		buf.WriteString(fmt.Sprintf("Disk Usage :  Total %.2f GB, Used %.2f GB, Free %.2f GB", float64(dUsage.Total)/float64(1<<30), float64(dUsage.Used)/float64(1<<30), float64(dUsage.Free)/float64(1<<30)))
	}
	buf.WriteByte('\n')

	io.Copy(w, buf)
}
