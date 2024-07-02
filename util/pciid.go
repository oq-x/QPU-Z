package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type pciCache map[string]Vendor

func (cache pciCache) Vendor(id string) (v Vendor, ok bool) {
	id = strings.ToLower(id)

	v, ok = cache[id]
	return
}

func (cache pciCache) Device(vendorId, deviceId string) (d Device, ok bool) {
	vendorId, deviceId = strings.ToLower(vendorId), strings.ToLower(deviceId)

	d, ok = cache[vendorId].Devices[deviceId]
	return
}

func (cache pciCache) Subsystem(vendorId, deviceId, subsystemVendorId, subsystemId string) (s Subsystem, ok bool) {
	vendorId, deviceId, subsystemVendorId, subsystemId = strings.ToLower(vendorId), strings.ToLower(deviceId), strings.ToLower(subsystemVendorId), strings.ToLower(subsystemId)

	s, ok = cache[vendorId].Devices[deviceId].Subsystems[[2]string{subsystemVendorId, subsystemId}]

	return
}

var PCIDeviceCache pciCache

type Subsystem struct {
	Name         string
	VendorID, ID string
	Parent       Device
}

type Device struct {
	Name       string
	ID         string
	Subsystems map[[2]string]Subsystem
	Parent     Vendor
}

func (d Device) Subsystem(subsystemVendorId, subsystemId string) (s Subsystem, ok bool) {
	subsystemVendorId, subsystemId = strings.ToLower(subsystemVendorId), strings.ToLower(subsystemId)

	s, ok = d.Subsystems[[2]string{subsystemVendorId, subsystemId}]

	return
}

type Vendor struct {
	Name    string
	ID      string
	Devices map[string]Device
	Parent  pciCache
}

func (v Vendor) Device(id string) (d Device, ok bool) {
	id = strings.ToLower(id)

	d, ok = v.Devices[id]

	return
}

func (v Vendor) Subsystem(deviceId, subsystemVendorId, subsystemId string) (s Subsystem, ok bool) {
	deviceId, subsystemVendorId, subsystemId = strings.ToLower(deviceId), strings.ToLower(subsystemVendorId), strings.ToLower(subsystemId)

	s, ok = v.Devices[deviceId].Subsystems[[2]string{subsystemVendorId, subsystemId}]

	return
}

func FetchPCIID() error {
	var (
		currentVendor *Vendor
		currentDevice *Device
	)

	PCIDeviceCache = make(pciCache)

	res, err := http.Get("https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")

lloop:
	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue //comment
		}
		switch {
		case strings.HasPrefix(line, "		"): // new subsytem
			if currentDevice == nil {
				return fmt.Errorf("invalid pci data file")
			}
			line = line[2:]
			i := strings.Index(line, "  ")
			if i == -1 {
				return fmt.Errorf("invalid pci data file")
			}
			subsystemVID := line[:i][:4]
			subsystemID := line[:i][5:]

			currentDevice.Subsystems[[2]string{subsystemVID, subsystemID}] = Subsystem{
				VendorID: subsystemVID,
				ID:       subsystemID,
				Name:     line[i:],

				Parent: *currentDevice,
			}
		case strings.HasPrefix(line, "	"): // new device
			if currentDevice != nil {
				currentVendor.Devices[currentDevice.ID] = *currentDevice
			}
			line = line[1:]
			i := strings.Index(line, "  ")
			if i == -1 {
				return fmt.Errorf("invalid pci data file")
			}

			currentDevice = &Device{
				ID:         line[:i],
				Name:       line[i:],
				Subsystems: make(map[[2]string]Subsystem),

				Parent: *currentVendor,
			}
		default: // new vendor
			if currentVendor != nil {
				PCIDeviceCache[currentVendor.ID] = *currentVendor
			}
			i := strings.Index(line, "  ")
			if i == -1 {
				return fmt.Errorf("invalid pci data file")
			}
			if line[:i] == "ffff" { //illegal vendor id
				break lloop
			}
			currentVendor = &Vendor{
				ID:      line[:i],
				Name:    line[i:],
				Devices: make(map[string]Device),

				Parent: PCIDeviceCache,
			}
		}
	}
	return nil
}
