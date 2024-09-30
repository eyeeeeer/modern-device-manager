package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetDisplayDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var displayDevicesWin32 []Win32_PNPDevice
	var displayDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying display devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Monitor" {
			displayDevicesWin32 = append(displayDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, displayDevice := range displayDevicesWin32 {
		displayDevices = append(displayDevices, Device{
			Name:                   displayDevice.Name,
			ConfigManagerErrorCode: displayDevice.ConfigManagerErrorCode,
			Status:                 displayDevice.Status,
			Description:            displayDevice.Description,
			DeviceID:               displayDevice.DeviceID,
			PNPDeviceID:            displayDevice.PNPDeviceID,
		})
	}
	if len(displayDevices) < 1 {
		return []Device{}
	} else {
		return displayDevices
	}
}
