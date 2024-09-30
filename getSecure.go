package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetSecureDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var secDevicesWin32 []Win32_PNPDevice
	var secDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying security devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "SecurityDevices" {
			secDevicesWin32 = append(secDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, secDevice := range secDevicesWin32 {
		secDevices = append(secDevices, Device{
			Name:                   secDevice.Name,
			ConfigManagerErrorCode: secDevice.ConfigManagerErrorCode,
			Status:                 secDevice.Status,
			Description:            secDevice.Description,
			DeviceID:               secDevice.DeviceID,
			PNPDeviceID:            secDevice.PNPDeviceID,
		})
	}
	if len(secDevices) < 1 {
		return []Device{}
	} else {
		return secDevices
	}
}
