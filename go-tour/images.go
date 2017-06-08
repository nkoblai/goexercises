package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image represents the custom image type
type Image struct {
	width, height int
	pColor        uint8
}

// Bounds returns the image dimensions and location
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

// ColorModel returns the image color model
func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the image color per pixel
func (i *Image) At(x, y int) color.Color {
	return color.RGBA{
		i.pColor + uint8(x), i.pColor + uint8(y), 255, 255}
}

// Define your own Image type, implement the necessary methods, and call pic.ShowImage.
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
// ColorModel should return color.RGBAModel.
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
