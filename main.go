package main

import (
	heapwalk "first_malware/HeapWalk"
)

func main() {
	heapwalk.GetProcessHeap()

	ExecuteCalculator()
}
