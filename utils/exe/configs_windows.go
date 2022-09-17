//go:build windows

package exe

import (
	"golang.org/x/sys/windows"
)

var (
	prefixCmd = []string{"powershell", "-nologo", "-noprofile"}
)

func RunMeElevated() error {
	verb := "runas"
    exe, _ := os.Executable()
    cwd, _ := os.Getwd()
    args := strings.Join(os.Args[1:], " ")

    verbPtr, _ := syscall.UTF16PtrFromString(verb)
    exePtr, _ := syscall.UTF16PtrFromString(exe)
    cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
    argPtr, _ := syscall.UTF16PtrFromString(args)

    var showCmd int32 = 1 //SW_NORMAL

    return windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd
}

func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")

    if err != nil 
        return false
    }
	
    return true
}