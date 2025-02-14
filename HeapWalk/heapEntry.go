package heapwalk

import "unsafe"

type HeapBlock struct {
	wFlags      uint16
	iBlockIndex uint16
}

type HeapRegion struct {
	dwCommittedSize   uint32
	dwUnCommittedSize uint32
	lpFirstBlock      uintptr
	lpLastBlock       uintptr
}

// GoLang does not support union datatype and might this conversion DOES NOT currect
// Validate any data region before you rely on this manual mappings
//
// TODO: Review validation of fields value after mapping to unionData [24]byte
// Are they is the same as while called from C++?
// For verify data use other heap-related functions to find-out is the data valid?
type PROCESS_HEAP_ENTRY_UNION struct {
	lpData       uintptr
	cbData       uint32
	cbOverhead   byte
	iRegionIndex byte
	unionData    [24]byte
}

func (entry *PROCESS_HEAP_ENTRY_UNION) Block() *HeapBlock {
	return (*HeapBlock)(unsafe.Pointer(&entry.unionData[0]))
}

func (entry *PROCESS_HEAP_ENTRY_UNION) Region() *HeapRegion {
	return (*HeapRegion)(unsafe.Pointer(&entry.unionData[0]))
}

// https://learn.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-process_heap_entry#syntax
type PROCESS_HEAP_ENTRY struct {
	//lp: long pointer
	lpData uintptr
	//cb: count of bytes
	cbData       uint32
	cbOverhead   byte
	iRegionIndex byte
	wFlags       uint16
	iBlockIndex  uint16
	dwReserved   [3]uintptr
}
