package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetSVGCDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var svgcDevicesWin32 []Win32_PNPDevice
	var svgcDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying sound, video and game controllers devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "MEDIA" {
			svgcDevicesWin32 = append(svgcDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, svgcDevice := range svgcDevicesWin32 {
		svgcDevices = append(svgcDevices, Device{
			Name:                   svgcDevice.Name,
			ConfigManagerErrorCode: svgcDevice.ConfigManagerErrorCode,
			Status:                 svgcDevice.Status,
			Description:            svgcDevice.Description,
			DeviceID:               svgcDevice.DeviceID,
			PNPDeviceID:            svgcDevice.PNPDeviceID,
		})
	}
	if len(svgcDevices) < 1 {
		return []Device{}
	} else {
		return svgcDevices
	}
}
