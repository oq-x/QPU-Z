package pages

import (
	"fmt"
	"qpu-z/specs"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
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
		stick := sticks[strings.Split(tni, "\\")[0]]
		co.(*widget.RichText).Segments = []widget.RichTextSegment{
			&widget.TextSegment{
				Text:  "Size: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: stick.Size},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Type: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: stick.Type},

			&widget.SeparatorSegment{},

			&widget.TextSegment{
				Text:  "Speed: ",
				Style: widget.RichTextStyleStrong,
			},
			&widget.TextSegment{Text: stick.Speed},
		}

		if stick.SerialNumber != "" {
			co.(*widget.RichText).Segments = append(co.(*widget.RichText).Segments,
				&widget.SeparatorSegment{},
				&widget.TextSegment{
					Text:  "Serial Number: ",
					Style: widget.RichTextStyleStrong,
				},
				&widget.TextSegment{Text: stick.SerialNumber},
			)
		}
		co.(*widget.RichText).Refresh()
	})
	tree.OpenAllBranches()
	return tree
}
