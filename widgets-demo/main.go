package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Widgets Demo")
	label := widget.NewLabel("Label")
	cancelBtn := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})
	cont := container.NewGridWithRows(8)
	pb := widget.NewProgressBar()
	content := widget.NewRichTextWithText("Hello there! This is dummy content.")
	card := widget.NewCard("Title", "Subtitle", content)

	data := map[int]string{
		1: "List Item 1",
		2: "List Item 2",
		3: "List Item 3",
		4: "List Item 4",
		5: "List Item 5",
		6: "List Item 6",
		7: "List Item 7",
		8: "List Item 8",
		9: "List Item 9",
	}
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template Object")
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[id])
		},
	)
	cont.Add(label)
	cont.Add(cancelBtn)
	cont.Add(pb)
	cont.Add(card)
	cont.Add(list)
	w.SetContent(cont)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
