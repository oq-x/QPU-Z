package util

import (
	"fmt"
	"qpu-z/assets"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func GetIcon(vendor string, isRadeon bool) *canvas.Image {
	low := strings.ToLower(vendor)
	return Icon(low, isRadeon)
}

func Icon(name string, isRadeon bool) *canvas.Image {
	if name == "amd" && isRadeon {
		name += "-radeon"
	}
	resource, _ := assets.Assets.ReadFile(fmt.Sprintf("assets/%s.png", name))
	image := canvas.NewImageFromResource(fyne.NewStaticResource(name+"_icon", resource))
	image.SetMinSize(fyne.NewSize(84, 84))
	return image
}
