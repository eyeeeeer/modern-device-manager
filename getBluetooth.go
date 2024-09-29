package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetBluetoothDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var btDevicesWin32 []Win32_PNPDevice
	var btDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying Bluetooth devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Bluetooth" {
			btDevicesWin32 = append(btDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, apoDevice := range btDevicesWin32 {
		btDevices = append(btDevices, Device{
			Name:                   apoDevice.Name,
			ConfigManagerErrorCode: apoDevice.ConfigManagerErrorCode,
			Status:                 apoDevice.Status,
			Description:            apoDevice.Description,
			DeviceID:               apoDevice.DeviceID,
			PNPDeviceID:            apoDevice.PNPDeviceID,
		})
	}
	return btDevices
}
