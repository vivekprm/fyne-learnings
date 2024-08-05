package models

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	xPadding = 20
	yPadding = 100
)

type Series struct {
	val float32
	x   float32
}

type Graph struct {
	width  float32
	height float32
}

func (g Graph) Draw() fyne.CanvasObject {
	obj := []fyne.CanvasObject{}

	lineX := canvas.NewLine(color.Black)
	lineX.StrokeWidth = 1
	lineX.StrokeColor = color.Black
	lineX.Position1 = fyne.NewPos(xPadding, g.height-yPadding)
	lineX.Position2 = fyne.NewPos(g.width-100, g.height-yPadding)

	lineY := canvas.NewLine(color.Black)
	lineY.StrokeWidth = 1
	lineY.StrokeColor = color.Black
	lineY.Position1 = fyne.NewPos(xPadding, g.height-yPadding)
	lineY.Position2 = fyne.NewPos(xPadding, g.height-(g.height-yPadding))
	obj = append(obj, lineX, lineY)
	return container.NewWithoutLayout(obj...)
}

func NewBarGraph(width, height float32) Graph {
	return Graph{
		width:  width,
		height: height,
	}
}
