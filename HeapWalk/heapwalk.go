package heapwalk

import (
	dll "goShellcodeRunner/DLL"
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetProcessHeap() {
	procGetProcessHeap := dll.Kernel32.NewProc("GetProcessHeap")
	procGetProcessHeapS := dll.Kernel32.NewProc("GetProcessHeaps")
	procHeapWalk := dll.Kernel32.NewProc("HeapWalk")

	numberOfHeapS, _, lastErr := procGetProcessHeapS.Call()

	if lastErr == windows.NTE_OP_OK {
		log.Printf("[HEAP] Numbers Of Heap Region For This Process (%v)", numberOfHeapS)
	}

	heapHandle, _, lastErr := procGetProcessHeap.Call()
	if lastErr == windows.NTE_OP_OK {

	}

	var heapEntry PROCESS_HEAP_ENTRY

	//TODO: delete unsafe pointer to prevent memory leaks
	ret, _, _ := procHeapWalk.Call(heapHandle, uintptr(unsafe.Pointer(&heapEntry)))

	if ret != 0 {
		log.Printf("[HEAP] Process Heap Size (%v) Bytes", heapEntry.cbData)
		log.Printf("[HEAP] Process Heap Address 0x[%x]", heapEntry.lpData)

	}

}
