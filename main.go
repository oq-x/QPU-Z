package main

import (
	"embed"
	"fmt"

	"qpu-z/pages"
	"qpu-z/util"
	"runtime"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/*
var assets embed.FS

func main() {
	app := app.New()
	window := app.NewWindow("QPU-Z")

	if runtime.GOOS != "darwin" {
		content := widget.NewCard("Unsupported Platform", fmt.Sprintf("QPU-Z only works on macOS. You are using %s", runtime.GOOS), widget.NewButton("Quit", func() {
			window.Close()
		}))
		window.SetContent(content)
	} else {
		util.IORegistry = util.FetchIORegistry()
		tabs := container.NewAppTabs(
			container.NewTabItem("Board", pages.BoardPage()),
			container.NewTabItem("CPU", pages.CPUPage(assets)),
			container.NewTabItem("GPU", pages.GPUPage(assets)),
			container.NewTabItem("Memory", pages.MemoryPage()),
			container.NewTabItem("PCI", widget.NewLabel("World!")),
		)
		window.SetContent(tabs)
	}

	window.ShowAndRun()
}
