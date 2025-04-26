package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Print("Hello World, First!\n")
	a := app.New()
	w := a.NewWindow("Hello World")

	fmt.Print("Hello World!\n")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
