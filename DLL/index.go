package dll

import "syscall"

var (
	Kernel32 *syscall.LazyDLL = syscall.NewLazyDLL("kernel32.dll")

	NtDll *syscall.LazyDLL = syscall.NewLazyDLL("ntdll.dll")
)
