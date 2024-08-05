package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Containers Demo")

	label := widget.NewLabel("Menu")
	input := widget.NewPasswordEntry()
	text := widget.NewRichTextWithText("Dummy content. Please ignore")
	vbox := container.NewVBox(text, input)

	// Hsplit
	// cont := container.NewHSplit(label, vbox)

	// App Tabs
	// item1 := container.NewTabItem("Tab1", label)
	// item2 := container.NewTabItem("Tab3", vbox)
	// cont := container.NewAppTabs(item1, item2)

	// Doc Tabs
	// cont := container.NewDocTabs(item1, item2)

	// Scroll container
	cont := container.NewScroll(container.NewHBox(label, vbox))

	w.SetContent(cont)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
