package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetAPODevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var aposDevicesWin32 []Win32_PNPDevice
	var aposDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying APOs devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "AudioProcessingObject" {
			aposDevicesWin32 = append(aposDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	fmt.Println(aposDevicesWin32)
	for _, apoDevice := range aposDevicesWin32 {
		aposDevices = append(aposDevices, Device{
			Name:                   apoDevice.Name,
			ConfigManagerErrorCode: apoDevice.ConfigManagerErrorCode,
			Status:                 apoDevice.Status,
			Description:            apoDevice.Description,
			DeviceID:               apoDevice.DeviceID,
			PNPDeviceID:            apoDevice.PNPDeviceID,
		})
	}
	return aposDevices
}
