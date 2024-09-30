package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetSoftwareDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var sdDevicesWin32 []Win32_PNPDevice
	var sdDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying software devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "SoftwareDevice" {
			sdDevicesWin32 = append(sdDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, sdDevice := range sdDevicesWin32 {
		sdDevices = append(sdDevices, Device{
			Name:                   sdDevice.Name,
			ConfigManagerErrorCode: sdDevice.ConfigManagerErrorCode,
			Status:                 sdDevice.Status,
			Description:            sdDevice.Description,
			DeviceID:               sdDevice.DeviceID,
			PNPDeviceID:            sdDevice.PNPDeviceID,
		})
	}
	if len(sdDevices) < 1 {
		return []Device{}
	} else {
		return sdDevices
	}
}
