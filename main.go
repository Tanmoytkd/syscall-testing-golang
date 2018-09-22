package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		panic("No Command Given")
	}

	command, lookErr := exec.LookPath(args[1])
	if lookErr != nil {
		panic(lookErr)
	}

	args = args[1:]

	env := os.Environ()

	commandBytePtr, err := syscall.BytePtrFromString(command)
	if err != nil {
		panic(err)
	}
	argvSlicePtr, err := syscall.SlicePtrFromStrings(args)
	if err != nil {
		panic(err)
	}
	envBytePtr, err := syscall.SlicePtrFromStrings(env)
	if err != nil {
		panic(err)
	}

	_, _, err = syscall.Syscall(syscall.SYS_EXECVE,
		uintptr(unsafe.Pointer(commandBytePtr)),
		uintptr(unsafe.Pointer(&argvSlicePtr[0])),
		uintptr(unsafe.Pointer(&envBytePtr[0])))

	if error(err) != nil {
		panic(err)
	}

	fmt.Println("The process quits and " +
		"this message does not get shown on screen")
}
