package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

type DeviceOverviewData struct {
	Model           string
	Arch            int
	DeviceName      string
	CPUName         string
	GPUList         []Device
	RamSize         int64
	RamType         int
	SystemPartition []DrivePartition
}

type DrivePartition struct {
	Label       string
	Letter      string
	Size        int64
	FreeSpace   int64
	Type        int
	isBootable  bool
	isPageFile  bool
	isCrashDump bool
	isRecovery  bool
	isEFI       bool
	FSType      int
}

type AllDeviceTypes struct {
	Audio              []Device `json:"audio"`
	Apos               []Device `json:"apos"`
	Battery            []Device `json:"battery"`
	Biometric          []Device `json:"biometric"`
	Bluetooth          []Device `json:"bluetooth"`
	Camera             []Device `json:"camera"`
	Pc                 []Device `json:"pc"`
	Drive              []Device `json:"drive"`
	Gpu                []Device `json:"gpu"`
	Firmware           []Device `json:"firmware"`
	Hid                []Device `json:"hid"`
	Keyboard           []Device `json:"keyboard"`
	Mouse              []Device `json:"mouse"`
	Display            []Device `json:"display"`
	Network            []Device `json:"network"`
	Printq             []Device `json:"printq"`
	Cpu                []Device `json:"cpu"`
	Secure             []Device `json:"secure"`
	Softwarecomponents []Device `json:"softwarecomponents"`
	Softwaredevices    []Device `json:"softwaredevices"`
	Sound              []Device `json:"sound"`
	Memoryc            []Device `json:"memoryc"`
	Sysdev             []Device `json:"sysdev"`
	Usbc               []Device `json:"usbc"`
	Usbmgr             []Device `json:"usbmgr"`
}

type Device struct {
	Name                     string
	Size                     int64
	SerialNumber             string
	Manufacturer             string
	MediaType                string
	ConfigManagerErrorCode   uint32
	Model                    string
	Status                   string
	Description              string
	DeviceID                 string
	MPU401Address            uint32
	PNPDeviceID              string
	TimeOnBattery            uint32
	TimeToFullCharge         uint32
	SmartBatteryVersion      string
	EstimatedChargeRemaining uint16
	EstimatedRunTime         uint32
	ExpectedBatteryLife      uint32
	ExpectedLife             uint32
	FullChargeCapacity       uint32
	MaxRechargeTime          uint32
	DesignCapacity           uint32
	DesignVoltage            uint64
	BatteryRechargeTime      uint32
	BatteryStatus            uint16
	Chemistry                uint16
}

type Win32_PNPDevice struct {
	Description            string
	Name                   string
	Status                 string
	ConfigManagerErrorCode uint32
	DeviceID               string
	PNPDeviceID            string
	Manufacturer           string
	PNPClass               string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	cb, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		log.Fatal(err)
	}

	cbi, err := strconv.Atoi(cb)
	if err != nil {
		// ... handle error
		panic(err)
	}
	if cbi < 22000 {
		runtime.Quit(ctx)
		vbscript := `
					Set objWMI = GetObject("winmgmts:\\.\root\cimv2")
			Set colOS = objWMI.ExecQuery("SELECT * FROM Win32_OperatingSystem")
			
			For Each objOS in colOS
				If CInt(objOS.BuildNumber) < 22000 Then
					MsgBox "Oops... This app can be run only on Windows 11. Update your system and try again.", vbCritical, "Device Manager"
				End If
			Next
		`
		vbsFile, err := ioutil.TempFile("", "dvmgrosbt*.vbs")
		if err != nil {
			fmt.Println("error when creation temporarily script file:", err)
			return
		}
		defer os.Remove(vbsFile.Name())
		_, err = vbsFile.Write([]byte(vbscript))
		if err != nil {
			fmt.Println("write err", err)
			return
		}
		vbsFile.Close()
		cmd := exec.Command("wscript", vbsFile.Name())
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
		err = cmd.Run()
		if err != nil {
			fmt.Println("error message exec", err)
			return
		}
	}
}

func (a *App) GetAllDevicesList() AllDeviceTypes {
	var devices AllDeviceTypes
	devices.Audio = GetAudioDevices()
	devices.Apos = GetAPODevices()
	devices.Battery = GetBatteriesDevices()
	devices.Biometric = GetBiometricDevices()
	devices.Bluetooth = GetBluetoothDevices()
	devices.Camera = GetCameraDevices()
	devices.Pc = GetPCDevices()
	devices.Drive = GetDriveDevices()
	devices.Gpu = GetGPUDevices()
	devices.Firmware = GetFirmwareDevices()
	devices.Hid = GetHIDDevices()
	devices.Keyboard = GetKeyboardsDevices()
	devices.Mouse = GetMiceDevices()
	devices.Display = GetDisplayDevices()
	devices.Network = GetNetworkDevices()
	devices.Printq = GetPrintQDevices()
	devices.Cpu = GetCPUDevices()
	devices.Secure = GetSecureDevices()
	devices.Softwarecomponents = GetSoftwareComponentsDevices()
	devices.Softwaredevices = GetSoftwareDevices()
	devices.Sound = GetSVGCDevices()
	devices.Memoryc = GetStorageControllersDevices()
	devices.Sysdev = GetSystemDevices()
	devices.Usbc = GetUSBControllersDevices()
	devices.Usbmgr = GetUSBManagersDevices()
	return devices
}
