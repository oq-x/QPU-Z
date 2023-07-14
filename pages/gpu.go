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

func GPUPage(assets embed.FS) fyne.CanvasObject {
	gpus := specs.GetGPUs()
	cont := container.NewVBox()
	for _, gpu := range gpus {
		vendorText := gpu.Vendor
		if gpu.VendorID != "" {
			vendorText += fmt.Sprintf(" (%s)", gpu.VendorID)
		}
		vendor := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Vendor ID: %s", vendorText))
		vendorLogo := util.GetIcon(assets, gpu.Vendor, true)
		devid := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Device ID: %s", gpu.DeviceID))
		core := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Core: %s", gpu.Core))
		vram := widget.NewRichTextFromMarkdown(fmt.Sprintf("## VRAM: %s", gpu.VRAM))
		model := widget.NewRichTextFromMarkdown(fmt.Sprintf("# %s", gpu.Model))
		internal := widget.NewRichTextFromMarkdown("### Internal")
		card := widget.NewCard("", "", container.NewVBox(container.NewHBox(model), container.NewHBox(vendorLogo, container.NewVBox(vendor, devid), container.NewVBox())))
		if gpu.Internal {
			card.Content.(*fyne.Container).Objects[0].(*fyne.Container).Add(internal)
		}
		if gpu.Core != "" {
			card.Content.(*fyne.Container).Objects[1].(*fyne.Container).Objects[2].(*fyne.Container).Add(core)
		}
		if gpu.VRAM != "" {
			card.Content.(*fyne.Container).Objects[1].(*fyne.Container).Objects[2].(*fyne.Container).Add(vram)
		}
		cont.Add(card)
	}

	return cont
}
