package specs

import (
	"qpu-z/util"
	"strconv"
	"strings"
)

type CPU struct {
	Model        string
	Cores        int
	Threads      int
	Instructions []string
	Vendor       string
	Count        int

	GenerationDisplayName string
}

var intelFamilyModelMap = map[int]map[int]string{
	15: {
		1: "Netburst (Williamette)",
		2: "Netburst (Northwood)",
		3: "Netburst (Prescott)",
		4: "Netburst (Prescott)",
		6: "Netburst",
	},
	11: {
		0: "Knights Ferry",
		1: "Knights Corner",
	},
	6: {
		183: "Raptor Lake-S",
		186: "Raptor Lake-P",

		151: "Alder Lake-S",
		154: "Alder Lake-P",

		167: "Rocket Lake-S",

		141: "Tiger Lake-H",
		140: "Tiger Lake-U",

		126: "Ice Lake-U/Y",

		165: "Comet Lake-S/H",
		142: "Comet Lake/Coffee Lake/Whiskey Lake/Kaby Lake-U/Amber Lake-Y",

		102: "Cannon Lake-U",

		158: "Coffee Lake-S/H/E",

		94: "Skylake-DT/H/S",
		78: "Skylake-Y/U",

		71: "Broadwell-C/W/H",
		61: "Broadwell-U/Y/S",

		70: "Haswell-GT3E",
		69: "Haswell-ULT",
		60: "Haswell-S",

		58: "Ivy Bridge",

		42: "Sandy Bridge",

		37: "Westmere (Arrandale/Clarkdale)",
		31: "Nahelem (Auburndale/Havendale)",
		30: "Nahelem (Clarksfield)",
		23: "Penyrn/Wolfdale/Yorkfield",
		22: "Merom-L",
		15: "Merom",
		14: "Yonah",
		21: "Tolapai",
		13: "Dothan",
		9:  "Banias",
		11: "Tualatin",
		8:  "Coppermine",
		7:  "Katmai",
	},
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
	vendor := data["machdep.cpu.vendor"].(string)
	if vendor == "GenuineIntel" {
		vendor = "Intel"
	} else if vendor == "AuthenticAMD" {
		vendor = "AMD"
	}

	family := data["machdep.cpu.family"].(int)
	model := data["machdep.cpu.model"].(int)

	return CPU{
		Cores:                 data["machdep.cpu.core_count"].(int),
		Threads:               data["machdep.cpu.thread_count"].(int),
		Model:                 data["machdep.cpu.brand_string"].(string),
		Vendor:                vendor,
		Count:                 count,
		GenerationDisplayName: intelFamilyModelMap[family][model],
	}
}
