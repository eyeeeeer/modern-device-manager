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
	var batteriesDevicesWin32 []Win32_BatteryDevice
	var batteriesDevices []Device
	err := wmi.Query("SELECT * FROM Win32_Battery", &batteriesDevicesWin32)
	if err != nil {
		fmt.Println("Error querying batteries devices:", err)
		return nil
	}
	fmt.Println(batteriesDevicesWin32)
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
