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

type pciClassCache map[string]DeviceClass

func (cache pciClassCache) Class(classId string) (cl DeviceClass, ok bool) {
	classId = strings.ToLower(classId)

	cl, ok = cache[classId]

	return
}

func (cache pciClassCache) Subclass(classId, subclassId string) (cl DeviceSubclass, ok bool) {
	classId, subclassId = strings.ToLower(classId), strings.ToLower(subclassId)

	cl, ok = cache[classId].DeviceSubclasses[subclassId]

	return
}

func (cache pciClassCache) ProgramInterface(classId, subclassId, programInterfaceId string) (inf ProgramInterface, ok bool) {
	classId, subclassId, programInterfaceId = strings.ToLower(classId), strings.ToLower(subclassId), strings.ToLower(programInterfaceId)

	inf, ok = cache[classId].DeviceSubclasses[subclassId].ProgramInterfaces[programInterfaceId]

	return
}

var PCIDeviceCache pciCache
var PCIClassCache pciClassCache

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

func pciidhandledef(currentVendor *Vendor, currentDevice *Device, dc pciCache, line string) (continu bool, err error) {
	switch {
	case line[0:1] == "\t" && line[1:2] == "\t": // new subsytem
		if currentDevice == nil {
			return false, fmt.Errorf("invalid pci data file")
		}
		line = line[2:]
		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}
		subsystemVID := line[:i][:4]
		subsystemID := line[:i][5:]

		currentDevice.Subsystems[[2]string{subsystemVID, subsystemID}] = Subsystem{
			VendorID: subsystemVID,
			ID:       subsystemID,
			Name:     line[i+2:],

			Parent: *currentDevice,
		}
	case line[0:1] == "\t": // new device
		if currentDevice != nil {
			currentVendor.Devices[currentDevice.ID] = *currentDevice
		}
		line = line[1:]
		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}

		*currentDevice = Device{
			ID:         line[:i],
			Name:       line[i+2:],
			Subsystems: make(map[[2]string]Subsystem),

			Parent: *currentVendor,
		}
	default: // new vendor
		if currentVendor.ID != "" {
			if currentDevice.ID != "" {
				currentVendor.Devices[currentDevice.ID] = *currentDevice
			}
			dc[currentVendor.ID] = *currentVendor
		}
		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}
		if line[:i] == "ffff" { //
			return false, nil
		}
		*currentVendor = Vendor{
			ID:      line[:i],
			Name:    line[i+2:],
			Devices: make(map[string]Device),

			Parent: PCIDeviceCache,
		}
	}
	return true, nil
}

func pciidhandlecl(currentClass *DeviceClass, currentSubclass *DeviceSubclass, cc pciClassCache, line string) (continu bool, err error) {
	switch {
	case line[0:1] == "\t" && line[1:2] == "\t":
		line = line[2:]

		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}

		currentSubclass.ProgramInterfaces[line[:i]] = ProgramInterface{
			ID:   line[:i],
			Name: line[i+2:],

			Parent: *currentSubclass,
		}
	case line[0:1] == "\t":
		if currentSubclass != nil {
			currentClass.DeviceSubclasses[currentSubclass.ID] = *currentSubclass
		}
		line = line[1:]

		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}

		*currentSubclass = DeviceSubclass{
			ID:                line[:i],
			Name:              line[i+1:],
			ProgramInterfaces: make(map[string]ProgramInterface),

			Parent: *currentClass,
		}
	case line[0] == 'C':
		if currentClass.ID != "" {
			if currentSubclass.ID != "" {
				currentClass.DeviceSubclasses[currentSubclass.ID] = *currentSubclass
			}
			cc[currentClass.ID] = *currentClass
		}
		line = line[2:]

		i := strings.Index(line, "  ")
		if i == -1 {
			return false, fmt.Errorf("invalid pci data file")
		}

		*currentClass = DeviceClass{
			ID:               line[:i],
			Name:             line[i+2:],
			DeviceSubclasses: make(map[string]DeviceSubclass),

			Parent: cc,
		}
	}
	return true, nil
}

func FetchPCIID() error {
	var (
		currentVendor Vendor
		currentDevice Device

		currentClass    DeviceClass
		currentSubclass DeviceSubclass
	)

	PCIDeviceCache = make(pciCache)
	PCIClassCache = make(pciClassCache)

	res, err := http.Get("https://raw.githubusercontent.com/pciutils/pciids/master/pci.ids")
	if err != nil {
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")

	var atClasses bool
	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue //comment
		}
		if !atClasses {
			c, err := pciidhandledef(&currentVendor, &currentDevice, PCIDeviceCache, line)
			if err != nil {
				return err
			}
			if !c {
				atClasses = true
			}
		} else {
			c, err := pciidhandlecl(&currentClass, &currentSubclass, PCIClassCache, line)
			if err != nil {
				return err
			}
			if !c {
				break
			}
		}
	}
	return nil
}

type DeviceClass struct {
	Name, ID         string
	DeviceSubclasses map[string]DeviceSubclass

	Parent pciClassCache
}

func (class DeviceClass) Subclass(subclassId string) (cl DeviceSubclass, ok bool) {
	subclassId = strings.ToLower(subclassId)

	cl, ok = class.DeviceSubclasses[subclassId]

	return
}

func (class DeviceClass) ProgramInterface(subclassId, programInterfaceId string) (inf ProgramInterface, ok bool) {
	subclassId, programInterfaceId = strings.ToLower(subclassId), strings.ToLower(programInterfaceId)

	inf, ok = class.DeviceSubclasses[subclassId].ProgramInterfaces[programInterfaceId]

	return
}

type DeviceSubclass struct {
	Name, ID          string
	ProgramInterfaces map[string]ProgramInterface

	Parent DeviceClass
}

func (class DeviceSubclass) ProgramInterface(programInterfaceId string) (inf ProgramInterface, ok bool) {
	programInterfaceId = strings.ToLower(programInterfaceId)

	inf, ok = class.ProgramInterfaces[programInterfaceId]

	return
}

type ProgramInterface struct {
	Name, ID string

	Parent DeviceSubclass
}
