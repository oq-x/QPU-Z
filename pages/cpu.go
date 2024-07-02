package pages

import (
	"fmt"
	"qpu-z/specs"

	"qpu-z/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CPUPage() fyne.CanvasObject {
	cpu := specs.GetCPU()
	if cpu.Count >= 2 {
		cpu.Model += fmt.Sprintf(" (x%d)", cpu.Count)
	}
	cores := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Cores: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: fmt.Sprint(cpu.Cores),
		},
	)
	threads := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Threads: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: fmt.Sprint(cpu.Threads),
		},
	)
	vendor := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Vendor: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: cpu.Vendor,
		},
	)
	generation := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Generation: ",
			Style: widget.RichTextStyleStrong,
		},
		&widget.TextSegment{
			Text: fmt.Sprintf("%s (%d/%d)", cpu.GenerationDisplayName, cpu.GenerationFamily, cpu.GenerationModel),
		},
	)

	icon := util.GetIcon(cpu.Vendor, false)
	icon.SetMinSize(fyne.NewSize(200, 200))
	icon.FillMode = canvas.ImageFillContain
	icon.ScaleMode = canvas.ImageScaleFastest
	card := widget.NewCard(cpu.Model, "", container.NewHBox(icon, container.NewVBox(vendor, generation, cores, threads)))
	return container.NewVBox(card)
}
