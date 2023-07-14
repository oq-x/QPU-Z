package specs

import (
	"qpu-z/util"
	"strconv"
	"strings"
)

type GPU struct {
	Model    string
	DeviceID string
	VendorID string
	Vendor   string
	VRAM     uint64
	VRAMText string
	Core     string
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

func GetVendor(data string) (name string, id string) {
	sp := strings.Split(data, "(")
	if len(sp) == 1 {
		return data, ""
	}
	name = strings.TrimSpace(sp[0])
	id = strings.TrimSuffix(sp[1], ")")
	return
}

func GetGPUs() []GPU {
	output := Command("system_profiler SPDisplaysDataType | awk '/Chipset Model:/{print \"\";print} /Vendor:|Device ID:|VRAM \\(Total\\):/'")
	gs := strings.Split(output, "\n\n")
	gpus := make([]GPU, 0)
	for _, gpu := range gs {
		spl := strings.Split(gpu, "\n")
		if spl[0] == "" {
			spl = spl[1:]
			gpu = strings.Join(spl, "\n")
		}
		if !strings.HasPrefix(gpu, "      ") {
			continue
		}
		g := make(map[string]string)
		for _, l := range spl {
			if l == "" {
				continue
			}
			l = strings.TrimSpace(l)
			sp := strings.Split(l, ":")
			key := strings.TrimSpace(sp[0])
			value := strings.TrimSpace(sp[1])
			g[key] = value
		}
		vname, vid := GetVendor(g["Vendor"])
		if vid == "" && vname == "Intel" {
			vid = "8086"
		}
		devid := strings.Split(g["Device ID"], "x")[1]
		vevid := strings.Split(vid, "x")[1]
		gpuExtended := util.GetDevice(vevid, devid)
		core := strings.Split(gpuExtended.Name, "[")[0]
		gpu := GPU{Model: g["Chipset Model"], DeviceID: g["Device ID"], Vendor: vname, VendorID: vid, VRAMText: g["VRAM (Total)"], Core: core}
		if g["VRAM (Total)"] != "" {
			gpu.VRAM = CalculateVRAM(g["VRAM (Total)"])
		}
		gpus = append(gpus, gpu)
	}
	return gpus
}
