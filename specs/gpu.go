package specs

import (
	"fmt"
	"math"
	"qpu-z/util"
	"strconv"
	"strings"
)

type GPU struct {
	Model       string
	DeviceID    string
	VendorID    string
	SubsystemID string
	Vendor      string
	VRAM        string
	Core        string
	Internal    bool
}

func CalculateVRAM(data string) uint64 {
	num, _ := strconv.Atoi(strings.TrimSpace(strings.Split(data, " ")[0]))

	if strings.HasSuffix(data, "GB") {
		return uint64(num * 1073741824)
	} else if strings.HasSuffix(data, "MB") {
		return uint64(num * 1048576)
	}
	return uint64(num)
}

func GetVendor(id string) string {
	switch id {
	case "8086":
		{
			return "Intel"
		}
	case "1002":
		{
			return "AMD"
		}
	case "10de":
		{
			return "Nvidia"
		}
	}
	return "Unknown"
}

func parseData(data []byte) string {
	str := dataString(data)
	sp := strings.Split(str, "")
	sp = sp[:len(sp)-4]
	return fmt.Sprintf("%s%s%s%s", sp[2], sp[3], sp[0], sp[1])
}

func dataString(data []byte) string {
	var builder strings.Builder
	for _, b := range data {
		builder.WriteString(fmt.Sprintf("%02X", b))
	}
	return builder.String()
}

func GetGPUs() []GPU {
	gpusd := make(map[string]map[string]interface{})
	for _, v := range util.IORegistry.PciDevices {
		if v["IOName"] == "display" {
			gpusd[v.IORegistryEntryName()] = v
		}
	}
	gpus := make([]GPU, 0)
	for _, g := range gpusd {
		vendorId := parseData(g["vendor-id"].([]byte))
		deviceId := parseData(g["device-id"].([]byte))
		subsystemId := parseData(g["subsystem-vendor-id"].([]byte)) + parseData(g["subsystem-id"].([]byte))
		gpuExtended, ok := util.PCIDeviceCache.Device(vendorId, deviceId)
		if !ok {
			continue
		}
		core := strings.Split(gpuExtended.Name, "[")[0]
		vram := ""
		if g["VRAM,totalMB"] != nil {
			v := g["VRAM,totalMB"].(uint64)
			if v < 1024 {
				vram = fmt.Sprintf("%dMB", v)
			} else {
				v := float64(v) / 1024
				vram = fmt.Sprintf("%dGB", int(math.Round(v)))
			}
		}
		gpu := GPU{Model: string(g["model"].([]byte)), DeviceID: deviceId, VendorID: vendorId, Core: core, SubsystemID: subsystemId, Vendor: GetVendor(vendorId), VRAM: vram}
		if g["IORegistryEntryName"] == "IGPU" {
			gpu.Internal = true
		}
		gpus = append(gpus, gpu)
	}
	return gpus
}
