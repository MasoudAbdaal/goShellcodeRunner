package dll

import "syscall"

// TODO: Refactor all function to use New_kernel32()
var (
	New_kernel32 kernel32DLL
	Kernel32     *syscall.LazyDLL = syscall.NewLazyDLL("kernel32.dll")
	Use32dll     *syscall.LazyDLL = syscall.NewLazyDLL("User32.dll")
	NtDll        *syscall.LazyDLL = syscall.NewLazyDLL("ntdll.dll")
)
