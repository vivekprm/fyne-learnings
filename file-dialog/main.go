package main

import (
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("File Dialog")
	e := widget.NewEntry()
	btn := widget.NewButton("Open", func() {
		open(w, e)
	})
	cont := container.NewVBox(e, btn)
	cont.Resize(fyne.NewSize(150, 20))
	w.SetContent(cont)
	w.ShowAndRun()
}

func open(window fyne.Window, e *widget.Entry) {
	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
		}
		if r == nil {
			return
		}
		data, err := ioutil.ReadAll(r)
		if err == nil {
			_ = r.Close()
			e.SetText(string(data))
			saveURI := r.URI()
			fmt.Println(saveURI)
		} else {
			dialog.ShowError(err, window)
		}
	}, window)
}
