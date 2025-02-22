package process

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// https://learn.microsoft.com/en-us/windows/win32/procthread/process-security-and-access-rights
const PROCESS_ALL_ACCESS = windows.STANDARD_RIGHTS_REQUIRED | windows.SYNCHRONIZE | 0xFFFF

func CreateCmdProcess() syscall.Handle {
	// TODO: In Bash, when you add a SPACE before your command, it hides from history!
	// How can I do the same in cmd or Powershell?
	// What other techniques is there for hide ourself?
	// https://www.linkedin.com/posts/activity-7290113056188112898-RBXE?utm_source=share&utm_medium=member_desktop&rcm=ACoAADR0WIUBHui16sjo_P6svUUR2yJg4DCoe7w
	cmdPtr, err := syscall.UTF16PtrFromString("C:\\Windows\\System32\\cmd.exe")
	cmdArgPtr, err2 := syscall.UTF16FromString(" /c ping /t c2.com")

	var startupInfo syscall.StartupInfo
	var procInfo syscall.ProcessInformation

	startupInfo.Cb = uint32(unsafe.Sizeof(startupInfo))

	if err != nil || err2 != nil {
		log.Panicf("[PROC] Error Happened In  UTF16PtrFromString Function (%v) (%v) \n", err, err2)
	}
	log.Printf("[PROC] Strings Has Been Initialized Successfully \n")

	err = syscall.CreateProcess(cmdPtr,
		&cmdArgPtr[0],
		nil, nil, false,
		// CREATE_NEW_CONSOLE = 0x00000010
		// CREATE_BREAKAWAY_FROM_JOB = 0x01000000
		// CREATE_DEFAULT_ERROR_MODE = 0x04000000
		// CREATE_NO_WINDOW = 0x08000000
		// CREATE_SEPARATE_WOW_VDM = 0x00000800
		// CREATE_UNICODE_ENVIRONMENT = 0x00000400
		// DEBUG_PROCESS = 0x00000001
		// https://learn.microsoft.com/en-us/windows/win32/procthread/process-creation-flags#flags

		// Used NewConsole for debuggging and end process!
		0x00000010, nil, nil,
		&startupInfo,
		&procInfo)

	if err != nil {
		log.Panicf("[PROC] Error In CreateProcess Function \"%v\" \n", err)
	}

	log.Printf("[PROC] cmd.exe Created Successfully With PID: (%v) \n", procInfo.ProcessId)

	log.Printf("[PROC] Try To Make A Handle From PID(%v) (cmd.exe) \n", procInfo.ProcessId)

	// createHandle_Old contents....

	log.Printf("[PROC] Handle From PID(%v) = %v \n", procInfo.ProcessId, procInfo.Process)

	return procInfo.Process
}

// func createHandle_Old() {
// 	pHandle, err := syscall.OpenProcess(PROCESS_ALL_ACCESS, false, procInfo.ProcessId)

// 	if err != nil {

// 		log.Panicf("[PROC] Error While Getting Handle \"%v\" \n", err)
// 	}

// 	log.Printf("[PROC] Handle From PID(%v) = %v \n", procInfo.ProcessId, pHandle)

// 	return pHandle
// }
