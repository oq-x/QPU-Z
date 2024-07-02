package main

import (
	"fmt"

	"qpu-z/pages"
	"qpu-z/util"
	"runtime"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("HacSpeccer")

	if runtime.GOOS != "darwin" {
		content := widget.NewCard("Unsupported Platform", fmt.Sprintf("HacSpeccer only works on macOS. You are using %s", runtime.GOOS), widget.NewButton("Quit", func() {
			window.Close()
		}))
		window.SetContent(content)
	} else {
		util.FetchPCIID()
		util.IORegistry = util.FetchIORegistry()
		tabs := container.NewAppTabs(
			container.NewTabItem("Board", pages.BoardPage()),
			container.NewTabItem("CPU", pages.CPUPage()),
			container.NewTabItem("GPU", pages.GPUPage()),
			container.NewTabItem("Memory", pages.MemoryPage()),
			container.NewTabItem("PCI", pages.PCIPage()),
		)
		window.SetContent(tabs)
	}

	window.ShowAndRun()
}
