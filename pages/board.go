package pages

import (
	"fmt"
	"qpu-z/specs"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var appleBoard specs.AppleBoard

func OpenCoreBoardPage() fyne.CanvasObject {
	board, ok := specs.GetBoard()
	if !ok {
		text := widget.NewRichTextFromMarkdown("If you are running the OpenCore bootloader, please make sure `Misc -> Security -> ExposeSensitiveData` is set to `8` or higher (`10` is recommended) in `config.plist`.")
		text.Wrapping = fyne.TextWrapBreak
		return widget.NewCard("Can't get board data", "", text)
	}
	vendor := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Vendor: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: board.Vendor,
		},
	)
	model := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Model: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: fmt.Sprintf("%s (%s)", board.Product, board.Board),
		},
	)
	opencoreVersion := widget.NewRichText(
		&widget.TextSegment{
			Text:  "OpenCore Version: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: board.OpenCoreVersion,
		},
	)
	return container.NewVBox(vendor, model, opencoreVersion)
}

func AppleBoardPage() fyne.CanvasObject {
	appleBoard = specs.GetAppleBoard()
	manufacturer := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Manufacturer: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: appleBoard.Manufacturer,
		},
	)
	model := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Model: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: appleBoard.Model,
		},
	)
	serialNumber := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Serial Number: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: appleBoard.SerialNumber,
		},
	)
	boardID := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Board ID: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: appleBoard.BoardID,
		},
	)
	return container.NewVBox(manufacturer, model, serialNumber, boardID)
}

func BoardPage() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Apple", AppleBoardPage()),
		container.NewTabItem("OpenCore", OpenCoreBoardPage()),
	)
	if strings.HasPrefix(appleBoard.Manufacturer, "Apple") {
		tabs.DisableIndex(1)
	}

	return tabs
}
