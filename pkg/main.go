package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/vivekprm/fyne-learnings/pkg/models"
	"github.com/vivekprm/fyne-learnings/pkg/sorting"
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

	sorting.InsertionSort(arr)
	print(arr)
}

func print(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Print(arr[len(arr)-1])
}
