package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetMiceDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var miceDevicesWin32 []Win32_PNPDevice
	var miceDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying mice and other pointing devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Mouse" {
			miceDevicesWin32 = append(miceDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, miceDevice := range miceDevicesWin32 {
		miceDevices = append(miceDevices, Device{
			Name:                   miceDevice.Name,
			ConfigManagerErrorCode: miceDevice.ConfigManagerErrorCode,
			Status:                 miceDevice.Status,
			Description:            miceDevice.Description,
			DeviceID:               miceDevice.DeviceID,
			PNPDeviceID:            miceDevice.PNPDeviceID,
		})
	}
	if len(miceDevices) < 1 {
		return []Device{}
	} else {
		return miceDevices
	}
}
