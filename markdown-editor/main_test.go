package main_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestText_Selected(t *testing.T) {
	e := widget.NewEntry()
	test.Type(e, "Hello")
	assert.Equal(t, "Hello", e.Text)

	test.DoubleTap(e)
	assert.Equal(t, "Hello", e.SelectedText())
	assert.Equal(t, 5, e.CursorColumn)
}
