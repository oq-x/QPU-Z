package specs

import (
	"qpu-z/util"
	"strings"
)

type AppleBoard struct {
	Manufacturer string
	Model        string
	BoardID      string
	SerialNumber string
}

type Board struct {
	Vendor string
	Model  string
}

func GetAppleBoard() AppleBoard {
	datastr, _ := util.Command("ioreg -l -p IODeviceTree -r -n / -d 1 | grep -iE 'board-id\"|manufacturer\"|model\"|IOPlatformSerialNumber' | awk '{$1=$1;print}' FS='[<>]' OFS=' '")
	data := make(map[string]string)
	for _, l := range strings.Split(string(datastr), "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		sp := strings.Split(l, " = ")
		key := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(sp[0]), "\""), "\"")
		value := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(sp[1]), "\""), "\"")
		data[key] = value
	}
	return AppleBoard{Manufacturer: data["manufacturer"], SerialNumber: data["IOPlatformSerialNumber"], Model: data["model"], BoardID: data["board-id"]}
}

func GetBoard() (Board, bool) {
	datastr, _ := util.Command("nvram 4D1FDA02-38C7-4A6A-9CC6-4BCCA8B30102:oem-vendor 4D1FDA02-38C7-4A6A-9CC6-4BCCA8B30102:oem-board")
	board := Board{}
	for _, l := range strings.Split(string(datastr), "\n") {
		if strings.HasPrefix(l, "nvram: Error") {
			return board, false
		}
		l = strings.TrimPrefix(l, "4D1FDA02-38C7-4A6A-9CC6-4BCCA8B30102:oem-")
		if strings.HasPrefix(l, "vendor") {
			board.Vendor = strings.TrimSpace(strings.TrimPrefix(l, "vendor"))
		} else if strings.HasPrefix(l, "board") {
			board.Model = strings.TrimSpace(strings.TrimPrefix(l, "board"))
		}
	}
	return board, true
}
