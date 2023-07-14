package util

import (
	"io"
	"net/http"
	"strings"
)

var deviceCache map[string]Device

type Device struct {
	Name    string
	ID      string
	Devices map[string]Device
}

func ParsePCI() map[string]Device {
	if deviceCache == nil {
		vendors := make(map[string]Device)
		res, err := http.Get("https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids")
		if err != nil {
			return vendors
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return vendors
		}
		raw := string(body)
		nc := ""
		for _, l := range strings.Split(raw, "\n") {
			if strings.HasPrefix(l, "#") {
				continue
			}
			nc += l + "\n"
		}
		sp := strings.Split(nc, "\n")
		var currentVendor Device
		var currentDevice Device
		for i, l := range sp {
			if !strings.HasPrefix(l, "\t") {
				if len(sp) > i+1 && !strings.HasPrefix(sp[i+1], "\t") {
					continue
				}
				ext := strings.Split(strings.TrimSpace(l), "  ")
				if len(ext) < 2 {
					continue
				}
				currentVendor = Device{
					ID:      ext[0],
					Name:    ext[1],
					Devices: make(map[string]Device),
				}
				vendors[ext[0]] = currentVendor
			} else {
				if !strings.HasPrefix(l, "\t\t") {
					ext := strings.Split(strings.TrimSpace(l), "  ")
					if len(ext) < 2 {
						continue
					}
					currentVendor.Devices[ext[0]] = Device{
						ID:      ext[0],
						Name:    ext[1],
						Devices: make(map[string]Device),
					}
					if len(sp) > i+1 && strings.HasPrefix(sp[i+1], "\t\t") {
						currentDevice = currentVendor.Devices[ext[0]]
					}
				} else {
					ext := strings.Split(strings.TrimSpace(l), "  ")
					if len(ext) < 2 {
						continue
					}
					currentDevice.Devices[ext[0]] = Device{
						ID:   ext[0],
						Name: ext[1],
					}
				}
			}
		}
		deviceCache = vendors
	}
	return deviceCache
}

func GetDevice(v string, d string) Device {
	vendors := ParsePCI()
	vendor := vendors[strings.ToLower(v)]
	if vendor.ID == "" {
		return Device{}
	}
	device := vendor.Devices[strings.ToLower(d)]
	if device.ID == "" {
		return Device{}
	}
	return device
}
