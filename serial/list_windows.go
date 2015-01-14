package serial

import (
	"syscall"
	"unsafe"
)

func listInternal() []SerialPort {
	var handle syscall.Handle
	subkey, _ := syscall.UTF16PtrFromString("HARDWARE\\DEVICEMAP\\SERIALCOMM")
	if syscall.RegOpenKeyEx(syscall.HKEY_LOCAL_MACHINE, subkey, 0, syscall.KEY_READ, &handle) != nil {
		return nil
	}
	defer syscall.RegCloseKey(handle)

	// Gets number of ports
	var count uint32
	var maxNameLen uint32
	var maxValueLen uint32
	if syscall.RegQueryInfoKey(handle, nil, nil, nil, nil, nil, nil, &count, &maxNameLen, &maxValueLen, nil, nil) != nil {
		return nil
	}

	list := make([]SerialPort, count)

	modadvapi32 := syscall.NewLazyDLL("advapi32.dll")
	procRegEnumValueW := modadvapi32.NewProc("RegEnumValueW")
	
	var i uint32
	for i = 0; i < count; i++ {
		nameLen := maxNameLen + 1
		valueLen := maxValueLen + 1
		name := make([]uint16, nameLen)
		value := make([]uint16, valueLen)

		//RegEnumValue
		r0, _, _ := syscall.Syscall9(procRegEnumValueW.Addr(), 8, uintptr(handle), uintptr(i), uintptr(unsafe.Pointer(&name[0])), uintptr(unsafe.Pointer(&nameLen)), 0, 0, uintptr(unsafe.Pointer(&value[0])), uintptr(unsafe.Pointer(&valueLen)), 0)
		if r0 != 0 {
			return nil
		}

		list[i] = SerialPort{syscall.UTF16ToString(value), "\\\\.\\" + syscall.UTF16ToString(value)}
	}

	return list
}
