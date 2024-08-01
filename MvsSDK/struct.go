package MvsSDK

import "C"

const MvMaxDeviceNum = 256
const InfoMaxBufferSize = 64

type MvCcDeviceInfo struct {
	MajorVer    uint16
	MinorVer    uint16
	MacAddrHigh uint32
	MacAddrLow  uint32
	TLayerType  uint32
	DevTypeInfo uint32
	Reserved    [3]uint32
	SpecialInfo sSpecialInfo
}

type sSpecialInfo struct {
	MvGigeDeviceInfo MvGigeDeviceInfo ///< [OUT] \~chinese GigE设备信息              \~english GigE Device Info
	MvUsb3DeviceInfo MvUsb3DeviceInfo ///< [OUT] \~chinese USB设备信息               \~english USB Device Info
	MvCamlDevInfo    MvCamlDevInfo    ///< [OUT] \~chinese CameraLink设备信息        \~english CameraLink Device Info
	MvCmlDeviceInfo  MvCmlDeviceInfo  ///< [OUT] \~chinese 采集卡CameraLink设备信息     \~english CameraLink Device Info On Frame Grabber
	MvCxpDeviceInfo  MvCxpDeviceInfo  ///< [OUT] \~chinese 采集卡CoaXPress设备信息     \~english CoaXPress Device Info On Frame Grabber
	MvXofDeviceInfo  MvXofDeviceInfo  ///< [OUT] \~chinese 采集卡XoF设备信息          \~english XoF Device Info On Frame Grabber
}

//unsigned int        nIpCfgOption;                               ///< [OUT] \~chinese IP配置选项             \~english IP Configuration Options
//unsigned int        nIpCfgCurrent;                              ///< [OUT] \~chinese 当前IP配置             \~english IP Configuration
//unsigned int        nCurrentIp;                                 ///< [OUT] \~chinese 当前IP地址             \~english Current Ip
//unsigned int        nCurrentSubNetMask;                         ///< [OUT] \~chinese 当前子网掩码           \~english Curtent Subnet Mask
//unsigned int        nDefultGateWay;                             ///< [OUT] \~chinese 当前网关               \~english Current Gateway
//unsigned char       chManufacturerName[32];                     ///< [OUT] \~chinese 制造商名称             \~english Manufacturer Name
//unsigned char       chModelName[32];                            ///< [OUT] \~chinese 型号名称               \~english Model Name
//unsigned char       chDeviceVersion[32];                        ///< [OUT] \~chinese 设备版本               \~english Device Version
//unsigned char       chManufacturerSpecificInfo[48];             ///< [OUT] \~chinese 制造商的具体信息       \~english Manufacturer Specific Information
//unsigned char       chSerialNumber[16];                         ///< [OUT] \~chinese 序列号                 \~english Serial Number
//unsigned char       chUserDefinedName[16];                      ///< [OUT] \~chinese 用户自定义名称         \~english User Defined Name
//unsigned int        nNetExport;                                 ///< [OUT] \~chinese 网口IP地址             \~english NetWork IP Address
//
//unsigned int        nReserved[4];                               ///<       \~chinese 预留                   \~english Reserved

type MvGigeDeviceInfo struct {
	IpCfgOption              uint32
	IpCfgCurrent             uint32
	CurrentIp                uint32
	CurrentSubNetMask        uint32
	DefultGateWay            uint32
	ManufacturerName         [32]byte
	ModelName                [32]byte
	DeviceVersion            [32]byte
	ManufacturerSpecificInfo [48]byte
	SerialNumber             [16]byte
	UserDefinedName          [16]byte
	NetExport                uint32
	Reserved                 [4]uint32
}

//unsigned char       CrtlInEndPoint;                             ///< [OUT] \~chinese 控制输入端点           \~english Control input endpoint
//unsigned char       CrtlOutEndPoint;                            ///< [OUT] \~chinese 控制输出端点           \~english Control output endpoint
//unsigned char       StreamEndPoint;                             ///< [OUT] \~chinese 流端点                 \~english Flow endpoint
//unsigned char       EventEndPoint;                              ///< [OUT] \~chinese 事件端点               \~english Event endpoint
//unsigned short      idVendor;                                   ///< [OUT] \~chinese 供应商ID号             \~english Vendor ID Number
//unsigned short      idProduct;                                  ///< [OUT] \~chinese 产品ID号               \~english Device ID Number
//unsigned int        nDeviceNumber;                              ///< [OUT] \~chinese 设备索引号             \~english Device Number
//unsigned char       chDeviceGUID[INFO_MAX_BUFFER_SIZE];         ///< [OUT] \~chinese 设备GUID号             \~english Device GUID Number
//unsigned char       chVendorName[INFO_MAX_BUFFER_SIZE];         ///< [OUT] \~chinese 供应商名字             \~english Vendor Name
//unsigned char       chModelName[INFO_MAX_BUFFER_SIZE];          ///< [OUT] \~chinese 型号名字               \~english Model Name
//unsigned char       chFamilyName[INFO_MAX_BUFFER_SIZE];         ///< [OUT] \~chinese 家族名字               \~english Family Name
//unsigned char       chDeviceVersion[INFO_MAX_BUFFER_SIZE];      ///< [OUT] \~chinese 设备版本               \~english Device Version
//unsigned char       chManufacturerName[INFO_MAX_BUFFER_SIZE];   ///< [OUT] \~chinese 制造商名字             \~english Manufacturer Name
//unsigned char       chSerialNumber[INFO_MAX_BUFFER_SIZE];       ///< [OUT] \~chinese 序列号                 \~english Serial Number
//unsigned char       chUserDefinedName[INFO_MAX_BUFFER_SIZE];    ///< [OUT] \~chinese 用户自定义名字         \~english User Defined Name
//unsigned int        nbcdUSB;                                    ///< [OUT] \~chinese 支持的USB协议          \~english Support USB Protocol
//unsigned int        nDeviceAddress;                             ///< [OUT] \~chinese 设备地址               \~english Device Address
//unsigned int        nReserved[2];                               ///<       \~chinese 预留                   \~english Reserved

type MvUsb3DeviceInfo struct {
	CrtlInEndPoint   uint8
	CrtlOutEndPoint  uint8
	StreamEndPoint   uint8
	EventEndPoint    uint8
	IdVendor         uint16
	IdProduct        uint16
	DeviceNumber     uint32
	DeviceGUID       [InfoMaxBufferSize]byte
	VendorName       [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	FamilyName       [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	ManufacturerName [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	UserDefinedName  [InfoMaxBufferSize]byte
	BcdUSB           uint32
	DeviceAddress    uint32
	Reserved         [2]uint32
}

//typedef struct _MV_CamL_DEV_INFO_
//{
//unsigned char       chPortID[INFO_MAX_BUFFER_SIZE];             ///< [OUT] \~chinese 串口号                 \~english Port ID
//unsigned char       chModelName[INFO_MAX_BUFFER_SIZE];          ///< [OUT] \~chinese 型号名字               \~english Model Name
//unsigned char       chFamilyName[INFO_MAX_BUFFER_SIZE];         ///< [OUT] \~chinese 名称                   \~english Family Name
//unsigned char       chDeviceVersion[INFO_MAX_BUFFER_SIZE];      ///< [OUT] \~chinese 设备版本               \~english Device Version
//unsigned char       chManufacturerName[INFO_MAX_BUFFER_SIZE];   ///< [OUT] \~chinese 制造商名字             \~english Manufacturer Name
//unsigned char       chSerialNumber[INFO_MAX_BUFFER_SIZE];       ///< [OUT] \~chinese 序列号                 \~english Serial Number
//
//unsigned int        nReserved[38];                              ///<       \~chinese 预留                   \~english Reserved
//
//}MV_CamL_DEV_INFO;
//

type MvCamlDevInfo struct {
	PortID           [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	FamilyName       [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	ManufacturerName [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	Reserved         [38]uint32
}

//typedef struct _MV_CML_DEVICE_INFO_
//{
//unsigned char       chInterfaceID[INFO_MAX_BUFFER_SIZE];     ///  \~chinese 采集卡ID       \~english Interface ID of Frame Grabber
//unsigned char       chVendorName[INFO_MAX_BUFFER_SIZE];      ///< \~chinese 供应商名字       \~english Vendor name
//unsigned char       chModelName[INFO_MAX_BUFFER_SIZE];       ///< \~chinese 型号名字         \~english Model name
//unsigned char       chManufacturerInfo[INFO_MAX_BUFFER_SIZE];///< \~chinese 厂商信息         \~english Manufacturer information
//unsigned char       chDeviceVersion[INFO_MAX_BUFFER_SIZE];   ///< \~chinese 相机版本         \~english Device version
//unsigned char       chSerialNumber[INFO_MAX_BUFFER_SIZE];    ///< \~chinese 序列号           \~english Serial number
//unsigned char       chUserDefinedName[INFO_MAX_BUFFER_SIZE]; ///< \~chinese 用户自定义名字   \~english User defined name
//unsigned char       chDeviceID[INFO_MAX_BUFFER_SIZE];        ///< \~chinese 相机ID            \~english Device ID
//unsigned int        nReserved[7];                              ///< \~chinese 保留字段      \~english Reserved
//}MV_CML_DEVICE_INFO;

type MvCmlDeviceInfo struct {
	InterfaceID      [InfoMaxBufferSize]byte
	VendorName       [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	ManufacturerInfo [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	UserDefinedName  [InfoMaxBufferSize]byte
	DeviceID         [InfoMaxBufferSize]byte
	Reserved         [7]uint32
}

//typedef struct _MV_CXP_DEVICE_INFO_
//{
//unsigned char       chInterfaceID[INFO_MAX_BUFFER_SIZE];     ///  \~chinese 采集卡ID       \~english Interface ID of Frame Grabber
//unsigned char       chVendorName[INFO_MAX_BUFFER_SIZE];      ///< \~chinese 供应商名字       \~english Vendor name
//unsigned char       chModelName[INFO_MAX_BUFFER_SIZE];       ///< \~chinese 型号名字         \~english Model name
//unsigned char       chManufacturerInfo[INFO_MAX_BUFFER_SIZE];///< \~chinese 厂商信息         \~english Manufacturer information
//unsigned char       chDeviceVersion[INFO_MAX_BUFFER_SIZE];   ///< \~chinese 相机版本         \~english Device version
//unsigned char       chSerialNumber[INFO_MAX_BUFFER_SIZE];    ///< \~chinese 序列号           \~english Serial number
//unsigned char       chUserDefinedName[INFO_MAX_BUFFER_SIZE]; ///< \~chinese 用户自定义名字   \~english User defined name
//unsigned char       chDeviceID[INFO_MAX_BUFFER_SIZE];        ///< \~chinese 相机ID            \~english Device ID
//unsigned int        nReserved[7];                              ///< \~chinese 保留字段      \~english Reserved
//}MV_CXP_DEVICE_INFO;

type MvCxpDeviceInfo struct {
	InterfaceID      [InfoMaxBufferSize]byte
	VendorName       [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	ManufacturerInfo [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	UserDefinedName  [InfoMaxBufferSize]byte
	DeviceID         [InfoMaxBufferSize]byte
	Reserved         [7]uint32
}

//typedef struct _MV_XOF_DEVICE_INFO_
//{
//unsigned char       chInterfaceID[INFO_MAX_BUFFER_SIZE];     ///  \~chinese 采集卡ID       \~english Interface ID of Frame Grabber
//unsigned char       chVendorName[INFO_MAX_BUFFER_SIZE];      ///< \~chinese 供应商名字       \~english Vendor name
//unsigned char       chModelName[INFO_MAX_BUFFER_SIZE];       ///< \~chinese 型号名字         \~english Model name
//unsigned char       chManufacturerInfo[INFO_MAX_BUFFER_SIZE];///< \~chinese 厂商信息         \~english Manufacturer information
//unsigned char       chDeviceVersion[INFO_MAX_BUFFER_SIZE];   ///< \~chinese 相机版本         \~english Device version
//unsigned char       chSerialNumber[INFO_MAX_BUFFER_SIZE];    ///< \~chinese 序列号           \~english Serial number
//unsigned char       chUserDefinedName[INFO_MAX_BUFFER_SIZE]; ///< \~chinese 用户自定义名字   \~english User defined name
//unsigned char       chDeviceID[INFO_MAX_BUFFER_SIZE];        ///< \~chinese 相机ID            \~english Device ID
//unsigned int        nReserved[7];                              ///< \~chinese 保留字段      \~english Reserved
//}MV_XOF_DEVICE_INFO;

type MvXofDeviceInfo struct {
	InterfaceID      [InfoMaxBufferSize]byte
	VendorName       [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	ManufacturerInfo [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	UserDefinedName  [InfoMaxBufferSize]byte
	DeviceID         [InfoMaxBufferSize]byte
	Reserved         [7]uint32
}
