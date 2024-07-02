package specs

import (
	"qpu-z/util"
	"strconv"
	"strings"
)

type CPU struct {
	Model   string
	Cores   int
	Threads int
	Vendor  string
	Count   int

	GenerationDisplayName             string
	GenerationFamily, GenerationModel int
}

func GetCPU() CPU {
	cmd, _ := util.Command("sysctl machdep.cpu")
	c, _ := util.Command("sysctl hw.packages")
	count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(string(c), ":")[1]))

	data := make(map[string]any)
	for _, l := range strings.Split(string(cmd), "\n") {
		if l == "" {
			continue
		}
		sp := strings.Split(l, ": ")
		i, e := strconv.Atoi(sp[1])
		if e == nil {
			data[sp[0]] = i
		} else {
			data[sp[0]] = sp[1]
		}
	}

	family := data["machdep.cpu.family"].(int)
	model := data["machdep.cpu.model"].(int)
	extmodel := data["machdep.cpu.extmodel"].(int)
	var genDisplayName = "Unknown"

	vendor := data["machdep.cpu.vendor"].(string)
	if vendor == "GenuineIntel" {
		vendor = "Intel"
		genDisplayName = intelFamilyModelMap[family][model]
	} else if vendor == "AuthenticAMD" {
		vendor = "AMD"
		genDisplayName = amdFamilyModelMap[family][[2]int{model, extmodel}]
	}

	return CPU{
		Cores:                 data["machdep.cpu.core_count"].(int),
		Threads:               data["machdep.cpu.thread_count"].(int),
		Model:                 data["machdep.cpu.brand_string"].(string),
		Vendor:                vendor,
		Count:                 count,
		GenerationDisplayName: genDisplayName,
		GenerationFamily:      family,
		GenerationModel:       model,
	}
}
