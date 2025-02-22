package main

import (
	heapwalk "goShellcodeRunner/HeapWalk"
	shellcode "goShellcodeRunner/Shellcode"
)

func main() {
	heapwalk.GetProcessHeap()

	// shellcode.ExecuteCalculator()

	shellcode.ClassicInjection()
}
