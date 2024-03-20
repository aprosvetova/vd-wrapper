package main

import (
	"errors"
	"os"
	"strings"

	"github.com/valyala/fastjson"
)

const configPath = "C:\\ProgramData\\Virtual Desktop\\StreamerSettings.json"

func clearIdentifiers() error {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return errors.New("can't read config json")
	}
	var p fastjson.Parser
	j, err := p.ParseBytes(b)
	if err != nil {
		return errors.New("config json is not json")
	}
	o, err := j.Object()
	if err != nil {
		return errors.New("unexpected config json format")
	}
	var toDelete []string
	o.Visit(func(key []byte, v *fastjson.Value) {
		ks := string(key)
		if strings.HasPrefix(ks, "Protected") || strings.HasPrefix(ks, "SavedPico") {
			toDelete = append(toDelete, ks)
		}
	})
	for _, key := range toDelete {
		o.Del(key)
	}
	err = os.WriteFile(configPath, o.MarshalTo(nil), os.ModePerm)
	if err != nil {
		return errors.New("can't save config json")
	}
	return nil
}
