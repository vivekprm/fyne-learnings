package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	barWidth            = 30
	barHeightMultiplier = 8
	barPadding          = 5
)

func main() {
	a := app.New()
	w := a.NewWindow("Sorting demo")
	bar1 := createBar(18)
	bar2 := createBar(11)
	bar2.Move(fyne.NewPos(barWidth+barPadding, 0))
	hb := container.NewWithoutLayout(bar1, bar2)

	sPos := bar1.Position()
	ePos := bar2.Position()
	canvas.NewPositionAnimation(sPos, ePos, time.Second*2, func(p fyne.Position) {
		bar1.Move(p)
		canvas.Refresh(bar1)
	}).Start()
	canvas.NewPositionAnimation(ePos, sPos, time.Second*2, func(p fyne.Position) {
		bar2.Move(p)
		canvas.Refresh(bar2)
	}).Start()

	w.SetContent(hb)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

func createBar(val int) *fyne.Container {
	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	rectc := container.NewWithoutLayout(rect)
	rect.Resize(fyne.NewSize(barWidth, float32(val*barHeightMultiplier)))
	rect.Refresh()

	text := canvas.NewText(fmt.Sprintf("%d", val), color.White)
	text.Move(fyne.NewPos(0+(barWidth-text.TextSize)/2, 0))
	textc := container.NewWithoutLayout(text)
	hb := container.NewVBox(rectc, textc)
	return hb
}
