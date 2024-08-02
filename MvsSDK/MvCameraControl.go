package MvsSDK

/*
#cgo pkg-config: hik-mvs
#include "MvCameraControl.h"
*/
import "C"
import (
	"fmt"
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
func Initialize() MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_Initialize()))
}

// Finalize 反初始化SDK，释放资源
func Finalize() MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_Finalize()))
}

// EnumerateTls 获取支持的传输层
// return 当前支持的传输层协议类型和数值总和
func EnumerateTls() int32 {
	return int32(C.MV_CC_EnumerateTls())
}

// EnumDevices 枚举设备，支持枚举对应采集卡上的相机
func EnumDevices(nTLayerType uint32) (MvErrorCode, []MvCcDeviceInfo) {
	pstDevList := C.struct__MV_CC_DEVICE_INFO_LIST_{}

	code := MvErrorCode(int32(C.MV_CC_EnumDevices(C.uint(nTLayerType), (*C.struct__MV_CC_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)))))
	var res []MvCcDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCcDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return code, res
}

func EnumDevicesEx(nTLayerType uint32, strManufacturerName string) (MvErrorCode, []MvCcDeviceInfo) {
	pstDevList := C.struct__MV_CC_DEVICE_INFO_LIST_{}

	code := MvErrorCode(int32(C.MV_CC_EnumDevicesEx(C.uint(nTLayerType), (*C.struct__MV_CC_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)), C.CString(strManufacturerName))))
	var res []MvCcDeviceInfo

	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCcDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return code, res
}

func EnumDevicesEx2(nTLayerType uint32, strManufacturerName string, enSortMethod int) (MvErrorCode, []MvCcDeviceInfo) {
	pstDevList := C.struct__MV_CC_DEVICE_INFO_LIST_{}

	code := MvErrorCode(int32(C.MV_CC_EnumDevicesEx2(C.uint(nTLayerType), (*C.struct__MV_CC_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)), C.CString(strManufacturerName), C.MV_SORT_METHOD(enSortMethod))))

	var res []MvCcDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCcDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return code, res
}

func IsDeviceAccessible(pstDevInfo MvCcDeviceInfo, nAccessMode uint32) bool {

	code := byte(C.MV_CC_IsDeviceAccessible((*C.struct__MV_CC_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)), C.uint(nAccessMode)))
	fmt.Println(code)
	return code == 1
}

type Device struct {
	handel unsafe.Pointer
}

func (d *Device) CreateHandle(pstDevInfo MvCcDeviceInfo) MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_CreateHandle(&d.handel, (*C.struct__MV_CC_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)))))
}

func (d *Device) CreateHandleWithoutLog(pstDevInfo MvCcDeviceInfo) MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_CreateHandleWithoutLog(&d.handel, (*C.struct__MV_CC_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)))))
}

func (d *Device) IsDeviceConnected() bool {
	return byte(C.MV_CC_IsDeviceConnected(d.handel)) == 1
}
func (d *Device) OpenDevice(nAccessMode uint32, nSwitchoverKey uint16) MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_OpenDevice(d.handel, C.uint(nAccessMode), C.ushort(nSwitchoverKey))))

}

// MV_ALL_MATCH_INFO Not implementation

// GetDeviceInfo 获取设备信息，返回：MV_ERROR ，MV_CC_DEVICE_INFO
func (d *Device) GetDeviceInfo() (MvErrorCode, MvCcDeviceInfo) {
	var pstDevInfo MvCcDeviceInfo
	code := MvErrorCode(int32(C.MV_CC_CreateHandleWithoutLog(&d.handel, (*C.struct__MV_CC_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)))))
	return code, pstDevInfo
}

func (d *Device) CloseDevice() MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_CloseDevice(d.handel)))
}

func (d *Device) DestroyHandle() MvErrorCode {
	return MvErrorCode(int32(C.MV_CC_DestroyHandle(d.handel)))
}

type DeviceControl struct {
	*Device
}

func (d *Device) NewDeviceControl() *DeviceControl {
	return &DeviceControl{Device: d}
}

func (d *DeviceControl) InvalidateNodes() MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_InvalidateNodes(d.Device.handel)))
	return code
}

func (d *DeviceControl) GetIntValueEx(key string) (MvErrorCode, MvccIntValueEx) {
	var value MvccIntValueEx

	code := MvErrorCode(int32(C.MV_CC_GetIntValueEx(d.Device.handel, C.CString(key), (*C.MVCC_INTVALUE_EX)(unsafe.Pointer(&value)))))
	return code, value
}

func (d *DeviceControl) SetIntValueEx(key string, value int64) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetIntValueEx(d.Device.handel, C.CString(key), C.int64_t(value))))
	return code
}

func (d *DeviceControl) GetEnumValue(key string) (MvErrorCode, MvccEnumValue) {
	var value MvccEnumValue

	code := MvErrorCode(int32(C.MV_CC_GetEnumValue(d.Device.handel, C.CString(key), (*C.MVCC_ENUMVALUE)(unsafe.Pointer(&value)))))
	return code, value
}

func (d *DeviceControl) SetEnumValue(key string, value uint32) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetEnumValue(d.Device.handel, C.CString(key), C.uint(value))))
	return code
}

func (d *DeviceControl) GetEnumEntrySymbolic(key string) (MvErrorCode, MvccEnumEntry) {
	var value MvccEnumEntry

	code := MvErrorCode(int32(C.MV_CC_GetEnumEntrySymbolic(d.Device.handel, C.CString(key), (*C.MVCC_ENUMENTRY)(unsafe.Pointer(&value)))))
	return code, value
}

func (d *DeviceControl) SetEnumValueByString(key string, value string) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetEnumValueByString(d.Device.handel, C.CString(key), C.CString(value))))
	return code
}

func (d *DeviceControl) GetFloatValue(key string) (MvErrorCode, MvccFloatValue) {
	var value MvccFloatValue

	code := MvErrorCode(int32(C.MV_CC_GetFloatValue(d.Device.handel, C.CString(key), (*C.MVCC_FLOATVALUE)(unsafe.Pointer(&value)))))
	return code, value
}

func (d *DeviceControl) SetFloatValue(key string, value float32) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetFloatValue(d.Device.handel, C.CString(key), C.float(value))))
	return code
}

func (d *DeviceControl) GetBoolValue(key string) (MvErrorCode, bool) {
	var value byte

	code := MvErrorCode(int32(C.MV_CC_GetBoolValue(d.Device.handel, C.CString(key), (*C.char)(unsafe.Pointer(&value)))))
	return code, value == 1
}

func (d *DeviceControl) SetBoolValue(key string, value bool) MvErrorCode {
	var wr byte
	if value {
		wr = 1
	} else {
		wr = 0
	}
	code := MvErrorCode(int32(C.MV_CC_SetBoolValue(d.Device.handel, C.CString(key), C.char(wr))))
	return code
}

func (d *DeviceControl) GetStringValue(key string) (MvErrorCode, MvccStringValue) {
	var value MvccStringValue
	code := MvErrorCode(int32(C.MV_CC_GetStringValue(d.Device.handel, C.CString(key), (*C.MVCC_STRINGVALUE)(unsafe.Pointer(&value)))))
	return code, value
}

func (d *DeviceControl) SetStringValue(key string, value string) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetStringValue(d.Device.handel, C.CString(key), C.CString(value))))
	return code
}

func (d *DeviceControl) SetCommandValue(key string) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetCommandValue(d.Device.handel, C.CString(key))))
	return code
}

func (d *DeviceControl) FeatureLoad(fileName string) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_FeatureLoad(d.Device.handel, C.CString(fileName))))
	return code
}

func (d *DeviceControl) FeatureLoadEx(fileName string) MvErrorCode {
	var value MvccNodeErrorList

	code := MvErrorCode(int32(C.MV_CC_FeatureLoadEx(d.Device.handel, C.CString(fileName), (*C.MVCC_NODE_ERROR_LIST)(unsafe.Pointer(&value)))))
	return code
}

func (d *DeviceControl) FeatureSave(fileName string) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_FeatureSave(d.Device.handel, C.CString(fileName))))
	return code
}

func (d *DeviceControl) OpenParamsGUI() MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_OpenParamsGUI(d.Device.handel)))
	return code
}

func (d *DeviceControl) ReadMemory(buf *[]byte, addr int64, size int64) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_ReadMemory(d.Device.handel, unsafe.Pointer(buf), C.longlong(addr), C.longlong(size))))
	return code
}

func (d *DeviceControl) WriteMemory(buf *[]byte, addr int64, size int64) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_WriteMemory(d.Device.handel, unsafe.Pointer(buf), C.longlong(addr), C.longlong(size))))
	return code
}

func (d *DeviceControl) GetGenICamXML(buf *[]byte, size int64) (MvErrorCode, uint32) {
	var readSize uint32
	code := MvErrorCode(int32(C.MV_XML_GetGenICamXML(d.Device.handel, (*C.uchar)(unsafe.Pointer(buf)), C.uint(size), (*C.uint)(unsafe.Pointer(&readSize)))))
	return code, readSize
}

func (d *DeviceControl) GetNodeAccessMode(name string) (MvErrorCode, int) {
	var value uint32
	code := MvErrorCode(int32(C.MV_XML_GetNodeAccessMode(d.Device.handel, C.CString(name), &value)))
	return code, int(value)
}

func (d *DeviceControl) GetNodeInterfaceType(name string) (MvErrorCode, int) {
	var value uint32
	code := MvErrorCode(int32(C.MV_XML_GetNodeInterfaceType(d.Device.handel, C.CString(name), &value)))
	return code, int(value)
}

type DeviceImage struct {
	*Device
}

func (d *Device) NewDeviceImage() *DeviceImage {
	return &DeviceImage{Device: d}
}

func (d *DeviceImage) StartGrabbing() MvErrorCode {

	code := MvErrorCode(int32(C.MV_CC_StartGrabbing(d.Device.handel)))
	return code
}

func (d *DeviceImage) StopGrabbing() MvErrorCode {

	code := MvErrorCode(int32(C.MV_CC_StopGrabbing(d.Device.handel)))
	return code
}

func (d *DeviceImage) GetImageForRGB(data *[]byte, dataSize uint32, timeOut int32) (MvErrorCode, MVFrameOutInfoEx) {
	var pFrameInfo MVFrameOutInfoEx

	code := MvErrorCode(int32(C.MV_CC_GetImageForRGB(d.handel, (*C.uchar)(unsafe.Pointer(data)), C.uint(dataSize), (*C.MV_FRAME_OUT_INFO_EX)(unsafe.Pointer(&pFrameInfo)), C.int(timeOut))))
	return code, pFrameInfo
}

func (d *DeviceImage) GetImageForBGR(data *[]byte, dataSize uint32, timeOut int32) (MvErrorCode, MVFrameOutInfoEx) {
	var pFrameInfo MVFrameOutInfoEx

	code := MvErrorCode(int32(C.MV_CC_GetImageForBGR(d.handel, (*C.uchar)(unsafe.Pointer(data)), C.uint(dataSize), (*C.MV_FRAME_OUT_INFO_EX)(unsafe.Pointer(&pFrameInfo)), C.int(timeOut))))
	return code, pFrameInfo
}

func (d *DeviceImage) GetImageBuffer(waitTime uint32, pFrame *MvFrameOut) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_GetImageBuffer(d.Device.handel, (*C.MV_FRAME_OUT)(unsafe.Pointer(pFrame)), C.uint(waitTime))))
	return code
}

func (d *DeviceImage) FreeImageBuffer(pFrame *MvFrameOut) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_FreeImageBuffer(d.Device.handel, (*C.MV_FRAME_OUT)(unsafe.Pointer(pFrame)))))
	return code
}

func (d *DeviceImage) GetOneFrameTimeout(pData *[MaxFrameSize]byte, pDataSize uint32, timeOut uint32) (MvErrorCode, MVFrameOutInfoEx) {
	var pFrameInfo MVFrameOutInfoEx

	code := MvErrorCode(int32(C.MV_CC_GetOneFrameTimeout(d.Device.handel, (*C.uchar)(unsafe.Pointer(pData)), C.uint(pDataSize), (*C.MV_FRAME_OUT_INFO_EX)(unsafe.Pointer(&pFrameInfo)), C.uint(timeOut))))
	return code, pFrameInfo
}

func (d *DeviceImage) ClearImageBuffer() MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_ClearImageBuffer(d.Device.handel)))
	return code
}

func (d *DeviceImage) GetValidImageNum() (MvErrorCode, uint32) {
	var numRe uint32
	code := MvErrorCode(int32(C.MV_CC_GetValidImageNum(d.Device.handel, (*C.uint)(unsafe.Pointer(&numRe)))))

	return code, numRe
}

func (d *DeviceImage) SetImageNodeNum(num uint32) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetImageNodeNum(d.Device.handel, C.uint(num))))

	return code
}

func (d *DeviceImage) SetGrabStrategy(enGrabStrategy uint32) MvErrorCode {
	code := MvErrorCode(int32(C.MV_CC_SetGrabStrategy(d.Device.handel, (C.MV_GRAB_STRATEGY)(enGrabStrategy))))

	return code
}
