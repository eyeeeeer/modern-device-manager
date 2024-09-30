package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetPCDevices() []Device { // Usless category :)
	var devicesWin32PnP []Win32_PNPDevice
	var pcDevicesWin32 []Win32_PNPDevice
	var pcDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying PC devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Computer" {
			pcDevicesWin32 = append(pcDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, pcDevice := range pcDevicesWin32 {
		pcDevices = append(pcDevices, Device{
			Name:                   pcDevice.Name,
			ConfigManagerErrorCode: pcDevice.ConfigManagerErrorCode,
			Status:                 pcDevice.Status,
			Description:            pcDevice.Description,
			DeviceID:               pcDevice.DeviceID,
			PNPDeviceID:            pcDevice.PNPDeviceID,
		})
	}
	if len(pcDevices) < 1 {
		return []Device{}
	} else {
		return pcDevices
	}
}
