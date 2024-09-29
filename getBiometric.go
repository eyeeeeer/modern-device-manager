package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetBiometricDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var biometricDevicesWin32 []Win32_PNPDevice
	var biometricDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying Biometric devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Biometric" {
			biometricDevicesWin32 = append(biometricDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, biometricDevice := range biometricDevicesWin32 {
		biometricDevices = append(biometricDevices, Device{
			Name:                   biometricDevice.Name,
			ConfigManagerErrorCode: biometricDevice.ConfigManagerErrorCode,
			Status:                 biometricDevice.Status,
			Description:            biometricDevice.Description,
			DeviceID:               biometricDevice.DeviceID,
			PNPDeviceID:            biometricDevice.PNPDeviceID,
		})
	}
	if len(biometricDevices) < 1 {
		return []Device{}
	} else {
		return biometricDevices
	}
}
