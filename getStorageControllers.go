package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetStorageControllersDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var scntrlrDevicesWin32 []Win32_PNPDevice
	var scntrlrDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying storage controllers devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "SCSIAdapter" {
			scntrlrDevicesWin32 = append(scntrlrDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, scntrlrDevice := range scntrlrDevicesWin32 {
		scntrlrDevices = append(scntrlrDevices, Device{
			Name:                   scntrlrDevice.Name,
			ConfigManagerErrorCode: scntrlrDevice.ConfigManagerErrorCode,
			Status:                 scntrlrDevice.Status,
			Description:            scntrlrDevice.Description,
			DeviceID:               scntrlrDevice.DeviceID,
			PNPDeviceID:            scntrlrDevice.PNPDeviceID,
		})
	}
	if len(scntrlrDevices) < 1 {
		return []Device{}
	} else {
		return scntrlrDevices
	}
}
