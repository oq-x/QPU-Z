package util

import (
	"fmt"
	"os/exec"

	"howett.net/plist"
)

var IORegistry ioregistry

type ioregistry struct {
	Devices map[string]map[string]interface{}

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

	var devices = make(map[string]map[string]any)

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
	for _, ch := range rootCh0ACPIPCI0ACPI["IORegistryEntryChildren"].([]any) {
		m, ok := ch.(map[string]any)
		if !ok {
			continue
		}
		devices[m["IORegistryEntryName"].(string)] = m
	}

	return ioregistry{Devices: devices, GPUAccelerated: !nok}
}
