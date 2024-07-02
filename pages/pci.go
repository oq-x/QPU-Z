package pages

import (
	"fmt"
	"qpu-z/util"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func PCIPage() fyne.CanvasObject {
	accel := util.IORegistry.GPUAccelerated
	accelText := widget.NewRichText(
		&widget.TextSegment{
			Style: widget.RichTextStyle{
				ColorName: util.Condition(accel, theme.ColorNameSuccess, theme.ColorNameError),
				TextStyle: fyne.TextStyle{
					Bold: true,
				},
			},
			Text: util.Condition(accel, "Graphical Acceleration Enabled", "No Graphical Acceleration"),
		},
	)

	var deviceTree = container.NewVBox()

	for _, d := range util.IORegistry.PciDevices {
		vid, _ := d["vendor-id"].([]byte)
		id, _ := d["device-id"].([]byte)

		svid, _ := d["subsystem-vendor-id"].([]byte)
		sid, _ := d["subsystem-id"].([]byte)

		kexts, kok := d["IORegistryEntryChildren"].([]any)
		var kextamount int

		vendor, ok := util.PCIDeviceCache.Vendor(util.PlistDataToLEHex(vid))
		if !ok {
			continue
		}
		var name string
		device, ok := vendor.Device(util.PlistDataToLEHex(id))
		if !ok {
			name = fmt.Sprintf("%s (%s:%s)",
				string(d["name"].([]byte)),
				util.PlistDataToLEHex(vid),
				util.PlistDataToLEHex(id),
			)
		} else {
			name = device.Name
			if len(svid) != 0 && len(sid) != 0 {
				sub, ok := device.Subsystem(util.PlistDataToLEHex(svid), util.PlistDataToLEHex(sid))
				if ok {
					name = sub.Name
				}
			}
		}

		name = strings.TrimSpace(name)

		model, ok := d["model"]
		if ok {
			name = string(model.([]byte))
		}

		rt := widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyleSubHeading,

				Text: name,
			},
		)

		classcode := d.DataKey("class-code")
		if len(classcode) == 4 {
			cl := fmt.Sprintf("%02x", classcode[2])
			subcl := fmt.Sprintf("%02x", classcode[1])

			class, ok := util.PCIClassCache.Subclass(cl, subcl)
			if ok {
				rt.Segments = append(rt.Segments,
					&widget.TextSegment{
						Style: widget.RichTextStyleStrong,

						Text: "Type: ",
					},
					&widget.TextSegment{Text: class.Name},
				)
			}
		}

		if kok {
			for _, k := range kexts {
				kext := k.(map[string]any)
				bundid, ok := kext["CFBundleIdentifier"]
				if !ok {
					continue
				}
				rt.Segments = append(rt.Segments, &widget.TextSegment{
					Text: fmt.Sprintf("%s (%s)", kext["IORegistryEntryName"].(string), bundid),
				})
				kextamount++
			}
		}
		var kstatus = widget.NewRichText(
			&widget.TextSegment{
				Style: widget.RichTextStyle{
					SizeName:  theme.SizeNameSubHeadingText,
					TextStyle: widget.RichTextStyleStrong.TextStyle,
					ColorName: theme.ColorNameError,
				},
				Text: "0 kexts",
			},
		)
		if kextamount != 0 {
			kstatus = widget.NewRichText(
				&widget.TextSegment{
					Style: widget.RichTextStyle{
						SizeName:  theme.SizeNameSubHeadingText,
						TextStyle: widget.RichTextStyleStrong.TextStyle,
						ColorName: theme.ColorNameSuccess,
					},
					Text: fmt.Sprintf("%d kexts", kextamount),
				},
			)
		}

		deviceTree.Add(
			container.NewHBox(rt, layout.NewSpacer(), kstatus),
		)
	}

	return container.NewBorder(
		accelText, nil, nil, nil, container.NewVScroll(deviceTree),
	)
}
