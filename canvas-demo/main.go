package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Canvas Demo")

	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	// To move and resize wrap it into a container
	rectWithContainer := container.NewWithoutLayout(rect)
	rect.Resize(fyne.NewSize(50, 100))
	rect.Move(fyne.NewPos(10, 20))

	btn := widget.NewButton("Run", func() {
		var current fyne.Size
		initialSize := fyne.NewSize(50, 100)
		finalSize := fyne.NewSize(50, 300)
		anim := canvas.NewSizeAnimation(initialSize, finalSize, 5*time.Second, func(s fyne.Size) {
			current = s
			rect.Resize(current)
		})
		anim.Start()
	})
	vbox := container.NewVBox(rectWithContainer, btn)
	w.SetContent(vbox)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
