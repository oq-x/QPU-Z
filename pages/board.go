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

func HackintoshBoardPage() fyne.CanvasObject {
	board, ok := specs.GetBoard()
	if !ok {
		text := widget.NewRichTextFromMarkdown("## If you are running a hackintosh, please make sure `Misc -> Security -> ExposeSensitiveData` is set to `8` or higher in `config.plist`.")
		return widget.NewCard("Couldn't get board data.", "", text)
	}
	vendor := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Vendor: %s", board.Vendor))
	model := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Model: %s", board.Model))
	return container.NewVBox(vendor, model)
}

func AppleBoardPage() fyne.CanvasObject {
	appleBoard = specs.GetAppleBoard()
	manufacturer := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Manufacturer: %s", appleBoard.Manufacturer))
	model := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Model: %s", appleBoard.Model))
	serialNumber := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Serial Number: %s", appleBoard.SerialNumber))
	boardID := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Board ID: %s", appleBoard.BoardID))
	return container.NewVBox(manufacturer, model, serialNumber, boardID)
}

func BoardPage() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Apple", AppleBoardPage()),
		container.NewTabItem("Hackintosh", HackintoshBoardPage()),
	)
	if strings.HasPrefix(appleBoard.Manufacturer, "Apple") {
		tabs.DisableIndex(1)
	}

	return tabs
}
