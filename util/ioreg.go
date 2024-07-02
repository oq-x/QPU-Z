package util

import (
	"fmt"
	"os/exec"

	"howett.net/plist"
)

type RegPCIDevice map[string]any

func (r RegPCIDevice) BuiltIn() bool {
	v, ok := r["built-in"].([]byte)
	return ok && len(v) != 0 && v[0] == 1
}

func (r RegPCIDevice) DataKey(key string) []byte {
	v, _ := r[key].([]byte)

	return v
}

func (r RegPCIDevice) StringKey(key string) string {
	v, _ := r[key].(string)

	return v
}

func (r RegPCIDevice) IORegistryEntryName() string {
	return r.StringKey("IORegistryEntryName")
}

func (r RegPCIDevice) Name() string {
	return r.StringKey("name")
}

func (r RegPCIDevice) IOName() string {
	return r.StringKey("IOName")
}

func (r RegPCIDevice) VendorID() string {
	v, ok := r["vendor-id"].([]byte)

	if !ok || len(v) != 4 {
		return ""
	}

	return PlistDataToLEHex(v)
}

func (r RegPCIDevice) DeviceID() string {
	v, ok := r["device-id"].([]byte)

	if !ok || len(v) != 4 {
		return ""
	}

	return PlistDataToLEHex(v)
}

func (r RegPCIDevice) SubsystemVendorID() string {
	v, ok := r["subsystem-vendor-id"].([]byte)

	if !ok || len(v) != 4 {
		return ""
	}

	return PlistDataToLEHex(v)
}

func (r RegPCIDevice) SubsystemID() string {
	v, ok := r["subsystem-id"].([]byte)

	if !ok || len(v) != 4 {
		return ""
	}

	return PlistDataToLEHex(v)
}

var IORegistry ioregistry

type ioregistry struct {
	PciDevices []RegPCIDevice

	GPUAccelerated bool
}

func Command(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func PlistDataToLEHex(data []byte) string {
	return fmt.Sprintf("%02x%02x", data[1], data[0])
}

func FetchIORegistry() ioregistry {
	var data = make(map[string]interface{})
	l, _ := Command("ioreg -a -l")
	plist.Unmarshal(l, data)

	_, nok := data["IONDRVFramebufferGeneration"]
	rootCh0, _ := data["IORegistryEntryChildren"].([]any)[0].(map[string]any)
	rootCh0ACPI, _ := rootCh0["IORegistryEntryChildren"].([]any)[0].(map[string]any)
	var rootCh0ACPIPCI0 map[string]any
	for _, ch := range rootCh0ACPI["IORegistryEntryChildren"].([]any) {
		m, ok := ch.(map[string]any)
		if !ok {
			continue
		}
		if m["IORegistryEntryName"] == "PCI0" {
			rootCh0ACPIPCI0 = m
			break
		}
	}
	rootCh0ACPIPCI0ACPI, _ := rootCh0ACPIPCI0["IORegistryEntryChildren"].([]any)[0].(map[string]any)
	rootCh0ACPIPCI0ACPIch := rootCh0ACPIPCI0ACPI["IORegistryEntryChildren"].([]any)
	var devices = make([]RegPCIDevice, len(rootCh0ACPIPCI0ACPIch))

	for i, ch := range rootCh0ACPIPCI0ACPIch {
		m, ok := ch.(map[string]any)
		if !ok {
			continue
		}

		devices[i] = m
	}

	return ioregistry{PciDevices: devices, GPUAccelerated: !nok}
}
