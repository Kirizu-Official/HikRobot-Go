package MvsSDK

const MvOK = 0x00000000

type MvErrorCode int

const (
	MvEHandle MvErrorCode = 0x80000000 // MvEHandle: \~chinese 错误或无效的句柄 \~english Error or invalid handle

	MvESupport MvErrorCode = 0x80000001 // MvESupport: \~chinese 不支持的功能 \~english Not supported function

	MvEBufOver MvErrorCode = 0x80000002 // MvEBufOver: \~chinese 缓存已满 \~english Buffer overflow

	MvECallOrder MvErrorCode = 0x80000003 // MvECallOrder: \~chinese 函数调用顺序错误 \~english Function calling order error

	MvEParameter MvErrorCode = 0x80000004 // MvEParameter: \~chinese 错误的参数 \~english Incorrect parameter

	MvEResource MvErrorCode = 0x80000006 // MvEResource: \~chinese 资源申请失败 \~english Applying resource failed

	MvENoData MvErrorCode = 0x80000007 // MvENoData: \~chinese 无数据 \~english No data

	MvEPrecondition MvErrorCode = 0x80000008 // MvEPrecondition: \~chinese 前置条件有误，或运行环境已发生变化 \~english Precondition error, or running environment changed

	MvEVersion MvErrorCode = 0x80000009 // MvEVersion: \~chinese 版本不匹配 \~english Version mismatches

	MvENoEnoughBuf MvErrorCode = 0x8000000A // MvENoEnoughBuf: \~chinese 传入的内存空间不足 \~english Insufficient memory

	MvEAbnormalImage MvErrorCode = 0x8000000B // MvEAbnormalImage: \~chinese 异常图像，可能是丢包导致图像不完整 \~english Abnormal image, maybe incomplete image because of lost packet

	MvELoadLibrary MvErrorCode = 0x8000000C // MvELoadLibrary: \~chinese 动态导入DLL失败 \~english Load library failed

	MvENoOutBuf MvErrorCode = 0x8000000D // MvENoOutBuf: \~chinese 没有可输出的缓存 \~english No Available Buffer

	MvEEncrypt MvErrorCode = 0x8000000E // MvEEncrypt: \~chinese 加密错误 \~english Encryption error

	MvEOpenFile MvErrorCode = 0x8000000F // MvEOpenFile: \~chinese 打开文件出现错误 \~english Open file error

	MvEBufInUse MvErrorCode = 0x80000010 // MvEBufInUse: \~chinese 缓存地址已使用 \~english Buffer already in use

	MvEBufInvalid MvErrorCode = 0x80000011 // MvEBufInvalid: \~chinese 无效的缓存地址 \~english Buffer address invalid

	MvENoAlignBuf MvErrorCode = 0x80000012 // MvENoAlignBuf: \~chinese 缓存对齐异常 \~english Buffer alignment error

	MvENoEnoughBufNum MvErrorCode = 0x80000013 // MvENoEnoughBufNum: \~chinese 缓存个数不足 \~english Insufficient cache count

	MvEPortInUse MvErrorCode = 0x80000014 // MvEPortInUse: \~chinese 串口被占用 \~english Port is in use

	MvEImageDecodec MvErrorCode = 0x80000015 // MvEImageDecodec: \~chinese 解码错误(SDK校验图像异常) \~english Decoding error (SDK verification image exception)

	MvEUnknow MvErrorCode = 0x800000FF // MvEUnknow: \~chinese 未知的错误 \~english Unknown error
)

const (
	MvEGcGeneric MvErrorCode = 0x80000100 // MvEGcGeneric: \~chinese 通用错误 \~english General error

	MvEGcArgument MvErrorCode = 0x80000101 // MvEGcArgument: \~chinese 参数非法 \~english Illegal parameters

	MvEGcRange MvErrorCode = 0x80000102 // MvEGcRange: \~chinese 值超出范围 \~english The value is out of range

	MvEGcProperty MvErrorCode = 0x80000103 // MvEGcProperty: \~chinese 属性 \~english Property

	MvEGcRuntime MvErrorCode = 0x80000104 // MvEGcRuntime: \~chinese 运行环境有问题 \~english Running environment error

	MvEGcLogical MvErrorCode = 0x80000105 // MvEGcLogical: \~chinese 逻辑错误 \~english Logical error

	MvEGcAccess MvErrorCode = 0x80000106 // MvEGcAccess: \~chinese 节点访问条件有误 \~english Node accessing condition error

	MvEGcTimeout MvErrorCode = 0x80000107 // MvEGcTimeout: \~chinese 超时 \~english Timeout

	MvEGcDynamicCast MvErrorCode = 0x80000108 // MvEGcDynamicCast: \~chinese 转换异常 \~english Transformation exception

	MvEGcUnknow MvErrorCode = 0x800001FF // MvEGcUnknow: \~chinese GenICam未知错误 \~english GenICam unknown error
)

const (
	MvENotImplemented MvErrorCode = 0x80000200 // MvENotImplemented: \~chinese 命令不被设备支持 \~english The command is not supported by device

	MvEInvalidAddress MvErrorCode = 0x80000201 // MvEInvalidAddress: \~chinese 访问的目标地址不存在 \~english The target address being accessed does not exist

	MvEWriteProtect MvErrorCode = 0x80000202 // MvEWriteProtect: \~chinese 目标地址不可写 \~english The target address is not writable

	MvEAccessDenied MvErrorCode = 0x80000203 // MvEAccessDenied: \~chinese 设备无访问权限 \~english No permission

	MvEBusy MvErrorCode = 0x80000204 // MvEBusy: \~chinese 设备忙，或网络断开 \~english Device is busy, or network disconnected

	MvEPacket MvErrorCode = 0x80000205 // MvEPacket: \~chinese 网络包数据错误 \~english Network data packet error

	MvENetEr MvErrorCode = 0x80000206 // MvENetEr: \~chinese 网络相关错误 \~english Network error

	MvEKeyVerification MvErrorCode = 0x8000020F // MvEKeyVerification: \~chinese 秘钥校验错误 \~english SwitchKey error

	MvEIpConflict MvErrorCode = 0x80000221 // MvEIpConflict: \~chinese 设备IP冲突 \~english Device IP conflict
)
const (
	MvEUsbRead MvErrorCode = 0x80000300 // MvEUsbRead: \~chinese 读usb出错 \~english Reading USB error

	MvEUsbWrite MvErrorCode = 0x80000301 // MvEUsbWrite: \~chinese 写usb出错 \~english Writing USB error

	MvEUsbDevice MvErrorCode = 0x80000302 // MvEUsbDevice: \~chinese 设备异常 \~english Device exception

	MvEUsbGenicam MvErrorCode = 0x80000303 // MvEUsbGenicam: \~chinese GenICam相关错误 \~english GenICam error

	MvEUsbBandwidth MvErrorCode = 0x80000304 // MvEUsbBandwidth: \~chinese 带宽不足 \~english Insufficient bandwidth

	MvEUsbDriver MvErrorCode = 0x80000305 // MvEUsbDriver: \~chinese 驱动不匹配或者未装驱动 \~english Driver mismatch or unmounted drive

	MvEUsbUnknow MvErrorCode = 0x800003FF // MvEUsbUnknow: \~chinese USB未知的错误 \~english USB unknown error
)
const (
	MvEUpgFileMismatch MvErrorCode = 0x80000400 // MvEUpgFileMismatch: \~chinese 升级固件不匹配 \~english Firmware mismatches

	MvEUpgLanguageMismatch MvErrorCode = 0x80000401 // MvEUpgLanguageMismatch: \~chinese 升级固件语言不匹配 \~english Firmware language mismatches

	MvEUpgConflict MvErrorCode = 0x80000402 // MvEUpgConflict: \~chinese 升级冲突（设备已经在升级了再次请求升级即返回此错误） \~english Upgrading conflicted (repeated upgrading requests during device upgrade)

	MvEUpgInnerErr MvErrorCode = 0x80000403 // MvEUpgInnerErr: \~chinese 升级时设备内部出现错误 \~english Camera internal error during upgrade

	MvEUpgUnknow MvErrorCode = 0x800004FF // MvEUpgUnknow: \~chinese 升级时未知错误 \~english Unknown error during upgrade
)

// GigEVision IP Configuration
const (
	MvIpCfgStatic = 0x05000000 // MvIpCfgStatic: \~chinese 静态 \~english Static
	MvIpCfgDhcp   = 0x06000000 // MvIpCfgDhcp: \~chinese DHCP \~english DHCP
	MvIpCfgLla    = 0x04000000 // MvIpCfgLla: \~chinese LLA \~english LLA
)

// GigEVision Net Transfer Mode
const (
	MvNetTransDriver = 0x00000001 // MvNetTransDriver: \~chinese 驱动 \~english Driver
	MvNetTransSocket = 0x00000002 // MvNetTransSocket: \~chinese Socket \~english Socket
)

// CameraLink Baud Rates (CLUINT32)
const (
	MvCamlBaudrate9600    = 0x00000001 // MvCamlBaudrate9600: \~chinese 9600 \~english 9600
	MvCamlBaudrate19200   = 0x00000002 // MvCamlBaudrate19200: \~chinese 19200 \~english 19200
	MvCamlBaudrate38400   = 0x00000004 // MvCamlBaudrate38400: \~chinese 38400 \~english 38400
	MvCamlBaudrate57600   = 0x00000008 // MvCamlBaudrate57600: \~chinese 57600 \~english 57600
	MvCamlBaudrate115200  = 0x00000010 // MvCamlBaudrate115200: \~chinese 115200 \~english 115200
	MvCamlBaudrate230400  = 0x00000020 // MvCamlBaudrate230400: \~chinese 230400 \~english 230400
	MvCamlBaudrate460800  = 0x00000040 // MvCamlBaudrate460800: \~chinese 460800 \~english 460800
	MvCamlBaudrate921600  = 0x00000080 // MvCamlBaudrate921600: \~chinese 921600 \~english 921600
	MvCamlBaudrateAutomax = 0x40000000 // MvCamlBaudrateAutomax: \~chinese 最大值 \~english Auto Max

)

// Exception message type
const (
	MvExceptionDevDisconnect = 0x00008001 // MvExceptionDevDisconnect: \~chinese 设备断开连接 \~english The device is disconnected
	MvExceptionVersionCheck  = 0x00008002 // MvExceptionVersionCheck: \~chinese SDK与驱动版本不匹配 \~english SDK does not match the driver version
)

// Device Transport Layer Protocol Type
const (
	MvUnknowDevice          = 0x00000000 // MvUnknowDevice: \~chinese 未知设备类型，保留意义 \~english Unknown Device Type, Reserved
	MvGigeDevice            = 0x00000001 // MvGigeDevice: \~chinese GigE设备 \~english GigE Device
	Mv1394Device            = 0x00000002 // Mv1394Device: \~chinese 1394-a/b 设备 \~english 1394-a/b Device
	MvUsbDevice             = 0x00000004 // MvUsbDevice: \~chinese USB 设备 \~english USB Device
	MvCameralinkDevice      = 0x00000008 // MvCameralinkDevice: \~chinese CameraLink设备 \~english CameraLink Device
	MvVirGigeDevice         = 0x00000010 // MvVirGigeDevice: \~chinese 虚拟GigE设备 \~english Virtual GigE Device
	MvVirUsbDevice          = 0x00000020 // MvVirUsbDevice: \~chinese 虚拟USB设备 \~english Virtual USB Device
	MvGentlGigeDevice       = 0x00000040 // MvGentlGigeDevice: \~chinese 自研网卡下GigE设备 \~english GenTL GigE Device
	MvGentlCameralinkDevice = 0x00000080 // MvGentlCameralinkDevice: \~chinese CameraLink相机设备 \~english GenTL CameraLink Camera Device
	MvGentlCxpDevice        = 0x00000100 // MvGentlCxpDevice: \~chinese CoaXPress设备 \~english GenTL CoaXPress Device
	MvGentlXofDevice        = 0x00000200 // MvGentlXofDevice: \~chinese XoF设备 \~english GenTL XoF Device
)

// Device Access Mode
const (
	MvAccessExclusive                  = 1 // MvAccessExclusive: \~chinese 独占权限，其他APP只允许读CCP寄存器 \~english Exclusive authority, other APP is only allowed to read the CCP register
	MvAccessExclusiveWithSwitch        = 2 // MvAccessExclusiveWithSwitch: \~chinese 可以从5模式下抢占权限，然后以独占权限打开 \~english You can seize the authority from the 5 mode, and then open with exclusive authority
	MvAccessControl                    = 3 // MvAccessControl: \~chinese 控制权限，其他APP允许读所有寄存器 \~english Control authority, allows other APP reading all registers
	MvAccessControlWithSwitch          = 4 // MvAccessControlWithSwitch: \~chinese 可以从5的模式下抢占权限，然后以控制权限打开 \~english You can seize the authority from the 5 mode, and then open with control authority
	MvAccessControlSwitchEnable        = 5 // MvAccessControlSwitchEnable: \~chinese 以可被抢占的控制权限打开 \~english Open with seized control authority
	MvAccessControlSwitchEnableWithKey = 6 // MvAccessControlSwitchEnableWithKey: \~chinese 可以从5的模式下抢占权限，然后以可被抢占的控制权限打开 \~english You can seize the authority from the 5 mode, and then open with seized control authority
	MvAccessMonitor                    = 7 // MvAccessMonitor: \~chinese 读模式打开设备，适用于控制权限下 \~english Open with read mode and is available under control authority
)

const MvMaxXmlSymbolicNum = 64
const MvMaxSymbolicLen = 64
const MvMaxNodeNameLen = 64
const MvMaxNodeErrorNum = 64
