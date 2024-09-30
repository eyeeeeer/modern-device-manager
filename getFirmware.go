package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetFirmwareDevices() []Device { // yet another useless category
	var devicesWin32PnP []Win32_PNPDevice
	var firmwareDevicesWin32 []Win32_PNPDevice
	var firmwareDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying Firmware devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Firmware" {
			firmwareDevicesWin32 = append(firmwareDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, firmwareDevice := range firmwareDevicesWin32 {
		firmwareDevices = append(firmwareDevices, Device{
			Name:                   firmwareDevice.Name,
			ConfigManagerErrorCode: firmwareDevice.ConfigManagerErrorCode,
			Status:                 firmwareDevice.Status,
			Description:            firmwareDevice.Description,
			DeviceID:               firmwareDevice.DeviceID,
			PNPDeviceID:            firmwareDevice.PNPDeviceID,
		})
	}
	if len(firmwareDevices) < 1 {
		return []Device{}
	} else {
		return firmwareDevices
	}
}
