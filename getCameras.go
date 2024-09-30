package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetCameraDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var camDevicesWin32 []Win32_PNPDevice
	var camDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying camera devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Camera" {
			camDevicesWin32 = append(camDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, btDevice := range camDevicesWin32 {
		camDevices = append(camDevices, Device{
			Name:                   btDevice.Name,
			ConfigManagerErrorCode: btDevice.ConfigManagerErrorCode,
			Status:                 btDevice.Status,
			Description:            btDevice.Description,
			DeviceID:               btDevice.DeviceID,
			PNPDeviceID:            btDevice.PNPDeviceID,
		})
	}
	if len(camDevices) < 1 {
		return []Device{}
	} else {
		return camDevices
	}
}
