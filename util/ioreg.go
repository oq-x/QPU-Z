package util

import (
	"os/exec"

	"howett.net/plist"
)

var IORegistry ioregistry

type ioregistry struct {
	Devices map[string]map[string]interface{}
}

func Command(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func loopChildren(class map[string]interface{}) map[string]interface{} {
	children, ok := class["IORegistryEntryChildren"].([]interface{})
	if (class["IORegistryEntryChildren"] == nil || class["model"] != nil) || !ok {
		return class
	}
	for _, v := range children {
		v, ok := v.(map[string]interface{})
		if ok {
			if v["IORegistryEntryChildren"] == nil || class["model"] != nil {
				return v
			} else {
				return loopChildren(v)
			}
		}
	}
	return map[string]interface{}{}
}

func FetchIORegistry() ioregistry {
	var data = make(map[string]interface{})
	l, _ := Command("ioreg -a -k IOName")
	plist.Unmarshal(l, data)
	ch := data["IORegistryEntryChildren"].([]interface{})[0].(map[string]interface{})["IORegistryEntryChildren"].([]interface{})[0].(map[string]interface{})["IORegistryEntryChildren"].([]interface{})
	devices := make(map[string]map[string]interface{})
	for _, v := range ch {
		v, ok := v.(map[string]interface{})
		if ok {
			if v["IORegistryEntryName"] == "PCI0" {
				ch := v["IORegistryEntryChildren"].([]interface{})[0].(map[string]interface{})["IORegistryEntryChildren"].([]interface{})
				for _, v := range ch {
					v, ok := v.(map[string]interface{})
					if !ok || v["IORegistryEntryChildren"] == nil {
						continue
					}

					device := loopChildren(v)
					if device["model"] == nil {
						continue
					}
					devices[device["IORegistryEntryName"].(string)] = device
				}
			}
		}
	}
	return ioregistry{Devices: devices}
}
