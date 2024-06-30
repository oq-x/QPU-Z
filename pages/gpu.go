package pages

import (
	"embed"
	"fmt"
	"qpu-z/specs"
	"qpu-z/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GPUPage(assets embed.FS) fyne.CanvasObject {
	gpus := specs.GetGPUs()
	cont := container.NewVBox()
	for _, gpu := range gpus {
		vendorText := gpu.Vendor
		if gpu.VendorID != "" {
			vendorText += fmt.Sprintf(" (0x%s)", gpu.VendorID)
		}
		vendorLogo := util.GetIcon(assets, gpu.Vendor, true)
		vendorLogo.FillMode = canvas.ImageFillContain
		vendorLogo.ScaleMode = canvas.ImageScaleFastest

		vendor := widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyleStrong,
				Text:  "Vendor ID: ",
			},
			&widget.TextSegment{
				Text: vendorText,
			},
		)
		devid := widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyleStrong,
				Text:  "Device ID: ",
			},
			&widget.TextSegment{
				Text: "0x" + gpu.DeviceID,
			},
		)
		core := widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyleStrong,
				Text:  "Core: ",
			},
			&widget.TextSegment{
				Text: gpu.Core,
			},
		)
		vram := widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyleStrong,
				Text:  "Video-RAM: ",
			},
			&widget.TextSegment{
				Text: gpu.VRAM,
			},
		)
		model := widget.NewRichTextFromMarkdown(fmt.Sprintf("# %s", gpu.Model))

		internalText := widget.NewLabel("Internal")
		internalText.TextStyle.Bold = true
		internalRect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		internalRect.CornerRadius = 20
		internal := container.NewStack(internalRect, internalText)

		card := widget.NewCard("", "", container.NewVBox(container.NewHBox(model), container.NewHBox(vendorLogo, container.NewVBox(vendor, devid))))
		if gpu.Internal {
			card.Content.(*fyne.Container).Objects[0].(*fyne.Container).Add(container.NewCenter(internal))
		}
		if gpu.Core != "" {
			card.Content.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*fyne.Container).Add(core)
		}
		if gpu.VRAM != "" {
			card.Content.(*fyne.Container).Objects[1].(*fyne.Container).Objects[1].(*fyne.Container).Add(vram)
		}
		cont.Add(card)
	}

	return cont
}
