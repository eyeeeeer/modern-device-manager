package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetPrintQDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var printDevicesWin32 []Win32_PNPDevice
	var printDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying print queue devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "PrintQueue" {
			printDevicesWin32 = append(printDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, printDevice := range printDevicesWin32 {
		printDevices = append(printDevices, Device{
			Name:                   printDevice.Name,
			ConfigManagerErrorCode: printDevice.ConfigManagerErrorCode,
			Status:                 printDevice.Status,
			Description:            printDevice.Description,
			DeviceID:               printDevice.DeviceID,
			PNPDeviceID:            printDevice.PNPDeviceID,
		})
	}
	if len(printDevices) < 1 {
		return []Device{}
	} else {
		return printDevices
	}
}
