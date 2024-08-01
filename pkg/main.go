package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/vivekprm/fyne-learnings/pkg/models"
)

const (
	width  = 800
	height = 800
)

func main() {
	os.Setenv("FYNE_THEME", "light")
	a := app.New()
	w := a.NewWindow("Insertion Sort Simulation")
	arr := []int{15, 9, 6, 11, 8, 10}
	g := models.NewBarGraph(arr, width, height)
	graph := g.Draw()
	w.SetContent(graph)
	w.Resize(fyne.NewSize(width, height))
	w.ShowAndRun()
}
