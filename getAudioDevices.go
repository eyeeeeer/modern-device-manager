package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

type Win32_SoundDevice struct {
	Name                   string
	ConfigManagerErrorCode uint32
	Description            string
	DeviceID               string
	Manufacturer           string
	MPU401Address          uint32
	PNPDeviceID            string
	ProductName            string
	Status                 string
}

func GetAudioDevices() []Device {
	var audioDevicesWin32 []Win32_SoundDevice
	var audioDevices []Device
	err := wmi.Query("SELECT * FROM Win32_SoundDevice", &audioDevicesWin32)
	if err != nil {
		fmt.Println("Error querying audio devices:", err)
		return nil
	}
	for _, audioDevice := range audioDevicesWin32 {
		audioDevices = append(audioDevices, Device{
			Name:                   audioDevice.Name,
			Model:                  "",
			ConfigManagerErrorCode: audioDevice.ConfigManagerErrorCode,
			Status:                 audioDevice.Status,
			SerialNumber:           "",
			Description:            audioDevice.Description,
			DeviceID:               audioDevice.DeviceID,
			MediaType:              "",
			Manufacturer:           audioDevice.Manufacturer,
			Size:                   0,
			MPU401Address:          audioDevice.MPU401Address,
			PNPDeviceID:            audioDevice.PNPDeviceID,
		})
	}
	return audioDevices
}
