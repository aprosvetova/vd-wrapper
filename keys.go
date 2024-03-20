package main

import (
	"errors"

	"git.tcp.direct/kayos/sendkeys"
)

func enterUsernames(oculus, pico, vive []string) error {
	if len(oculus) > 4 || len(pico) > 4 || len(vive) > 4 {
		return errors.New("only 4 usernames per platform supported")
	}
	k, err := sendkeys.NewKBWrapWithOptions()
	if err != nil {
		return err
	}

	// Getting to the first Oculus username field
	k.Tab()
	k.Tab()
	k.Tab()
	k.Tab()

	oculus = padSliceTo4(oculus)
	pico = padSliceTo4(pico)
	vive = padSliceTo4(vive)

	usernames := append(append(oculus, pico...), vive...)

	// Enter all usernames, skip unused fields
	for _, username := range usernames {
		k.Type(username)
		k.Tab()
	}

	// Hit Save
	k.Enter()

	return nil
}

func padSliceTo4(slice []string) []string {
	for len(slice) < 4 {
		slice = append(slice, "")
	}
	return slice
}
