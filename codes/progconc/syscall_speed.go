package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strconv"
)

/*	syscall_speed.go
	referring to syscall_speed.c

	By repeatedly invoking a simple system call (getppid()), we can get some
	idea of the cost of making system calls.

	Usage: time syscall_speed functionType numCalls

	functionType: system or user
*/
func main() {
	// Use system call in default
	functionType := "system"
	// Repeat 10000000 times in default
	numCalls := 10000000

	functionType = os.Args[1]
	if len(os.Args) > 2 {
		if v, err := strconv.Atoi(os.Args[2]); err == nil {
			numCalls = v
		}
	}

	var f func() int
	switch functionType {
	case "user":
		f = myFunc
		fmt.Println("Calling normal function")
	case "system":
		f = unix.Getppid
		fmt.Println("Calling getppid()")
	default:
		os.Exit(-1)
	}

	for i := 0; i < numCalls; i++ {
		f()
	}
}

func myFunc() int {
	return 1
}
