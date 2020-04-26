// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const KEY_READ REGSAM = 0x20019
const KEY_WRITE REGSAM = 0x20006

const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

const (
	ERROR_NO_MORE_ITEMS = 259
)

const (
	SC_MANAGER_CONNECT            = 0x0001
	SC_MANAGER_CREATE_SERVICE     = 0x0002
	SC_MANAGER_ENUMERATE_SERVICE  = 0x0004
	SC_MANAGER_LOCK               = 0x0008
	SC_MANAGER_QUERY_LOCK_STATUS  = 0x0010
	SC_MANAGER_MODIFY_BOOT_CONFIG = 0x0020
	SC_MANAGER_ALL_ACCESS         = STANDARD_RIGHTS_REQUIRED | SC_MANAGER_CONNECT | SC_MANAGER_CREATE_SERVICE | SC_MANAGER_ENUMERATE_SERVICE | SC_MANAGER_LOCK | SC_MANAGER_QUERY_LOCK_STATUS | SC_MANAGER_MODIFY_BOOT_CONFIG
)

const (
	SERVICE_QUERY_CONFIG         = 0x0001
	SERVICE_CHANGE_CONFIG        = 0x0002
	SERVICE_QUERY_STATUS         = 0x0004
	SERVICE_ENUMERATE_DEPENDENTS = 0x0008
	SERVICE_START                = 0x0010
	SERVICE_STOP                 = 0x0020
	SERVICE_PAUSE_CONTINUE       = 0x0040
	SERVICE_INTERROGATE          = 0x0080
	SERVICE_USER_DEFINED_CONTROL = 0x0100
	SERVICE_ALL_ACCESS           = STANDARD_RIGHTS_REQUIRED | SERVICE_QUERY_CONFIG | SERVICE_CHANGE_CONFIG | SERVICE_QUERY_STATUS | SERVICE_ENUMERATE_DEPENDENTS | SERVICE_START | SERVICE_STOP | SERVICE_PAUSE_CONTINUE | SERVICE_INTERROGATE | SERVICE_USER_DEFINED_CONTROL
)

const (
	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_KERNEL_DRIVER       = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  = 0x00000002
	SERVICE_INTERACTIVE_PROCESS = 0x00000100
)

const (
	SERVICE_BOOT_START   = 0x00000000
	SERVICE_SYSTEM_START = 0x00000001
	SERVICE_AUTO_START   = 0x00000002
	SERVICE_DEMAND_START = 0x00000003
	SERVICE_DISABLED     = 0x00000004
)

const (
	SERVICE_ERROR_IGNORE   = 0x00000000
	SERVICE_ERROR_NORMAL   = 0x00000001
	SERVICE_ERROR_SEVERE   = 0x00000002
	SERVICE_ERROR_CRITICAL = 0x00000003
)

const (
	DELETE                  = 0x00010000
	READ_CONTROL            = 0x00020000
	WRITE_DAC               = 0x00040000
	WRITE_OWNER             = 0x00080000
	SYNCHRONIZE             = 0x00100000
	STANDARD_RIGHTS_READ    = READ_CONTROL
	STANDARD_RIGHTS_WRITE   = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL
	STANDARD_RIGHTS_ALL     = 0x001f0000
	SPECIFIC_RIGHTS_ALL     = 0x0000ffff
)

const (
	SERVICES_ACTIVE_DATABASE = "ServicesActive"
	SERVICES_FAILED_DATABASE = "ServicesFailed"

	SC_GROUP_IDENTIFIER = '+'

	SERVICE_NO_CHANGE = 0xffffffff

	SERVICE_ACTIVE    = 0x00000001
	SERVICE_INACTIVE  = 0x00000002
	SERVICE_STATE_ALL = SERVICE_ACTIVE | SERVICE_INACTIVE
)

const (
	SERVICE_CONTROL_STOP                  = 0x00000001
	SERVICE_CONTROL_PAUSE                 = 0x00000002
	SERVICE_CONTROL_CONTINUE              = 0x00000003
	SERVICE_CONTROL_INTERROGATE           = 0x00000004
	SERVICE_CONTROL_SHUTDOWN              = 0x00000005
	SERVICE_CONTROL_PARAMCHANGE           = 0x00000006
	SERVICE_CONTROL_NETBINDADD            = 0x00000007
	SERVICE_CONTROL_NETBINDREMOVE         = 0x00000008
	SERVICE_CONTROL_NETBINDENABLE         = 0x00000009
	SERVICE_CONTROL_NETBINDDISABLE        = 0x0000000A
	SERVICE_CONTROL_DEVICEEVENT           = 0x0000000B
	SERVICE_CONTROL_HARDWAREPROFILECHANGE = 0x0000000C
	SERVICE_CONTROL_POWEREVENT            = 0x0000000D
	SERVICE_CONTROL_SESSIONCHANGE         = 0x0000000E
	SERVICE_CONTROL_PRESHUTDOWN           = 0x0000000F
	SERVICE_CONTROL_TIMECHANGE            = 0x00000010
	SERVICE_CONTROL_TRIGGEREVENT          = 0x00000020
)

const (
	SERVICE_STOPPED          = 0x00000001
	SERVICE_START_PENDING    = 0x00000002
	SERVICE_STOP_PENDING     = 0x00000003
	SERVICE_RUNNING          = 0x00000004
	SERVICE_CONTINUE_PENDING = 0x00000005
	SERVICE_PAUSE_PENDING    = 0x00000006
	SERVICE_PAUSED           = 0x00000007
)

const (
	SERVICE_ACCEPT_STOP                  = 0x00000001
	SERVICE_ACCEPT_PAUSE_CONTINUE        = 0x00000002
	SERVICE_ACCEPT_SHUTDOWN              = 0x00000004
	SERVICE_ACCEPT_PARAMCHANGE           = 0x00000008
	SERVICE_ACCEPT_NETBINDCHANGE         = 0x00000010
	SERVICE_ACCEPT_HARDWAREPROFILECHANGE = 0x00000020
	SERVICE_ACCEPT_POWEREVENT            = 0x00000040
	SERVICE_ACCEPT_SESSIONCHANGE         = 0x00000080
	SERVICE_ACCEPT_PRESHUTDOWN           = 0x00000100
	SERVICE_ACCEPT_TIMECHANGE            = 0x00000200
	SERVICE_ACCEPT_TRIGGEREVENT          = 0x00000400
)

const (
	SERVICE_RUNS_IN_SYSTEM_PROCESS = 0x00000001

	SERVICE_CONFIG_DESCRIPTION              = 1
	SERVICE_CONFIG_FAILURE_ACTIONS          = 2
	SERVICE_CONFIG_DELAYED_AUTO_START_INFO  = 3
	SERVICE_CONFIG_FAILURE_ACTIONS_FLAG     = 4
	SERVICE_CONFIG_SERVICE_SID_INFO         = 5
	SERVICE_CONFIG_REQUIRED_PRIVILEGES_INFO = 6
	SERVICE_CONFIG_PRESHUTDOWN_INFO         = 7
	SERVICE_CONFIG_TRIGGER_INFO             = 8
	SERVICE_CONFIG_PREFERRED_NODE           = 9

	SERVICE_NOTIFY_STATUS_CHANGE_1 = 1
	SERVICE_NOTIFY_STATUS_CHANGE_2 = 2

	SERVICE_NOTIFY_STATUS_CHANGE = SERVICE_NOTIFY_STATUS_CHANGE_2

	SERVICE_NOTIFY_STOPPED          = 0x00000001
	SERVICE_NOTIFY_START_PENDING    = 0x00000002
	SERVICE_NOTIFY_STOP_PENDING     = 0x00000004
	SERVICE_NOTIFY_RUNNING          = 0x00000008
	SERVICE_NOTIFY_CONTINUE_PENDING = 0x00000010
	SERVICE_NOTIFY_PAUSE_PENDING    = 0x00000020
	SERVICE_NOTIFY_PAUSED           = 0x00000040
	SERVICE_NOTIFY_CREATED          = 0x00000080
	SERVICE_NOTIFY_DELETED          = 0x00000100
	SERVICE_NOTIFY_DELETE_PENDING   = 0x00000200

	SERVICE_STOP_REASON_FLAG_MIN       = 0x00000000
	SERVICE_STOP_REASON_FLAG_UNPLANNED = 0x10000000
	SERVICE_STOP_REASON_FLAG_CUSTOM    = 0x20000000
	SERVICE_STOP_REASON_FLAG_PLANNED   = 0x40000000
	SERVICE_STOP_REASON_FLAG_MAX       = 0x80000000

	SERVICE_STOP_REASON_MAJOR_MIN             = 0x00000000
	SERVICE_STOP_REASON_MAJOR_OTHER           = 0x00010000
	SERVICE_STOP_REASON_MAJOR_HARDWARE        = 0x00020000
	SERVICE_STOP_REASON_MAJOR_OPERATINGSYSTEM = 0x00030000
	SERVICE_STOP_REASON_MAJOR_SOFTWARE        = 0x00040000
	SERVICE_STOP_REASON_MAJOR_APPLICATION     = 0x00050000
	SERVICE_STOP_REASON_MAJOR_NONE            = 0x00060000
	SERVICE_STOP_REASON_MAJOR_MAX             = 0x00070000
	SERVICE_STOP_REASON_MAJOR_MIN_CUSTOM      = 0x00400000
	SERVICE_STOP_REASON_MAJOR_MAX_CUSTOM      = 0x00ff0000

	SERVICE_STOP_REASON_MINOR_MIN                       = 0x00000000
	SERVICE_STOP_REASON_MINOR_OTHER                     = 0x00000001
	SERVICE_STOP_REASON_MINOR_MAINTENANCE               = 0x00000002
	SERVICE_STOP_REASON_MINOR_INSTALLATION              = 0x00000003
	SERVICE_STOP_REASON_MINOR_UPGRADE                   = 0x00000004
	SERVICE_STOP_REASON_MINOR_RECONFIG                  = 0x00000005
	SERVICE_STOP_REASON_MINOR_HUNG                      = 0x00000006
	SERVICE_STOP_REASON_MINOR_UNSTABLE                  = 0x00000007
	SERVICE_STOP_REASON_MINOR_DISK                      = 0x00000008
	SERVICE_STOP_REASON_MINOR_NETWORKCARD               = 0x00000009
	SERVICE_STOP_REASON_MINOR_ENVIRONMENT               = 0x0000000a
	SERVICE_STOP_REASON_MINOR_HARDWARE_DRIVER           = 0x0000000b
	SERVICE_STOP_REASON_MINOR_OTHERDRIVER               = 0x0000000c
	SERVICE_STOP_REASON_MINOR_SERVICEPACK               = 0x0000000d
	SERVICE_STOP_REASON_MINOR_SOFTWARE_UPDATE           = 0x0000000e
	SERVICE_STOP_REASON_MINOR_SECURITYFIX               = 0x0000000f
	SERVICE_STOP_REASON_MINOR_SECURITY                  = 0x00000010
	SERVICE_STOP_REASON_MINOR_NETWORK_CONNECTIVITY      = 0x00000011
	SERVICE_STOP_REASON_MINOR_WMI                       = 0x00000012
	SERVICE_STOP_REASON_MINOR_SERVICEPACK_UNINSTALL     = 0x00000013
	SERVICE_STOP_REASON_MINOR_SOFTWARE_UPDATE_UNINSTALL = 0x00000014
	SERVICE_STOP_REASON_MINOR_SECURITYFIX_UNINSTALL     = 0x00000015
	SERVICE_STOP_REASON_MINOR_MMC                       = 0x00000016
	SERVICE_STOP_REASON_MINOR_NONE                      = 0x00000017
	SERVICE_STOP_REASON_MINOR_MAX                       = 0x00000018
	SERVICE_STOP_REASON_MINOR_MIN_CUSTOM                = 0x00000100
	SERVICE_STOP_REASON_MINOR_MAX_CUSTOM                = 0x0000FFFF

	SERVICE_CONTROL_STATUS_REASON_INFO = 1

	SERVICE_SID_TYPE_NONE         = 0x00000000
	SERVICE_SID_TYPE_UNRESTRICTED = 0x00000001
	SERVICE_SID_TYPE_RESTRICTED   = 0x00000002 | SERVICE_SID_TYPE_UNRESTRICTED

	SERVICE_TRIGGER_TYPE_DEVICE_INTERFACE_ARRIVAL = 1
	SERVICE_TRIGGER_TYPE_IP_ADDRESS_AVAILABILITY  = 2
	SERVICE_TRIGGER_TYPE_DOMAIN_JOIN              = 3
	SERVICE_TRIGGER_TYPE_FIREWALL_PORT_EVENT      = 4
	SERVICE_TRIGGER_TYPE_GROUP_POLICY             = 5
	SERVICE_TRIGGER_TYPE_CUSTOM                   = 20

	SERVICE_TRIGGER_DATA_TYPE_BINARY = 1
	SERVICE_TRIGGER_DATA_TYPE_STRING = 2
)

type (
	ACCESS_MASK uint32
	HKEY        HANDLE
	REGSAM      ACCESS_MASK
)

type SERVICE_DESCRIPTION struct {
	LpDescription *uint16
}

type SERVICE_STATUS struct {
	DwServiceType             uint32
	DwCurrentState            uint32
	DwControlsAccepted        uint32
	DwWin32ExitCode           uint32
	DwServiceSpecificExitCode uint32
	DwCheckPoint              uint32
	DwWaitHint                uint32
}

type SERVICE_TABLE_ENTRY struct {
	LpServiceName *uint16
	LpServiceProc uintptr
}

const (
	REG_NONE      uint64 = 0 // No value type
	REG_SZ               = 1 // Unicode nul terminated string
	REG_EXPAND_SZ        = 2 // Unicode nul terminated string
	// (with environment variable references)
	REG_BINARY                     = 3 // Free form binary
	REG_DWORD                      = 4 // 32-bit number
	REG_DWORD_LITTLE_ENDIAN        = 4 // 32-bit number (same as REG_DWORD)
	REG_DWORD_BIG_ENDIAN           = 5 // 32-bit number
	REG_LINK                       = 6 // Symbolic Link (unicode)
	REG_MULTI_SZ                   = 7 // Multiple Unicode strings
	REG_RESOURCE_LIST              = 8 // Resource list in the resource map
	REG_FULL_RESOURCE_DESCRIPTOR   = 9 // Resource list in the hardware description
	REG_RESOURCE_REQUIREMENTS_LIST = 10
	REG_QWORD                      = 11 // 64-bit number
	REG_QWORD_LITTLE_ENDIAN        = 11 // 64-bit number (same as REG_QWORD)
)

var (
	// Library
	libadvapi32 *windows.LazyDLL

	// Functions
	regCloseKey                  *windows.LazyProc
	regOpenKeyEx                 *windows.LazyProc
	regQueryValueEx              *windows.LazyProc
	regEnumValue                 *windows.LazyProc
	regSetValueEx                *windows.LazyProc
	openSCManager                *windows.LazyProc
	createService                *windows.LazyProc
	openService                  *windows.LazyProc
	deleteService                *windows.LazyProc
	closeServiceHandle           *windows.LazyProc
	lockServiceDatabase          *windows.LazyProc
	unlockServiceDatabase        *windows.LazyProc
	changeServiceConfig2         *windows.LazyProc
	startService                 *windows.LazyProc
	controlService               *windows.LazyProc
	queryServiceStatus           *windows.LazyProc
	startServiceCtrlDispatcher   *windows.LazyProc
	registerServiceCtrlHandlerEx *windows.LazyProc
	setServiceStatus             *windows.LazyProc
)

func init() {
	// Library
	libadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	// Functions
	regCloseKey = libadvapi32.NewProc("RegCloseKey")
	regOpenKeyEx = libadvapi32.NewProc("RegOpenKeyExW")
	regQueryValueEx = libadvapi32.NewProc("RegQueryValueExW")
	regEnumValue = libadvapi32.NewProc("RegEnumValueW")
	regSetValueEx = libadvapi32.NewProc("RegSetValueExW")
	openSCManager = libadvapi32.NewProc("OpenSCManagerW")
	createService = libadvapi32.NewProc("CreateServiceW")
	openService = libadvapi32.NewProc("OpenServiceW")
	deleteService = libadvapi32.NewProc("DeleteService")
	closeServiceHandle = libadvapi32.NewProc("CloseServiceHandle")
	lockServiceDatabase = libadvapi32.NewProc("LockServiceDatabase")
	unlockServiceDatabase = libadvapi32.NewProc("UnlockServiceDatabase")
	changeServiceConfig2 = libadvapi32.NewProc("ChangeServiceConfig2W")
	startService = libadvapi32.NewProc("StartServiceW")
	controlService = libadvapi32.NewProc("ControlService")
	queryServiceStatus = libadvapi32.NewProc("QueryServiceStatus")
	startServiceCtrlDispatcher = libadvapi32.NewProc("StartServiceCtrlDispatcherW")
	registerServiceCtrlHandlerEx = libadvapi32.NewProc("RegisterServiceCtrlHandlerExW")
	setServiceStatus = libadvapi32.NewProc("SetServiceStatus")
}

func RegCloseKey(hKey HKEY) int32 {
	ret, _, _ := syscall.Syscall(regCloseKey.Addr(), 1,
		uintptr(hKey),
		0,
		0)

	return int32(ret)
}

func RegOpenKeyEx(hKey HKEY, lpSubKey *uint16, ulOptions uint32, samDesired REGSAM, phkResult *HKEY) int32 {
	ret, _, _ := syscall.Syscall6(regOpenKeyEx.Addr(), 5,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpSubKey)),
		uintptr(ulOptions),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(phkResult)),
		0)

	return int32(ret)
}

func RegQueryValueEx(hKey HKEY, lpValueName *uint16, lpReserved, lpType *uint32, lpData *byte, lpcbData *uint32) int32 {
	ret, _, _ := syscall.Syscall6(regQueryValueEx.Addr(), 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)))

	return int32(ret)
}

func RegEnumValue(hKey HKEY, index uint32, lpValueName *uint16, lpcchValueName *uint32, lpReserved, lpType *uint32, lpData *byte, lpcbData *uint32) int32 {
	ret, _, _ := syscall.Syscall9(regEnumValue.Addr(), 8,
		uintptr(hKey),
		uintptr(index),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(unsafe.Pointer(lpcchValueName)),
		uintptr(unsafe.Pointer(lpReserved)),
		uintptr(unsafe.Pointer(lpType)),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(unsafe.Pointer(lpcbData)),
		0)
	return int32(ret)
}

func RegSetValueEx(hKey HKEY, lpValueName *uint16, lpReserved, lpDataType uint64, lpData *byte, cbData uint32) int32 {
	ret, _, _ := syscall.Syscall6(regSetValueEx.Addr(), 6,
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(lpReserved),
		uintptr(lpDataType),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(cbData))
	return int32(ret)
}

func OpenSCManager(lpMachineName, lpDatabaseName *string, dwDesiredAccess uint32) HANDLE {
	machineName := UTF16PtrFromString(lpMachineName)
	databaseName := UTF16PtrFromString(lpDatabaseName)

	ret, _, _ := syscall.Syscall(openSCManager.Addr(), 3,
		uintptr(unsafe.Pointer(machineName)),
		uintptr(unsafe.Pointer(databaseName)),
		uintptr(dwDesiredAccess))
	return HANDLE(ret)
}

func CreateService(hSCManager HANDLE, lpServiceName, lpDisplayName string, dwDesiredAccess, dwServiceType, dwStartType, dwErrorControl uint32,
	lpBinaryPathName string, lpLoadOrderGroup *string, lpdwTagId *uint32, lpDependencies, lpServiceStartName, lpPassword *string) HANDLE {
	serviceName, _ := syscall.UTF16PtrFromString(lpServiceName)
	displayName, _ := syscall.UTF16PtrFromString(lpDisplayName)
	binaryPathName, _ := syscall.UTF16PtrFromString(lpBinaryPathName)
	loadOrderGroup := UTF16PtrFromString(lpLoadOrderGroup)
	dependencies := UTF16PtrFromString(lpDependencies)
	serviceStartName := UTF16PtrFromString(lpServiceStartName)
	password := UTF16PtrFromString(lpPassword)

	ret, _, _ := syscall.Syscall15(createService.Addr(), 13,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(serviceName)),
		uintptr(unsafe.Pointer(displayName)),
		uintptr(dwDesiredAccess),
		uintptr(dwServiceType),
		uintptr(dwStartType),
		uintptr(dwErrorControl),
		uintptr(unsafe.Pointer(binaryPathName)),
		uintptr(unsafe.Pointer(loadOrderGroup)),
		uintptr(unsafe.Pointer(lpdwTagId)),
		uintptr(unsafe.Pointer(dependencies)),
		uintptr(unsafe.Pointer(serviceStartName)),
		uintptr(unsafe.Pointer(password)),
		0,
		0)
	return HANDLE(ret)
}

func OpenService(hSCManager HANDLE, lpServiceName string, dwDesiredAccess uint32) HANDLE {
	serviceName, _ := syscall.UTF16PtrFromString(lpServiceName)

	ret, _, _ := syscall.Syscall(openService.Addr(), 3,
		uintptr(hSCManager),
		uintptr(unsafe.Pointer(serviceName)),
		uintptr(dwDesiredAccess))
	return HANDLE(ret)
}

func DeleteService(hService HANDLE) BOOL {
	ret, _, _ := syscall.Syscall(deleteService.Addr(), 1,
		uintptr(hService),
		0,
		0)
	return BOOL(ret)
}

func CloseServiceHandle(hSCObject HANDLE) BOOL {
	ret, _, _ := syscall.Syscall(closeServiceHandle.Addr(), 1,
		uintptr(hSCObject),
		0,
		0)
	return BOOL(ret)
}

func LockServiceDatabase(hSCManager HANDLE) HANDLE {
	ret, _, _ := syscall.Syscall(lockServiceDatabase.Addr(), 1,
		uintptr(hSCManager),
		0,
		0)
	return HANDLE(ret)
}

func UnlockServiceDatabase(ScLock HANDLE) BOOL {
	ret, _, _ := syscall.Syscall(unlockServiceDatabase.Addr(), 1,
		uintptr(ScLock),
		0,
		0)
	return BOOL(ret)
}

func ChangeServiceConfig2(hService HANDLE, dwInfoLevel uint32, lpInfo uintptr) BOOL {
	ret, _, _ := syscall.Syscall(changeServiceConfig2.Addr(), 3,
		uintptr(hService),
		uintptr(dwInfoLevel),
		uintptr(lpInfo))
	return BOOL(ret)
}

func StartService(hService HANDLE, dwNumServiceArgs uint32, lpServiceArgVectors *string) BOOL {
	serviceArgVectors := UTF16PtrFromString(lpServiceArgVectors)

	ret, _, _ := syscall.Syscall(startService.Addr(), 3,
		uintptr(hService),
		uintptr(dwNumServiceArgs),
		uintptr(unsafe.Pointer(serviceArgVectors)))
	return BOOL(ret)
}

func ControlService(hService HANDLE, dwControl uint32, lpServiceStatus *SERVICE_STATUS) BOOL {
	ret, _, _ := syscall.Syscall(controlService.Addr(), 3,
		uintptr(hService),
		uintptr(dwControl),
		uintptr(unsafe.Pointer(lpServiceStatus)))
	return BOOL(ret)
}

func QueryServiceStatus(hService HANDLE, lpServiceStatus *SERVICE_STATUS) BOOL {
	ret, _, _ := syscall.Syscall(queryServiceStatus.Addr(), 2,
		uintptr(hService),
		uintptr(unsafe.Pointer(lpServiceStatus)),
		0)
	return BOOL(ret)
}

func StartServiceCtrlDispatcher(lpServiceStartTable *SERVICE_TABLE_ENTRY) BOOL {
	ret, _, _ := syscall.Syscall(startServiceCtrlDispatcher.Addr(), 1,
		uintptr(unsafe.Pointer(lpServiceStartTable)),
		0,
		0)
	return BOOL(ret)
}

func RegisterServiceCtrlHandlerEx(lpServiceName string, lpHandlerProc uintptr, lpContext uintptr) HANDLE {
	serviceName, _ := syscall.UTF16PtrFromString(lpServiceName)

	ret, _, _ := syscall.Syscall(registerServiceCtrlHandlerEx.Addr(), 3,
		uintptr(unsafe.Pointer(serviceName)),
		lpHandlerProc,
		lpContext)
	return HANDLE(ret)
}

func SetServiceStatus(hServiceStatus HANDLE, lpServiceStatus *SERVICE_STATUS) BOOL {
	ret, _, _ := syscall.Syscall(setServiceStatus.Addr(), 2,
		uintptr(hServiceStatus),
		uintptr(unsafe.Pointer(lpServiceStatus)),
		0)
	return BOOL(ret)
}
