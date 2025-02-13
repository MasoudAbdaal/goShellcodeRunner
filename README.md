# Shellcode Execution in Golang

This repository demonstrates a clean and modular approach to executing shellcode in a Windows environment using Golang. The project leverages Windows API calls via the `windows` and `syscall` packages to dynamically allocate memory, copy shellcode, change memory permissions, and create a new thread for execution.


## Features

- **Memory Allocation:**  
  Utilizes `VirtualAlloc` to reserve and commit memory for the shellcode.

- **Shellcode Copy:**  
  Uses `RtlMoveMemory` from `Ntdll.dll` to copy the decoded shellcode into the allocated memory space.

- **Memory Protection:**  
  Changes memory permissions from read-write to execute-read using `VirtualProtect`, ensuring the shellcode can be safely executed.

- **Thread Creation:**  
  Employs `CreateThread` from `kernel32.dll` to run the shellcode in a new thread.

- **Thread Synchronization:**  
  Implements `WaitForSingleObject` to pause the main program until the shellcode execution completes, preventing premature termination.



## Code Quality & Structure

- **Clean & Modular Code:**  
  The source code is organized into clear, self-contained functions (e.g., `AllocateShellcode`, `CopyShellcodeToMemory`, `ChangeShellcodeMemoryToRX`, `CreateThread`, and `RunShellcode`). This structure not only makes the code easy to read and maintain but also serves as a good reference for best practices in shellcode execution on Windows using Golang.

- **Descriptive Logging & Comments:**  
  Throughout the code, descriptive log messages and inline comments explain the purpose of each operation, enhancing readability and maintainability.


## How It Works

1. **Shellcode Decoding & Memory Allocation:**  
   The shellcode is stored as a hex string and is decoded into a byte slice using `hex.DecodeString`. Memory is then allocated using `VirtualAlloc` to provide a writable region.

2. **Copying & Setting Up Shellcode:**  
   The decoded shellcode is copied into the allocated memory using `RtlMoveMemory`. Memory permissions are changed from read-write to execute-read with `VirtualProtect` to allow execution.

3. **Execution via Thread Creation:**  
   A new thread is created with `CreateThread` that points directly to the shellcode's memory location. The main function waits for the thread to finish execution using `WaitForSingleObject`.


## Usage

1. **Setup:**  
   Ensure you have Go installed on your Windows machine and that your environment is properly configured.

2. **Build the Project:**  
   Compile the project by running:
   ```bash
   go build -o shellcodeRunner.exe
   ```

3. **Execute:**  
   Run the generated executable. The shellcode (provided as a hex string in the code) will be allocated, copied into memory, and executed within a new thread.

## Disclaimer

This project is intended for educational and research purposes only. Use it responsibly and ensure you have proper authorization before executing or testing shellcode in any environment.

---

Feel free to explore the code and use it as a learning tool to understand advanced Windows API interactions and dynamic code execution using Golang.
* [Main Source](https://www.scriptchildie.com/)