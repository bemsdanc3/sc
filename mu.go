package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("хуй его знает")

	w.SetContent(container.NewVBox())

	w.ShowAndRun()
}
