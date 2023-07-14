package pages

import (
	"fmt"
	"qpu-z/specs"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MemoryPage() fyne.CanvasObject {
	sticks := specs.GetMemory()
	ids := []string{}
	for i := 0; i < len(sticks); i++ {
		ids = append(ids, fmt.Sprint(i))
	}
	tree := widget.NewTree(func(tni widget.TreeNodeID) []widget.TreeNodeID {
		if tni == "" {
			return ids
		}
		return []string{tni + "\\0"}
	}, func(tni widget.TreeNodeID) bool {
		if tni == "" {
			return true
		}
		num, e := strconv.Atoi(tni)
		if e != nil {
			return false
		}
		if ids[num] != "" {
			return sticks[tni].Type != "Empty"
		}
		return false
	}, func(b bool) fyne.CanvasObject {
		if b {
			return widget.NewRichTextFromMarkdown("")
		}
		return container.NewVBox(layout.NewSpacer(), widget.NewRichTextFromMarkdown("## TEXT\n## TEXT\n## TEXT\n## TEXT"), layout.NewSpacer())
	}, func(tni widget.TreeNodeID, b bool, co fyne.CanvasObject) {
		if b {
			co.(*widget.RichText).ParseMarkdown(fmt.Sprintf("### %s", tni))
			return
		}
		stick := sticks[strings.Split(tni, "\\")[0]]
		text := fmt.Sprintf("## Size: %s\n", stick.Size)
		text += fmt.Sprintf("## Type: %s\n", stick.Type)
		text += fmt.Sprintf("## Speed: %s\n", stick.Speed)
		if stick.SerialNumber != "" {
			text += fmt.Sprintf("## Serial Number: %s\n", stick.SerialNumber)
		}
		co.(*fyne.Container).Objects[1] = widget.NewRichTextFromMarkdown(text)
	})
	return tree
}
