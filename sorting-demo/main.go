package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/vivekprm/fyne-learnings/sorting-demo/models"
	"github.com/vivekprm/fyne-learnings/sorting-demo/sorting"
)

const (
	width      = 800
	height     = 800
	barWidth   = 20
	barPadding = 10
)

func main() {
	os.Setenv("FYNE_THEME", "light")
	a := app.New()
	w := a.NewWindow("Insertion Sort Simulation")
	arr := []int{15, 9, 6, 11, 8, 10}
	g := models.NewBarGraph(width, height)
	graph := g.Draw()

	objs := []fyne.CanvasObject{}
	objs = append(objs, graph)
	for i := 0; i < len(arr); i++ {
		bar := models.NewBar(float32(arr[i]), float32(i*(barPadding+barWidth)), float32(arr[i]))
		barObj := bar.Draw()
		objs = append(objs, barObj)
	}
	w.SetContent(container.NewWithoutLayout(objs...))

	w.Resize(fyne.NewSize(width, height))
	w.ShowAndRun()

	sorting.InsertionSort(arr)
	print(arr)
}

func print(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Print(arr[len(arr)-1])
}
