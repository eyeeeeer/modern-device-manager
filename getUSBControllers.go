package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetUSBControllersDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var usbcntrlrDevicesWin32 []Win32_PNPDevice
	var usbcntrlrDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying USB controllers devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "USB" {
			usbcntrlrDevicesWin32 = append(usbcntrlrDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, usbcntrlrDevice := range usbcntrlrDevicesWin32 {
		usbcntrlrDevices = append(usbcntrlrDevices, Device{
			Name:                   usbcntrlrDevice.Name,
			ConfigManagerErrorCode: usbcntrlrDevice.ConfigManagerErrorCode,
			Status:                 usbcntrlrDevice.Status,
			Description:            usbcntrlrDevice.Description,
			DeviceID:               usbcntrlrDevice.DeviceID,
			PNPDeviceID:            usbcntrlrDevice.PNPDeviceID,
		})
	}
	if len(usbcntrlrDevices) < 1 {
		return []Device{}
	} else {
		return usbcntrlrDevices
	}
}
