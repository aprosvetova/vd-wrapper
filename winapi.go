package main

import (
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procFindWindowW         = user32.NewProc("FindWindowW")
	procSetForegroundWindow = user32.NewProc("SetForegroundWindow")
	procPostMessageW        = user32.NewProc("PostMessageW")
	messageBoxW             = user32.NewProc("MessageBoxW")
)

func messageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := messageBoxW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}

func alertError(caption string) int {
	const (
		NULL         = 0
		MB_OK        = 0
		MB_ICONERROR = 0x00000010
	)
	return messageBox(NULL, caption, "VD Wrapper "+Version, MB_OK|MB_ICONERROR)
}

func bringFocus(windowTitle string) error {
	lpszWindow, err := syscall.UTF16PtrFromString(windowTitle)
	if err != nil {
		return err
	}
	hwnd, _, _ := procFindWindowW.Call(uintptr(0), uintptr(unsafe.Pointer(lpszWindow)))
	if hwnd == 0 {
		return syscall.GetLastError()
	}
	_, _, _ = procSetForegroundWindow.Call(hwnd)
	return nil
}

func closeWindow(windowTitle string) error {
	lpszWindow, err := syscall.UTF16PtrFromString(windowTitle)
	if err != nil {
		return err
	}

	hwnd, _, _ := procFindWindowW.Call(uintptr(0), uintptr(unsafe.Pointer(lpszWindow)))
	if hwnd == 0 {
		return syscall.GetLastError()
	}

	_, _, errPostMessage := procPostMessageW.Call(hwnd, uintptr(0x0112), uintptr(0xF060), 0) // WM_SYSCOMMAND = 0x0112, SC_CLOSE = 0xF060
	if errPostMessage != nil && errPostMessage.(syscall.Errno) != 0 {
		return errPostMessage
	}

	return nil
}
