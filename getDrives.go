package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

type Win32_DiskDrive struct {
	Description            string
	ConfigManagerErrorCode uint32
	DeviceID               string
	Manufacturer           string
	Model                  string
	MediaType              string
	SerialNumber           string
	Status                 string
	Name                   string
	Size                   int64
}

func GetDriveDevices() []Device {
	var diskDrivesWin32 []Win32_DiskDrive
	var diskDrives []Device
	err := wmi.Query("SELECT * FROM Win32_DiskDrive", &diskDrivesWin32)
	if err != nil {
		fmt.Println("Error querying disk drives:", err)
		return nil
	}
	for _, drive := range diskDrivesWin32 {
		diskDrives = append(diskDrives, Device{
			Name:                   drive.Model,
			Model:                  drive.Model,
			ConfigManagerErrorCode: drive.ConfigManagerErrorCode,
			Status:                 drive.Status,
			SerialNumber:           drive.SerialNumber,
			Description:            drive.Description,
			DeviceID:               drive.DeviceID,
			MediaType:              drive.MediaType,
			Manufacturer:           drive.Manufacturer,
			Size:                   drive.Size,
			MPU401Address:          0,
			PNPDeviceID:            "",
		})
	}
	if len(diskDrives) < 1 {
		return []Device{}
	} else {
		return diskDrives
	}
}
