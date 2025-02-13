package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func ShowMessageBoxes() {
	//  In Win32, a window object is identified by a value known as a window handle.
	//  And the type of a window handle is an HWND
	// https://learn.microsoft.com/en-us/windows/apps/develop/ui-input/retrieve-hwnd
	hWnd := uintptr(0)
	windows.MessageBox(
		windows.HWND(hWnd),
		windows.StringToUTF16Ptr("Windows package!!"),
		windows.StringToUTF16Ptr("MessageBoxx"),
		windows.MB_OKCANCEL)

	// https://pkg.go.dev/syscall?GOOS=windows#LazyDLL
	// A LazyDLL implements access to a single DLL.
	// It will delay the load of the DLL until the first call to its LazyDLL.Handle method or to one of its LazyProc's Addr method.
	use32dll := syscall.NewLazyDLL("User32.dll")
	procMsgBox := use32dll.NewProc("MessageBoxW")

	lpText, _ := syscall.UTF16PtrFromString("Used Syscall Package")
	lpCaption, _ := syscall.UTF16PtrFromString("DLL Imported Message")
	uType := uint(2)

	procMsgBox.Call(hWnd,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType))
}
