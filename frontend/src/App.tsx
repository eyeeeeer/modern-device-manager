import React, { useState, useEffect } from "react";
import "./App.css";
import { Field, Link, ProgressBar, SearchBox } from "@fluentui/react-components";
import { ChevronLeftRegular, ChevronRightRegular } from "@fluentui/react-icons";
import {GetAllDevicesList} from "../wailsjs/go/main/App";

import catList from "./device_types.json";
import { main } from "../wailsjs/go/models";

type catListType = {
    name: string;
    icon: string;
    label: string;
}[];

interface DeviceList {
    audio: main.Device[];
    apos: main.Device[];
    battery: main.Device[];
    biometric: main.Device[];
    bluetooth: main.Device[];
    camera: main.Device[];
    pc: main.Device[];
    drive: main.Device[];
    gpu: main.Device[];
    firmware: main.Device[];
    hid: main.Device[];
    keyboard: main.Device[];
    mouse: main.Device[];
    display: main.Device[];
    network: main.Device[];
    printq: main.Device[];
    cpu: main.Device[];
    secure: main.Device[];
    softwarecomponents: main.Device[];
    softwaredevices: main.Device[];
    sound: main.Device[];
    memoryc: main.Device[];
    sysdev: main.Device[];
    usbc: main.Device[];
    usbmgr: main.Device[];
}

function App() {
    const [ catId, setCatId ] = useState("");
    const [ catIcon, setCatIcon ] = useState("");
    const [ catIndex, setCatIndex ] = useState(0);
    const [ deviceList, setDeviceList ] = useState<any>();

    useEffect(() => {
        const getDevicesList = async () => {
            const devices = await GetAllDevicesList();
            setDeviceList(devices);
        }

        getDevicesList();;
    }, []);

    return (
        <div className="container">
            <div className="deviceInfoPanel">
                <div className="devicePreviewContainer">
                    <img src="/previews/laptops/asus/tuf_gaming/2022.png" alt="Device Preview" className="devicePreview" draggable={false} />
                    <h1>ASUS TUF Gaming F15</h1>
                </div>
                <div className="deviceSpecsContainer">
                    <h3 className="deviceSpecs__pcName">eyeeeeer-laptop</h3>
                    <div className="deviceSpecs__container">
                        <div className="deviceSpecs__pcCPUsContainer">
                            <span>13th Gen Intel(R) Core(TM) i7-13620H@ 2.50GHz</span>
                        </div>
                        <div className="deviceSpecs__pcGPUsContainer">
                            <span>Intel(R) UHD Graphics</span>
                            <span>NVIDIA GeForce RTX 4070 Laptop GPU</span>
                        </div>
                        <div className="deviceSpecs__pcRAMContainer">
                            <span>16 GB DDR5</span>
                        </div>
                        <div className="deviceSpecs__pcPrimaryDriveContainer">
                            <span>Windows (C:) - 476 GB</span>
                            <ProgressBar value={85} max={100} />
                            <div className="deviceSpecs__pcPrimaryDriveSpaceLabelsContainer">
                                <Field validationMessage="404 GB used" validationState="none"></Field>
                                <Field validationMessage="71.3 GB free" validationState="none"></Field>
                            </div>
                        </div>
                        <Link href="ms-settings:about">
                            View more system information in Settings
                        </Link>
                    </div>
                </div>
            </div>
            <div className="deviceTypeListPanel">
                <div className="deviceTypeListPanel__Header">
                    {catId == "" ? (
                        <h2>Device type</h2>
                    ) : (
                        <div className="deviceTypeListPanel_HeaderNav">
                            <button onClick={() => {
                                setCatId("");
                                setCatIcon("");
                                setCatIndex(0);
                            }} title="Back">
                                <ChevronLeftRegular/>
                            </button>
                            <h2>{catList[catIndex].catName.replace("devices", "")} devices</h2>
                        </div>
                    )}
                    <SearchBox placeholder="Search" />
                </div>
                <div className="deviceTypeListPanel__TypesList">
                    {catId == "" ? (
                        catList.map((item, index) => (
                            <div key={item.catId} className="deviceTypeListPanel__TypesList__ListItem" onClick={() => {
                                setCatId(item.catId);
                                setCatIcon(item.catIcon);
                                setCatIndex(index);
                            }}>
                                <div>
                                    <div>
                                        <img src={"/icons/device_types/" + item.catIcon + ".ico"} alt={item.catName} draggable="false" className="deviceTypeListPanel__TypesList__ListItem__Icon" />
                                    </div>
                                    <span className="deviceTypeListPanel__TypesList__ListItem__Label">{item.catName}</span>
                                </div>
                                <ChevronRightRegular />
                            </div>
                        ))
                    ) : (
                        !deviceList[catId as keyof DeviceList] ? (
                            <div className="deviceTypeListPanel__DevicesCategoryEmptyContainer">
                                <span>You don't have a {catList[catIndex].catName.replace("devices", "")} devices.</span>
                            </div>
                        ) : (
                            <div className="deviceTypeListPanel__TypesList">
                                {deviceList[catId as keyof DeviceList].map((drive, index) => (
                                    <div className="deviceTypeListPanel__TypesList__ListItem" key={index}>
                                        <div>
                                            <div>
                                                <img src={"/icons/device_types/" + catIcon + ".ico"} alt={drive.Model} draggable="false" className="deviceTypeListPanel__TypesList__ListItem__Icon" />
                                            </div>
                                            <span className="deviceTypeListPanel__TypesList__ListItem__Label">
                                                {drive.Name}
                                            </span>
                                        </div>
                                        <ChevronRightRegular />
                                    </div>
                                ))}
                            </div>
                        )
                    )}
                </div>
            </div>
            <div className="deviceSelectedInfoPanel">
                <span>Select a device to view more information</span>
            </div>
        </div>
    )
}

export default App; 