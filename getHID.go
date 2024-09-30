package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetHIDDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var hidDevicesWin32 []Win32_PNPDevice
	var hidDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying human interface devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "HIDClass" {
			hidDevicesWin32 = append(hidDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, hidDevice := range hidDevicesWin32 {
		hidDevices = append(hidDevices, Device{
			Name:                   hidDevice.Name,
			ConfigManagerErrorCode: hidDevice.ConfigManagerErrorCode,
			Status:                 hidDevice.Status,
			Description:            hidDevice.Description,
			DeviceID:               hidDevice.DeviceID,
			PNPDeviceID:            hidDevice.PNPDeviceID,
		})
	}
	if len(hidDevices) < 1 {
		return []Device{}
	} else {
		return hidDevices
	}
}
