package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dialog Demo")

	btn := widget.NewButton("Open", func() {
		dialog.ShowConfirm("Confirmation", "Do you want to delete this item?", func(b bool) {
			fmt.Println(b)
		}, w)
	})
	cont := container.NewVBox(btn)
	cont.Resize(fyne.NewSize(150, 20))
	w.SetContent(cont)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
