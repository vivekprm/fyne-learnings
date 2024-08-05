package models

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	barWidth            = 30
	barHeightMultiplier = 15
	strokeWidth         = 0
)

type Bar struct {
	height    float32
	xPosition float32
	yPosition float32
}

func NewBar(height, xPosition, yPosition float32) Bar {
	return Bar{
		height:    height,
		xPosition: xPosition,
		yPosition: yPosition,
	}
}

func (b Bar) Draw() fyne.CanvasObject {
	text := canvas.NewText(strconv.FormatFloat(float64(b.height), 'f', 0, 32), color.Black)
	text.Move(fyne.NewPos(b.xPosition, b.yPosition+b.height*barHeightMultiplier))
	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	rect.Resize(fyne.NewSize(barWidth, b.height*barHeightMultiplier))
	rect.Move(fyne.NewPos(b.xPosition, b.yPosition))
	rect.StrokeColor = color.White
	rect.StrokeWidth = strokeWidth

	return container.NewWithoutLayout(text, rect)
}
