package main

import (
	"errors"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const streamerExecutablePath = "C:\\Program Files\\Virtual Desktop Streamer\\VirtualDesktop.Streamer.exe"
const streamerExecutableName = "VirtualDesktop.Streamer.exe"

func isVDProcessRunning() bool {
	cmd := exec.Command("tasklist", "/FO", "CSV", "/FI", "IMAGENAME eq "+streamerExecutableName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), streamerExecutableName)
}

func killVDProcess() error {
	cmd := exec.Command("taskkill", "/F", "/IM", streamerExecutableName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	return err
}

func launchVDProcess() error {
	err := exec.Command(streamerExecutablePath).Start()
	if err != nil {
		return nil
	}
	return waitForVDProcess()
}

func waitForVDProcess() error {
	timeout := time.After(5 * time.Second)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return errors.New("process didn't appear")
		case <-ticker.C:
			if isVDProcessRunning() {
				return nil
			}
		}
	}
}
