# Widget - Collection
```go
widget.NewList(
    func() int {
        return len(data)
    },
    func() fyne.CanvasObejct {
        return widget.NewLabel("Template Object")
    },
    func(id widget.ListItemID, o fyne.CanvasObject) {
        o.(*widget.Label).SetText(data[id])
    },
)
```