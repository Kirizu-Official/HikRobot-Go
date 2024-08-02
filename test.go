package main

import (
	"fmt"
	"github.com/Kirizu-Official/HikRobot-Go/MvsSDK"
	"os"
)

func main() {

	MvsSDK.GetSDKVersion()
	fmt.Println("MVS SDK初始化成功...")
	MvsSDK.Initialize()
	MvsSDK.EnumerateTls()

	code, res := MvsSDK.EnumDevices(MvsSDK.MvLayerGigeDevice)
	fmt.Println(code, len(res))

	fmt.Printf("code:%d, 成功找到 %d 个设备\n", code, len(res))
	fmt.Println(res[0].TLayerType, string(res[0].SpecialInfo.MvGigeDeviceInfo.ModelName[:]), string(res[0].SpecialInfo.MvGigeDeviceInfo.SerialNumber[:]))

	fmt.Println(MvsSDK.IsDeviceAccessible(res[0], MvsSDK.MvAccessExclusive))

	device := MvsSDK.Device{}
	fmt.Println(device.CreateHandle(res[0]), "create")
	fmt.Println(device.OpenDevice(MvsSDK.MvAccessExclusive, 0), "open")
	fmt.Println(device.IsDeviceConnected())
	control := device.NewDeviceControl()
	control.OpenParamsGUI()
	camera := device.NewDeviceImage()
	camera.OpenDevice(MvsSDK.MvAccessExclusive, 0)
	camera.SetImageNodeNum(1)
	camera.StartGrabbing()
	var data [MvsSDK.MaxFrameSize]byte
	result, info := camera.GetOneFrameTimeout(&data, 10*1024*1024, 500)
	camera.StopGrabbing()
	camera.CloseDevice()
	fmt.Println(result, info.Height, info.Height, info.ChunkHeight, info.ChunkWidth, info.FrameLen, info.PixelType)
	os.WriteFile("test.jpg", data[:info.FrameLen], 0755)

}
