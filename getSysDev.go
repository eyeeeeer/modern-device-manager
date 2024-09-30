package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetSystemDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var sysDevicesWin32 []Win32_PNPDevice
	var sysDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying system devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "System" {
			sysDevicesWin32 = append(sysDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, sysDevice := range sysDevicesWin32 {
		sysDevices = append(sysDevices, Device{
			Name:                   sysDevice.Name,
			ConfigManagerErrorCode: sysDevice.ConfigManagerErrorCode,
			Status:                 sysDevice.Status,
			Description:            sysDevice.Description,
			DeviceID:               sysDevice.DeviceID,
			PNPDeviceID:            sysDevice.PNPDeviceID,
		})
	}
	if len(sysDevices) < 1 {
		return []Device{}
	} else {
		return sysDevices
	}
}
