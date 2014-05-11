package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
	X int
    Y int
    Width int
    Height int
}

func (i Image) Bounds() image.Rectangle{
    return image.Rect(i.X, i.Y, i.Width, i.Height)
}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i Image) At(x,y int) color.Color {
    x_1:=uint8(x*y/2)
    y_1:=uint8(y*x/2)
    return color.RGBA{x_1,y_1,255,255}
}

func main() {
    m := Image{0,0,200,200}
    pic.ShowImage(m)
}