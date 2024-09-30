package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetGPUDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var gpuDevicesWin32 []Win32_PNPDevice
	var gpuDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying GPU devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Display" {
			gpuDevicesWin32 = append(gpuDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, gpuDevice := range gpuDevicesWin32 {
		gpuDevices = append(gpuDevices, Device{
			Name:                   gpuDevice.Name,
			ConfigManagerErrorCode: gpuDevice.ConfigManagerErrorCode,
			Status:                 gpuDevice.Status,
			Description:            gpuDevice.Description,
			DeviceID:               gpuDevice.DeviceID,
			PNPDeviceID:            gpuDevice.PNPDeviceID,
		})
	}
	if len(gpuDevices) < 1 {
		return []Device{}
	} else {
		return gpuDevices
	}
}
