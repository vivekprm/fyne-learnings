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
	data   []Series
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

	for i := 0; i < len(g.data); i++ {
		xPos := g.data[i].x + float32(i*xPadding)
		barVal := g.data[i].val
		yPos := g.height - 100 - barVal*barHeightMultiplier
		bar := NewBar(barVal, xPos, yPos)
		barObj := bar.Draw()
		obj = append(obj, barObj)
	}

	return container.NewWithoutLayout(obj...)
}

func NewBarGraph(arr []int, width, height float32) Graph {
	var data []Series
	pos := xPadding

	for i := 0; i < len(arr); i++ {
		data = append(data, Series{
			val: float32(arr[i]),
			x:   float32(pos),
		})
		pos += barWidth
	}
	return Graph{
		data:   data,
		width:  width,
		height: height,
	}
}
