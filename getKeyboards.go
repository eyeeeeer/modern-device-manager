package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetKeyboardsDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var keyboardDevicesWin32 []Win32_PNPDevice
	var keyboardDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying keyboard devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Keyboard" {
			keyboardDevicesWin32 = append(keyboardDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, keyboardDevice := range keyboardDevicesWin32 {
		keyboardDevices = append(keyboardDevices, Device{
			Name:                   keyboardDevice.Name,
			ConfigManagerErrorCode: keyboardDevice.ConfigManagerErrorCode,
			Status:                 keyboardDevice.Status,
			Description:            keyboardDevice.Description,
			DeviceID:               keyboardDevice.DeviceID,
			PNPDeviceID:            keyboardDevice.PNPDeviceID,
		})
	}
	if len(keyboardDevices) < 1 {
		return []Device{}
	} else {
		return keyboardDevices
	}
}
