package MvsSDK

/*
#cgo CFLAGS: -I../include
#cgo windows amd64 LDFLAGS: -L${SRCDIR}/../lib/win/32
#cgo windows 386 LDFLAGS: -L${SRCDIR}/../lib/win/64
#cgo linux amd64 LDFLAGS: -L${SRCDIR}/../lib/linux/64 -Wl,-rpath,$(SRCDIR)/../lib/linux/64
#cgo linux 386 LDFLAGS: -L${SRCDIR}/../lib/linux/32 -Wl,-rpath,$(SRCDIR)/../lib/linux/32
#cgo LDFLAGS: -lMvCameraControl -static
#include "MvCameraControl.h"
*/
import "C"
import (
	"unsafe"
)

// C 语言类型	CGO 类型	Go 语言类型
// char	C.char	byte
// singed char	C.schar	int8
// unsigned char	C.uchar	uint8
// short	C.short	int16
// unsigned short	C.ushort	uint16
// int	C.int	int32
// unsigned int	C.uint	uint32
// long	C.long	int32
// unsigned long	C.ulong	uint32
// long long int	C.longlong	int64
// unsigned long long int	C.ulonglong	uint64
// float	C.float	float32
// double	C.double	float64
// size_t	C.size_t	uint

// GetSDKVersion 获取SDK版本信息接口
func GetSDKVersion() uint32 {
	return uint32(C.MV_CC_GetSDKVersion())
}

// Initialize 初始化SDK
func Initialize() int {
	return int(C.MV_CC_Initialize())
}

// Finalize 反初始化SDK，释放资源
func Finalize() int {
	return int(C.MV_CC_Finalize())
}

// EnumerateTls 获取支持的传输层
// return 当前支持的传输层协议类型和数值总和
func EnumerateTls() int {
	return int(C.MV_CC_EnumerateTls())
}

// EnumDevices 枚举设备，支持枚举对应采集卡上的相机
func EnumDevices(infoStruct *interface{}) []MvCcDeviceInfo {
	/**
	  unsigned int        nDeviceNum;                                 ///< [OUT] \~chinese 在线设备数量           \~english Online Device Number
	  MV_CC_DEVICE_INFO*  pDeviceInfo[MV_MAX_DEVICE_NUM];             ///< [OUT] \~chinese 支持最多256个设备      \~english Support up to 256 devices
	*/
	pstDevList := C.struct__MV_CC_DEVICE_INFO_LIST_{}

	C.MV_CC_EnumDevices(C.uint(0x00000001), (*C.struct__MV_CC_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)))

	var result []MvCcDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	//data := (*MvCcDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0]))
	for i := 0; i < num; i++ {
		result = append(result, *(*MvCcDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}
	//fmt.Println("data", data)

	return result
	//fmt.Println(string(dataNew.SpecialInfo.MvGigeDeviceInfo.chModelName[:]))
	//return 0
}
