package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

type Win32_BatteryDevice struct {
	BatteryRechargeTime      uint32
	BatteryStatus            uint16
	Chemistry                uint16
	ConfigManagerErrorCode   uint32
	Description              string
	DesignCapacity           uint32
	DesignVoltage            uint64
	DeviceID                 string
	EstimatedChargeRemaining uint16
	EstimatedRunTime         uint32
	ExpectedBatteryLife      uint32
	ExpectedLife             uint32
	FullChargeCapacity       uint32
	MaxRechargeTime          uint32
	Name                     string
	PNPDeviceID              string
	SmartBatteryVersion      string
	Status                   string
	TimeOnBattery            uint32
	TimeToFullCharge         uint32
}

func GetBatteriesDevices() []Device {
	var batteriesDevicesWin32Battery []Win32_BatteryDevice
	var devicesWin32PnP []Win32_PNPDevice
	var batteriesDevicesWin32 []Win32_BatteryDevice
	var batteriesDevices []Device
	err := wmi.Query("SELECT * FROM Win32_Battery", &batteriesDevicesWin32Battery)
	err = wmi.Query("SELECT * FROM Win32_PnPEntity", &devicesWin32PnP)
	if err != nil {
		fmt.Println("Error querying batteries devices:", err)
		return nil
	}
	for _, pnpDevice := range devicesWin32PnP {
		if pnpDevice.PNPClass == "Battery" {
			batteriesDevicesWin32 = append(batteriesDevicesWin32, Win32_BatteryDevice{
				Name:                   pnpDevice.Name,
				ConfigManagerErrorCode: pnpDevice.ConfigManagerErrorCode,
				Status:                 pnpDevice.Status,
				Description:            pnpDevice.Description,
				DeviceID:               pnpDevice.DeviceID,
				PNPDeviceID:            pnpDevice.PNPDeviceID,
			})
		}
	}
	batteriesDevicesWin32 = append(batteriesDevicesWin32, batteriesDevicesWin32Battery...)
	for _, batteryDevice := range batteriesDevicesWin32 {
		batteriesDevices = append(batteriesDevices, Device{
			Name:                     batteryDevice.Name,
			BatteryRechargeTime:      batteryDevice.BatteryRechargeTime,
			BatteryStatus:            batteryDevice.BatteryStatus,
			Chemistry:                batteryDevice.Chemistry,
			ConfigManagerErrorCode:   batteryDevice.ConfigManagerErrorCode,
			DesignCapacity:           batteryDevice.DesignCapacity,
			DesignVoltage:            batteryDevice.DesignVoltage,
			EstimatedChargeRemaining: batteryDevice.EstimatedChargeRemaining,
			EstimatedRunTime:         batteryDevice.EstimatedRunTime,
			ExpectedBatteryLife:      batteryDevice.ExpectedBatteryLife,
			ExpectedLife:             batteryDevice.ExpectedLife,
			FullChargeCapacity:       batteryDevice.FullChargeCapacity,
			MaxRechargeTime:          batteryDevice.MaxRechargeTime,
			SmartBatteryVersion:      batteryDevice.SmartBatteryVersion,
			TimeOnBattery:            batteryDevice.TimeOnBattery,
			TimeToFullCharge:         batteryDevice.TimeToFullCharge,
			Status:                   batteryDevice.Status,
			Description:              batteryDevice.Description,
			DeviceID:                 batteryDevice.DeviceID,
			PNPDeviceID:              batteryDevice.PNPDeviceID,
		})
	}
	return batteriesDevices
}
