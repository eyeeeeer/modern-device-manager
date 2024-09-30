package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetSoftwareComponentsDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var scDevicesWin32 []Win32_PNPDevice
	var scDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying software components devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "SoftwareComponent" {
			scDevicesWin32 = append(scDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, scDevice := range scDevicesWin32 {
		scDevices = append(scDevices, Device{
			Name:                   scDevice.Name,
			ConfigManagerErrorCode: scDevice.ConfigManagerErrorCode,
			Status:                 scDevice.Status,
			Description:            scDevice.Description,
			DeviceID:               scDevice.DeviceID,
			PNPDeviceID:            scDevice.PNPDeviceID,
		})
	}
	if len(scDevices) < 1 {
		return []Device{}
	} else {
		return scDevices
	}
}
