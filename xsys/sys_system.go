package xsys

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"go-phoenix/handle"
	"net"
	"time"
)

type SysSystem struct {
}

func (o *SysSystem) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	buf := new(bytes.Buffer)

	// Network Addresses
	ars, err := net.InterfaceAddrs()
	if err != nil {
		buf.WriteString("Network End Point Addresses:  --\n")
	} else {
		buf.WriteString("Network End Point Addresses:\n")
		for index, a := range ars {
			buf.WriteString(fmt.Sprintf("\tAddress %02d:  [%s] %s\n", index+1, a.Network(), a))
		}
	}
	buf.WriteByte('\n')

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
		buf.WriteString("CPU Information :  --\n")
	} else {
		buf.WriteString("CPU Information :\n")
		for _, iState := range iStates {
			buf.WriteString(fmt.Sprintf("\t%s %d Cores\n", iState.ModelName, iState.Cores))
		}
	}

	// CPU Percents
	cpuPercents, err := cpu.Percent(0, true)
	if err != nil {
		buf.WriteString("CPU Percents :  --\n")
	} else {
		var total float64
		for _, cpuPercent := range cpuPercents {
			total = total + cpuPercent
		}

		buf.WriteString(fmt.Sprintf("CPU Average Percent :  %.2f%%\n", total/float64(len(cpuPercents))))
		buf.WriteString(fmt.Sprintf("CPU Percents :  \n"))
		for index, cpuPercent := range cpuPercents {
			buf.WriteString(fmt.Sprintf("\tCPU %d :  %.2f%%\n", index+1, cpuPercent))
		}
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

	return buf.String(), nil
}
