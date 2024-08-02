package MvsSDK

import "C"
import "unsafe"

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

/**
typedef struct _MVCC_INTVALUE_EX_T
{
    int64_t             nCurValue;                                  ///< [OUT] \~chinese 当前值                 \~english Current Value
    int64_t             nMax;                                       ///< [OUT] \~chinese 最大值                 \~english Max
    int64_t             nMin;                                       ///< [OUT] \~chinese 最小值                 \~english Min
    int64_t             nInc;                                       ///< [OUT] \~chinese Inc                    \~english Inc

    unsigned int        nReserved[16];                              ///<       \~chinese 预留                   \~english Reserved

}
*/

type MvccIntValueEx struct {
	CurValue int64
	Max      int64
	Min      int64
	Inc      int64
	Reserved [16]uint32
}

/**
typedef struct _MVCC_ENUMVALUE_T
{
    unsigned int        nCurValue;                                  ///< [OUT] \~chinese 当前值                 \~english Current Value
    unsigned int        nSupportedNum;                              ///< [OUT] \~chinese 数据的有效数据个数     \~english Number of valid data
    unsigned int        nSupportValue[MV_MAX_XML_SYMBOLIC_NUM];     ///< [OUT] \~chinese 支持的枚举值           \~english Support Value

    unsigned int        nReserved[4];                               ///<       \~chinese 预留                   \~english Reserved

}MVCC_ENUMVALUE;
*/

type MvccEnumValue struct {
	CurValue     uint32
	SupportedNum uint32
	SupportValue [MvMaxXmlSymbolicNum]uint32
	Reserved     [4]uint32
}

/**
typedef struct _MVCC_ENUMENTRY_T
{
    unsigned int        nValue;                                     ///< [IN]  \~chinese 指定值                 \~english Value
    char                chSymbolic[MV_MAX_SYMBOLIC_LEN];            ///< [OUT] \~chinese 指定值对应的符号       \~english Symbolic

    unsigned int        nReserved[4];                               ///< \~chinese 预留                         \~english Reserved

}MVCC_ENUMENTRY;
*/

type MvccEnumEntry struct {
	Value    uint32
	Symbolic [MvMaxSymbolicLen]byte
	Reserved [4]uint32
}

/**
typedef struct _MVCC_FLOATVALUE_T
{
    float               fCurValue;                                  ///< [OUT] \~chinese 当前值                 \~english Current Value
    float               fMax;                                       ///< [OUT] \~chinese 最大值                 \~english Max
    float               fMin;                                       ///< [OUT] \~chinese 最小值                 \~english Min

    unsigned int        nReserved[4];                               ///<       \~chinese 预留                   \~english Reserved

}MVCC_FLOATVALUE;
*/

type MvccFloatValue struct {
	CurValue float32
	Max      float32
	Min      float32
	Reserved [4]uint32
}

/**
typedef struct _MVCC_STRINGVALUE_T
{
    char                chCurValue[256];                            ///< [OUT] \~chinese 当前值                 \~english Current Value

    int64_t             nMaxLength;                                 ///< [OUT] \~chinese 最大长度               \~english MaxLength
    unsigned int        nReserved[2];                               ///<       \~chinese 预留                   \~english Reserved

}MVCC_STRINGVALUE;
*/

type MvccStringValue struct {
	CurValue  [256]byte
	MaxLength int64
	Reserved  [2]uint32
}

/**
typedef struct _MVCC_NODE_ERROR_T
{
    char                strName[MV_MAX_NODE_NAME_LEN];              ///< \~chinese 节点名称                     \~english Nodes Name
    MVCC_NODE_ERR_TYPE  enErrType;									///< \~chinese 错误类型                     \~english Error Type

    unsigned int        nReserved[4];                               ///< \~chinese 预留                         \~english Reserved

}MVCC_NODE_ERROR;
*/

type MvccNodeError struct {
	Name     [MvMaxNodeNameLen]byte
	ErrType  int32
	Reserved [4]uint32
}

/**
typedef struct _MVCC_NODE_ERROR_LIST_T
{
    unsigned int        nErrorNum;                                  ///< \~chinese 错误个数                     \~english Number of Error
    MVCC_NODE_ERROR     stNodeError[MV_MAX_NODE_ERROR_NUM];         ///< \~chinese 错误信息                     \~english Error Name

    unsigned int        nReserved[4];                               ///< \~chinese 预留                         \~english Reserved

}MVCC_NODE_ERROR_LIST;
*/

type MvccNodeErrorList struct {
	ErrorNum  uint32
	NodeError [MvMaxNodeErrorNum]MvccNodeError
	Reserved  [4]uint32
}

/**
typedef struct _MV_FRAME_OUT_
{
    unsigned char*          pBufAddr;                               ///< [OUT] \~chinese 图像指针地址           \~english  pointer of image
    MV_FRAME_OUT_INFO_EX    stFrameInfo;                            ///< [OUT] \~chinese 图像信息               \~english information of the specific image

    unsigned int            nRes[16];                               ///<       \~chinese 预留                   \~english Reserved

}MV_FRAME_OUT;
*/

type MvFrameOut struct {
	Addr      unsafe.Pointer
	FrameInfo MVFrameOutInfoEx
	Res       [16]uint32
}

/*
*
typedef struct _MV_FRAME_OUT_INFO_EX_

	{
	    unsigned short          nWidth;                                 ///< [OUT] \~chinese 图像宽(最大65535，超出请用nExtendWidth)    \~english Image Width (over 65535, use nExtendWidth)
	    unsigned short          nHeight;                                ///< [OUT] \~chinese 图像高(最大65535，超出请用nExtendHeight)   \~english Image Height(over 65535, use nExtendHeight)
	    enum MvGvspPixelType    enPixelType;                            ///< [OUT] \~chinese 像素格式               \~english Pixel Type

	    unsigned int            nFrameNum;                              ///< [OUT] \~chinese 帧号                   \~english Frame Number
	    unsigned int            nDevTimeStampHigh;                      ///< [OUT] \~chinese 时间戳高32位           \~english Timestamp high 32 bits
	    unsigned int            nDevTimeStampLow;                       ///< [OUT] \~chinese 时间戳低32位           \~english Timestamp low 32 bits
	    unsigned int            nReserved0;                             ///< [OUT] \~chinese 保留，8字节对齐        \~english Reserved, 8-byte aligned
	    int64_t                 nHostTimeStamp;                         ///< [OUT] \~chinese 主机生成的时间戳       \~english Host-generated timestamp

	    unsigned int            nFrameLen;                              ///< [OUT] \~chinese 帧的长度               \~english The Length of Frame

	    /// \~chinese 设备水印时标      \~english Device frame-specific time scale
	    unsigned int            nSecondCount;                           ///< [OUT] \~chinese 秒数                   \~english The Seconds
	    unsigned int            nCycleCount;                            ///< [OUT] \~chinese 周期数                 \~english The Count of Cycle
	    unsigned int            nCycleOffset;                           ///< [OUT] \~chinese 周期偏移量             \~english The Offset of Cycle

	    float                   fGain;                                  ///< [OUT] \~chinese 增益                   \~english Gain
	    float                   fExposureTime;                          ///< [OUT] \~chinese 曝光时间               \~english Exposure Time
	    unsigned int            nAverageBrightness;                     ///< [OUT] \~chinese 平均亮度               \~english Average brightness

	    /// \~chinese 白平衡相关        \~english White balance
	    unsigned int            nRed;                                   ///< [OUT] \~chinese 红色                   \~english Red
	    unsigned int            nGreen;                                 ///< [OUT] \~chinese 绿色                   \~english Green
	    unsigned int            nBlue;                                  ///< [OUT] \~chinese 蓝色                   \~english Blue

	    unsigned int            nFrameCounter;                          ///< [OUT] \~chinese 总帧数                 \~english Frame Counter
	    unsigned int            nTriggerIndex;                          ///< [OUT] \~chinese 触发计数               \~english Trigger Counting

	    unsigned int            nInput;                                 ///< [OUT] \~chinese 输入                   \~english Input
	    unsigned int            nOutput;                                ///< [OUT] \~chinese 输出                   \~english Output

	    /// \~chinese ROI区域           \~english ROI Region
	    unsigned short          nOffsetX;                               ///< [OUT] \~chinese 水平偏移量             \~english OffsetX
	    unsigned short          nOffsetY;                               ///< [OUT] \~chinese 垂直偏移量             \~english OffsetY
	    unsigned short          nChunkWidth;                            ///< [OUT] \~chinese Chunk宽                \~english The Width of Chunk
	    unsigned short          nChunkHeight;                           ///< [OUT] \~chinese Chunk高                \~english The Height of Chunk

	    unsigned int            nLostPacket;                            ///< [OUT] \~chinese 本帧丢包数             \~english Lost Packet Number In This Frame

	    unsigned int            nUnparsedChunkNum;                      ///< [OUT] \~chinese 未解析的Chunkdata个数  \~english Unparsed Chunk Number
	    union
	    {
	        MV_CHUNK_DATA_CONTENT*  pUnparsedChunkContent;              ///< [OUT] \~chinese 未解析的Chunk          \~english Unparsed Chunk Content
	        int64_t                 nAligning;                          ///< [OUT] \~chinese 校准                   \~english Aligning
	    }UnparsedChunkList;

	    unsigned int            nExtendWidth;                           ///< [OUT] \~chinese 图像宽(扩展变量)       \~english Image Width
	    unsigned int            nExtendHeight;                          ///< [OUT] \~chinese 图像高(扩展变量)       \~english Image Height

	    unsigned int            nReserved[34];                          ///<       \~chinese 预留                   \~english Reserved

}MV_FRAME_OUT_INFO_EX;
*/

type MVFrameOutInfoEx struct {
	Width     uint16
	Height    uint16
	PixelType uint32

	FrameNum         uint32
	DevTimeStampHigh uint32
	DevTimeStampLow  uint32
	Reserved0        uint32
	HostTimeStamp    int64

	FrameLen uint32

	SecondCount uint32
	CycleCount  uint32
	CycleOffset uint32

	Gain              float32
	ExposureTime      float32
	AverageBrightness uint32

	Red   uint32
	Green uint32
	Blue  uint32

	FrameCounter uint32
	TriggerIndex uint32

	Input  uint32
	Output uint32

	OffsetX     uint16
	OffsetY     uint16
	ChunkWidth  uint16
	ChunkHeight uint16

	LostPacket uint32

	UnparsedChunkNum uint32

	UnparsedChunkContent unsafe.Pointer // This is the void* pUnparsedChunkContent in the union
	nAligning            int64          ///< [OUT] \~chinese 校准                   \~english Aligning

	ExtendWidth  uint32
	ExtendHeight uint32

	Reserved [34]uint32
}

/**
typedef struct _MV_CC_PIXEL_CONVERT_PARAM_EX_
{
    unsigned int            nWidth;                                 ///< [IN]  \~chinese 图像宽                 \~english Width
    unsigned int            nHeight;                                ///< [IN]  \~chinese 图像高                 \~english Height

    enum MvGvspPixelType    enSrcPixelType;                         ///< [IN]  \~chinese 源像素格式             \~english Source pixel format
    unsigned char*          pSrcData;                               ///< [IN]  \~chinese 输入数据缓存           \~english Input data buffer
    unsigned int            nSrcDataLen;                            ///< [IN]  \~chinese 输入数据长度           \~english Input data length

    enum MvGvspPixelType    enDstPixelType;                         ///< [IN]  \~chinese 目标像素格式           \~english Destination pixel format
    unsigned char*          pDstBuffer;                             ///< [OUT] \~chinese 输出数据缓存           \~english Output data buffer
    unsigned int            nDstLen;                                ///< [OUT] \~chinese 输出数据长度           \~english Output data length
    unsigned int            nDstBufferSize;                         ///< [IN]  \~chinese 提供的输出缓冲区大小   \~english Provided output buffer size

    unsigned int            nRes[4];                                ///<       \~chinese 预留                   \~english Reserved

}MV_CC_PIXEL_CONVERT_PARAM_EX;
*/

type MVCCPixelConvertParamEx struct {
	Width     uint32
	Height    uint32
	SrcPixel  uint32
	DstPixel  uint32
	SrcData   unsafe.Pointer
	SrcLen    uint32
	DstBuffer unsafe.Pointer
	DstLen    uint32
	DstSize   uint32
	Res       [4]uint32
}
