package specs

import (
	"fmt"
	"qpu-z/util"
	"strings"
)

type MemoryModule struct {
	ID           string
	Size         string
	Type         string
	Speed        string
	Manufacturer string
	SerialNumber string
}

func GetMemory() map[string]MemoryModule {
	output, _ := util.Command("system_profiler SPMemoryDataType | awk '/Size:/{print \"\";print} /Type:|Speed:|Serial Number:|Manufacturer/'")
	modules := make(map[string]MemoryModule)
	for i, stick := range strings.Split(string(output), "\n\n") {
		id := fmt.Sprint(i)
		data := make(map[string]string)
		for _, l := range strings.Split(stick, "\n") {
			if l == "" {
				continue
			}
			l = strings.TrimSpace(l)
			sp := strings.Split(l, ":")
			val := strings.TrimSpace(sp[1])
			if val == "-" {
				continue
			}
			data[sp[0]] = val
		}
		modules[id] = MemoryModule{
			ID:           id,
			Size:         data["Size"],
			Type:         data["Type"],
			Speed:        data["Speed"],
			Manufacturer: data["Manufacturer"],
			SerialNumber: data["Serial Number"],
		}
	}
	return modules
}
