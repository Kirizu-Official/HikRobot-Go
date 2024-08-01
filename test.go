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

	var data interface{}
	result := MvsSDK.EnumDevices(&data)
	fmt.Printf("成功找到 %d 个设备\n", len(result))
	fmt.Println(string(result[0].SpecialInfo.MvGigeDeviceInfo.ManufacturerName[:]), string(result[0].SpecialInfo.MvGigeDeviceInfo.ModelName[:]))
}
