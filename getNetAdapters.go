package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetNetworkDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var netDevicesWin32 []Win32_PNPDevice
	var netDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying network devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Net" {
			netDevicesWin32 = append(netDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, netDevice := range netDevicesWin32 {
		netDevices = append(netDevices, Device{
			Name:                   netDevice.Name,
			ConfigManagerErrorCode: netDevice.ConfigManagerErrorCode,
			Status:                 netDevice.Status,
			Description:            netDevice.Description,
			DeviceID:               netDevice.DeviceID,
			PNPDeviceID:            netDevice.PNPDeviceID,
		})
	}
	if len(netDevices) < 1 {
		return []Device{}
	} else {
		return netDevices
	}
}
