package util

import (
	"embed"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func GetIcon(assets embed.FS, vendor string, isRadeon bool) *canvas.Image {
	low := strings.ToLower(vendor)
	return Icon(assets, low, isRadeon)
}

func Icon(assets embed.FS, name string, isRadeon bool) *canvas.Image {
	if name == "amd" && isRadeon {
		name += "-radeon"
	}
	resource, _ := assets.ReadFile(fmt.Sprintf("assets/%s.png", name))
	image := canvas.NewImageFromResource(fyne.NewStaticResource(name+"_icon", resource))
	image.SetMinSize(fyne.NewSize(84, 84))
	return image
}
