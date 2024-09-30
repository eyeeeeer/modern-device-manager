package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetUSBManagersDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var usbmanDevicesWin32 []Win32_PNPDevice
	var usbmanDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying USB Connector Managers devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "UCM" {
			usbmanDevicesWin32 = append(usbmanDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, usbmanDevice := range usbmanDevicesWin32 {
		usbmanDevices = append(usbmanDevices, Device{
			Name:                   usbmanDevice.Name,
			ConfigManagerErrorCode: usbmanDevice.ConfigManagerErrorCode,
			Status:                 usbmanDevice.Status,
			Description:            usbmanDevice.Description,
			DeviceID:               usbmanDevice.DeviceID,
			PNPDeviceID:            usbmanDevice.PNPDeviceID,
		})
	}
	if len(usbmanDevices) < 1 {
		return []Device{}
	} else {
		return usbmanDevices
	}
}
