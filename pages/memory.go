package pages

import (
	"fmt"
	"qpu-z/specs"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MemoryPage() fyne.CanvasObject {
	modules := specs.GetMemory()
	var ids = make([]string, len(modules))
	for i := range modules {
		ids[i] = fmt.Sprint(i)
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
			return modules[num].Type != "Empty"
		}
		return false
	}, func(b bool) fyne.CanvasObject {
		if b {
			return widget.NewRichTextFromMarkdown("")
		}
		return widget.NewRichText(
			&widget.TextSegment{
				Text:  "Size: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Type: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Speed: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{},

			&widget.SeparatorSegment{},

			&widget.TextSegment{Text: "Serial Number", Style: widget.RichTextStyleStrong},
			&widget.TextSegment{},
		)
	}, func(tni widget.TreeNodeID, b bool, co fyne.CanvasObject) {
		if b {
			co.(*widget.RichText).ParseMarkdown(fmt.Sprintf("### %s", tni))
			return
		}
		d, _ := strconv.Atoi(strings.Split(tni, "\\")[0])

		module := modules[d]
		co.(*widget.RichText).Segments = []widget.RichTextSegment{
			&widget.TextSegment{
				Text:  "Size: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: module.Size},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Type: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: module.Type},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Speed: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: module.Speed},
		}

		if module.SerialNumber != "" {
			co.(*widget.RichText).Segments = append(co.(*widget.RichText).Segments,
				&widget.SeparatorSegment{},
				&widget.TextSegment{
					Text:  "Serial Number: ",
					Style: widget.RichTextStyleStrong,
				},
				&widget.TextSegment{Text: module.SerialNumber},
			)
		}
		co.(*widget.RichText).Refresh()
	})
	tree.OpenAllBranches()
	return container.NewBorder(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("## %d memory modules:", len(modules))),
		nil, nil, nil, tree,
	)
}
