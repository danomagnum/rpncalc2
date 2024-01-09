package main

import (
	"fmt"
	"rpncalc/plugins"
	"rpncalc/rpncalc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{}

func makeUI(intp *rpncalc.Interpreter) (fyne.CanvasObject, *fyne.MainMenu) {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	estring := ""
	lstring := "Test"
	myentrystring := binding.BindString(&estring)
	mylabelstring := binding.BindString(&lstring)
	myentry := widget.NewEntryWithData(myentrystring)
	myentry.OnSubmitted = func(s string) {
		err := intp.Parse(s)
		if err != nil {
			mylabelstring.Set(err.Error())
		}
		myentry.SetText("")
		//mylabelstring.Set(s)
		data = make([]string, len(intp.Stack))
		for i := range intp.Stack {
			data[i] = intp.Stack[i].String()
		}
		//data = append(data, s)
		list.Refresh()
	}
	mylabel := widget.NewLabelWithData(mylabelstring)
	/*
		mybutton := widget.NewButton("Convert",
			func() {
				myinput, _ := strconv.Atoi(estring)
				myinput = myinput * 3
				myoutput := strconv.Itoa(myinput)
				mylabelstring.Set(myoutput)
			})
	*/
	footer := container.NewVBox(myentry, mylabel)
	file_menu_one := fyne.NewMenuItem("Test", func() { fmt.Print("test clicked") })
	file_menu := fyne.NewMenu("File", file_menu_one)
	main_menu := fyne.NewMainMenu(file_menu)
	return container.New(layout.NewBorderLayout(nil, footer, nil, nil), footer, list), main_menu
	//return container.NewVBox(list, myentry, mylabel)
}

func main() {
	a := app.New()
	w := a.NewWindow("Widget Binding")

	intp := rpncalc.NewInterpreter()
	intp.AddOperators(plugins.Extended_Math_Ops)
	intp.AddOperators(plugins.Conversion_Ops)

	ui, menu := makeUI(intp)
	w.SetMainMenu(menu)

	w.SetContent(ui)

	w.Resize(fyne.NewSize(300, 400))

	w.ShowAndRun()
}
