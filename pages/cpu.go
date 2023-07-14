package pages

import (
	"embed"
	"fmt"
	"qpu-z/specs"

	"qpu-z/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CPUPage(assets embed.FS) fyne.CanvasObject {
	cpu := specs.GetCPU()
	var extra map[string]string
	if cpu.Vendor == "Intel" {
		extra = util.IntelArkGetCPU(util.URLCPUName(cpu.Model))
	}
	gen := util.IntelArkGetGeneration(extra["CodeNameText"])
	if cpu.Count >= 2 {
		cpu.Model += fmt.Sprintf(" (x%d)", cpu.Count)
	}
	cores := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Cores: %d", cpu.Cores))
	threads := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Threads: %d", cpu.Threads))
	vendor := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Vendor: %s", cpu.Vendor))
	generation := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Generation: %s", gen))
	cache := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Cache: %s", extra["Cache"]))
	releaseDate := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Release Date: %s", extra["BornOnDate"]))
	maxClockSpeed := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Max Clock Speed: %s", extra["ClockSpeedMax"]))
	vt := widget.NewRichTextFromMarkdown(fmt.Sprintf("## VT-D: %s, VT-X: %s", extra["VTD"], extra["VTX"]))
	igpu := widget.NewRichTextFromMarkdown(fmt.Sprintf("## iGPU: %s", extra["ProcessorGraphicsModelId"]))
	icon := util.GetIcon(assets, cpu.Vendor, false)
	icon.SetMinSize(fyne.NewSize(215, 215))
	card := widget.NewCard(cpu.Model, "", container.NewHBox(icon, container.NewVBox(vendor, generation, releaseDate, maxClockSpeed), container.NewVBox(cores, threads, cache, vt)))
	if extra["ProcessorGraphicsModelId"] != "" {
		card.Content.(*fyne.Container).Objects[1].(*fyne.Container).Add(igpu)
	}
	return container.NewVBox(card)
}
