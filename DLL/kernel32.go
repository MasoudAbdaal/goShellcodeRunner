package dll

import (
	"syscall"
)

// TradeOff: While you using struct and methods, you should define a constructor for it
// Every time you want to use Kernel32, you have to New() it & it takes more memory area in comparesion of use:
// var Kernel32 *syscall.LazyDLL = syscall.NewLazyDLL("kernel32.dll")
type kernel32DLL struct {
}

func (k *kernel32DLL) newKernel32DLL() *kernel32DLL {
	return k
}

func (*kernel32DLL) VirtualAllocEx(
	processHandle *syscall.Handle,
	memorySize int,
	allocationType int,
	regionProtections int) (address uintptr, lastError error) {

	proc := Kernel32.NewProc("VirtualAllocEx")

	addr, _, lastErr := proc.Call(uintptr(*processHandle),
		uintptr(0),
		uintptr(memorySize),
		uintptr(allocationType),
		uintptr(regionProtections),
	)

	return addr, lastErr
}
