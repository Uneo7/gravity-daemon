package utils

import (
	"log"
	"os/exec"
	"strconv"
)

func getRam(pid int) int {
	cmd := exec.Command("/bin/sh", "-c", "pgrep -P "+strconv.Itoa(pid)+"| xargs pmap | grep total | awk '{print substr($2, 1, length($2)-1)}' | awk '{s+=$1} END {print s}'")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return -1
	}

	ramString := CleanupString(string(output))
	ram, err := strconv.Atoi(ramString)

	if err != nil {
		return -1
	}

	return ram

}

func getCpu(pid int) float64 {
	cmd := exec.Command("/bin/sh", "-c", "pgrep -P "+strconv.Itoa(pid)+" | xargs ps -o %cpu,cmd -p | awk '{cpu+=$1} END {print cpu}'")
	output, err := cmd.CombinedOutput()

	log.Println(cmd.Args)
	log.Println(string(output))

	if err != nil {
		return -1
	}

	cpuString := CleanupString(string(output))
	cpu, err := strconv.ParseFloat(cpuString, 64)

	if err != nil {
		return -1
	}

	return cpu
}

func Resources(server Server) (ram int, cpu float64) {

	//cpu = -1
	cpu = getCpu(server.Pid)
	ram = getRam(server.Pid)

	return
}
