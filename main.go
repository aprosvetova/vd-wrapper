package main

import (
	"fmt"
)

func main() {
	err := parseArgs()
	if err != nil {
		alertError(err.Error())
		return
	}

	if isVDProcessRunning() {
		err := killVDProcess()
		if err != nil {
			alertError(fmt.Sprintf("VD was already running, and I failed to kill it:\n%s", err))
			return
		}
	}

	err = clearIdentifiers()
	if err != nil {
		alertError(fmt.Sprintf("Couldn't clear the identifiers in the VD config:\n%s\n\nVD won't launch", err))
		return
	}

	err = launchVDProcess()
	if err != nil {
		alertError(fmt.Sprintf("Couldn't launch VD:\n%s", err))
		return
	}

	err = waitForVDWindow()
	if err != nil {
		alertError("VD window never showed up")
		return
	}

	err = bringFocusToVDWindow()
	if err != nil {
		alertError(fmt.Sprintf("Couldn't focus on the VD window:\n%s", err))
		return
	}

	err = enterUsernames(args.Oculus, args.Pico, args.Vive)
	if err != nil {
		alertError(fmt.Sprintf("Couldn't enter usernames:\n%s", err))
		return
	}

	err = closeVDWindow()
	if err != nil {
		alertError(fmt.Sprintf("Couldn't close the VD window:\n%s", err))
		return
	}
}
