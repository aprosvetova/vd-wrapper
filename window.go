package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const streamerWindowTitle = "Virtual Desktop Streamer"

func isVDWindowReady() bool {
	cmd := exec.Command("powershell", "-Command", "Get-Process | Where-Object { $_.MainWindowTitle -eq '"+streamerWindowTitle+"' } | Select-Object -ExpandProperty MainWindowTitle")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}
	return strings.Contains(out.String(), streamerWindowTitle)
}

func waitForVDWindow() error {
	timeout := time.After(10 * time.Second)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return errors.New("window didn't appear")
		case <-ticker.C:
			if isVDWindowReady() {
				return nil
			}
		}
	}
}

func bringFocusToVDWindow() error {
	return bringFocus(streamerWindowTitle)
}

func closeVDWindow() error {
	return closeWindow(streamerWindowTitle)
}
