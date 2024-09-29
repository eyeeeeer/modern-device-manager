export namespace main {
	
	export class Device {
	    Name: string;
	    Size: number;
	    SerialNumber: string;
	    Manufacturer: string;
	    MediaType: string;
	    ConfigManagerErrorCode: number;
	    Model: string;
	    Status: string;
	    Description: string;
	    DeviceID: string;
	    MPU401Address: number;
	    PNPDeviceID: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Size = source["Size"];
	        this.SerialNumber = source["SerialNumber"];
	        this.Manufacturer = source["Manufacturer"];
	        this.MediaType = source["MediaType"];
	        this.ConfigManagerErrorCode = source["ConfigManagerErrorCode"];
	        this.Model = source["Model"];
	        this.Status = source["Status"];
	        this.Description = source["Description"];
	        this.DeviceID = source["DeviceID"];
	        this.MPU401Address = source["MPU401Address"];
	        this.PNPDeviceID = source["PNPDeviceID"];
	    }
	}
	export class AllDeviceTypes {
	    audio: Device[];
	    apos: Device[];
	    battery: Device[];
	    biometric: Device[];
	    bluetooth: Device[];
	    camera: Device[];
	    pc: Device[];
	    drive: Device[];
	    gpu: Device[];
	    firmware: Device[];
	    hid: Device[];
	    keyboard: Device[];
	    mouse: Device[];
	    display: Device[];
	    network: Device[];
	    printq: Device[];
	    cpu: Device[];
	    secure: Device[];
	    softwarecomponents: Device[];
	    softwaredevices: Device[];
	    sound: Device[];
	    memoryc: Device[];
	    sysdev: Device[];
	    usbc: Device[];
	    usbmgr: Device[];
	
	    static createFrom(source: any = {}) {
	        return new AllDeviceTypes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.audio = this.convertValues(source["audio"], Device);
	        this.apos = this.convertValues(source["apos"], Device);
	        this.battery = this.convertValues(source["battery"], Device);
	        this.biometric = this.convertValues(source["biometric"], Device);
	        this.bluetooth = this.convertValues(source["bluetooth"], Device);
	        this.camera = this.convertValues(source["camera"], Device);
	        this.pc = this.convertValues(source["pc"], Device);
	        this.drive = this.convertValues(source["drive"], Device);
	        this.gpu = this.convertValues(source["gpu"], Device);
	        this.firmware = this.convertValues(source["firmware"], Device);
	        this.hid = this.convertValues(source["hid"], Device);
	        this.keyboard = this.convertValues(source["keyboard"], Device);
	        this.mouse = this.convertValues(source["mouse"], Device);
	        this.display = this.convertValues(source["display"], Device);
	        this.network = this.convertValues(source["network"], Device);
	        this.printq = this.convertValues(source["printq"], Device);
	        this.cpu = this.convertValues(source["cpu"], Device);
	        this.secure = this.convertValues(source["secure"], Device);
	        this.softwarecomponents = this.convertValues(source["softwarecomponents"], Device);
	        this.softwaredevices = this.convertValues(source["softwaredevices"], Device);
	        this.sound = this.convertValues(source["sound"], Device);
	        this.memoryc = this.convertValues(source["memoryc"], Device);
	        this.sysdev = this.convertValues(source["sysdev"], Device);
	        this.usbc = this.convertValues(source["usbc"], Device);
	        this.usbmgr = this.convertValues(source["usbmgr"], Device);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

