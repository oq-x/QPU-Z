package specs

import (
	"fmt"
	"strings"
)

type StickRAM struct {
	ID           string
	Size         string
	Type         string
	Speed        string
	Manufacturer string
	SerialNumber string
}

func GetMemory() map[string]StickRAM {
	output := Command("system_profiler SPMemoryDataType | awk '/Size:/{print \"\";print} /Type:|Speed:|Serial Number:|Manufacturer/'")
	sticks := make(map[string]StickRAM)
	for i, stick := range strings.Split(output, "\n\n") {
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
		sticks[id] = StickRAM{
			ID:           id,
			Size:         data["Size"],
			Type:         data["Type"],
			Speed:        data["Speed"],
			Manufacturer: data["Manufacturer"],
			SerialNumber: data["Serial Number"],
		}
	}
	return sticks
}
