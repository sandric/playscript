package main

import (
	"fmt"
	"time"
	"os"
	"syscall"
	"unsafe"
	"errors")

const TTY = "/dev/tty1"

func openTTY() int {
	fd, _ := syscall.Open(TTY, os.O_RDWR, 0666)
	return fd
}
func ioctl(fd int, request, argp uintptr) error {
	_, _, errorp := syscall.Syscall(
		syscall.SYS_IOCTL, 
		uintptr(fd), 
		request, 
		argp)

	return os.NewSyscallError("ioctl", errorp)
}

func printText(text string) error {
	fd := openTTY()
	var data = [] byte(text)

	for _, byte := range data {
		if  err := ioctl(fd, 
					 	 syscall.TIOCSTI, 
						 uintptr(unsafe.Pointer(&byte)));
			err != err {
				fmt.Println(err)
				return err
			} else {
				time.Sleep(200 * time.Millisecond)
			}
	}

	return errors.New("all done")
}

func main() {

	printText("ls -la\n")
}