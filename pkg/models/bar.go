package models

import (
	"fmt"
	"image/color"

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
	// xPos := float32(b.position) + 10
	// yPos := b.height - b.height*barHeightMultiplier
	text := canvas.NewText(fmt.Sprintf("%f", b.height), color.Black)
	text.Move(fyne.NewPos(b.xPosition+strokeWidth, b.height-100))
	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	rect.Resize(fyne.NewSize(barWidth, b.height*barHeightMultiplier))
	rect.Move(fyne.NewPos(b.xPosition, b.yPosition))
	rect.StrokeColor = color.White
	rect.StrokeWidth = strokeWidth

	return container.NewWithoutLayout(text, rect)
}
