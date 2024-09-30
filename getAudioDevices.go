package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

func GetAudioDevices() []Device {
	var devicesWin32PnP []Win32_PNPDevice
	var audioDevicesWin32 []Win32_PNPDevice
	var audioDevices []Device
	err := wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying audio I/O devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "AudioEndpoint" {
			audioDevicesWin32 = append(audioDevicesWin32, Win32_PNPDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	for _, audioDevice := range audioDevicesWin32 {
		audioDevices = append(audioDevices, Device{
			Name:                   audioDevice.Name,
			ConfigManagerErrorCode: audioDevice.ConfigManagerErrorCode,
			Status:                 audioDevice.Status,
			Description:            audioDevice.Description,
			DeviceID:               audioDevice.DeviceID,
			PNPDeviceID:            audioDevice.PNPDeviceID,
		})
	}
	if len(audioDevices) < 1 {
		return []Device{}
	} else {
		return audioDevices
	}
}
