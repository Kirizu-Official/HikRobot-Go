package main

import (
	"fmt"
	"github.com/Kirizu-Official/HikRobot-Go/MvsSDK"
)

func main() {

	MvsSDK.GetSDKVersion()
	fmt.Println("MVS SDK初始化成功...")
	MvsSDK.Initialize()
	MvsSDK.EnumerateTls()

	code, res := MvsSDK.EnumDevices(MvsSDK.MvGigeDevice)
	fmt.Println(code, len(res))

	fmt.Printf("code:%d, 成功找到 %d 个设备\n", code, len(res))
	fmt.Println(string(res[0].SpecialInfo.MvGigeDeviceInfo.ManufacturerName[:]), string(res[0].SpecialInfo.MvGigeDeviceInfo.ModelName[:]))

	fmt.Println(MvsSDK.IsDeviceAccessible(res[0], MvsSDK.MvAccessExclusive))

	device := MvsSDK.Device{}
	fmt.Println(device.CreateHandle(res[0]), "create")
	fmt.Println(device.OpenDevice(MvsSDK.MvAccessExclusive, 0), "open")
	fmt.Println(device.IsDeviceConnected())
}
