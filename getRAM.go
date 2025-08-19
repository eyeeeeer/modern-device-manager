package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRam() (RamGB int64) {
	cmd := exec.Command("powershell", "-Command",
		"(Get-CimInstance Win32_ComputerSystem).TotalPhysicalMemory")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	memStr := strings.TrimSpace(string(out))
	memBytes, err := strconv.ParseUint(memStr, 10, 64)
	if err != nil {
		panic(err)
	}
	memGB := memBytes / 1024 / 1024 / 1024
	return int64(memGB)
}
