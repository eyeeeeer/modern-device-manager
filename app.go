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
	Name                   string
	Size                   int64
	SerialNumber           string
	Manufacturer           string
	MediaType              string
	ConfigManagerErrorCode uint32
	Model                  string
	Status                 string
	Description            string
	DeviceID               string
	MPU401Address          uint32
	PNPDeviceID            string
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
					MsgBox "Oops... This app can be runned only on Windows 11. Upgrade your system and try again.", vbCritical, "Device Manager"
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
	devices.Drive = GetDriveDevices()
	return devices
}
