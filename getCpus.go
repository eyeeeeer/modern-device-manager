package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetCPUDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var cpuDevicesWin32 []Win32_PNPDevice
	var cpuDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying CPU devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Processor" {
			cpuDevicesWin32 = append(cpuDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, cpuDevice := range cpuDevicesWin32 {
		cpuDevices = append(cpuDevices, Device{
			Name:                   cpuDevice.Name,
			ConfigManagerErrorCode: cpuDevice.ConfigManagerErrorCode,
			Status:                 cpuDevice.Status,
			Description:            cpuDevice.Description,
			DeviceID:               cpuDevice.DeviceID,
			PNPDeviceID:            cpuDevice.PNPDeviceID,
		})
	}
	if len(cpuDevices) < 1 {
		return []Device{}
	} else {
		return cpuDevices
	}
}
