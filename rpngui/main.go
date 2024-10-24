package main

import (
	"fmt"
	"rpncalc/plugins"
	"rpncalc/rpncalc"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{}

var shutdown_complete sync.WaitGroup
var shutdown chan struct{}

type TapLabel struct {
	*widget.Label
	window *fyne.Window
}

func (mc *TapLabel) Tapped(*fyne.PointEvent) {
	fmt.Printf("Tapped on '%s'\n", mc.Text)

	w := *mc.window
	w.Clipboard().SetContent(mc.Text)
}

func NewTapLabel(text string, w *fyne.Window) *TapLabel {
	return NewTapLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{}, w)
}

func NewTapLabelWithStyle(text string, alignment fyne.TextAlign, style fyne.TextStyle, w *fyne.Window) *TapLabel {
	return &TapLabel{
		widget.NewLabelWithStyle(text, alignment, style),
		w,
	}
}

func main() {
	a := app.New()
	fyne.CurrentApp().Settings().SetTheme(myTheme{})
	mainWindow := a.NewWindow("RPN Calculator")

	shutdown = make(chan struct{})

	intp := rpncalc.NewInterpreter()
	intp.AddOperators(plugins.Extended_Math_Ops)
	intp.AddOperators(plugins.Conversion_Ops)

	cfg := LoadConfig()

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return NewTapLabel("template", &mainWindow)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			lbl := o.(*TapLabel)
			lbl.SetText(data[i])
		})

	lstring := ""
	mylabelstring := binding.BindString(&lstring)

	run_and_update := func(s string) {

		err := intp.Parse(s)
		if err != nil {
			mylabelstring.Set(err.Error())
		}

		data = make([]string, len(intp.Stack))
		for i := range intp.Stack {
			data[i] = intp.Stack[i].String()
		}
		list.Refresh()
	}

	myentry := newEnterEntry(cfg)
	myentry.OnSubmitted = func(s string) {
		myentry.SetText("")
		run_and_update(s)
	}
	myentry.OnChanged = func(s string) {

	}
	mylabel := widget.NewLabelWithData(mylabelstring)

	footer := container.NewVBox(myentry, mylabel)
	file_menu_one := fyne.NewMenuItem("Test", func() { fmt.Print("test clicked") })
	file_menu := fyne.NewMenu("File", file_menu_one)

	user_menus := make([]*fyne.MenuItem, 0)
	for k := range cfg.UserMenu {
		v := cfg.UserMenu[k]
		newitem := fyne.NewMenuItem(k, func() { run_and_update(v) })
		user_menus = append(user_menus, newitem)
	}
	user_menu := fyne.NewMenu("User", user_menus...)

	op_menus := make(map[string][]*fyne.MenuItem)
	ops := intp.Operators()
	for i := range ops {
		op := ops[i]
		names := strings.Split(op.Name, ".")
		prefix := names[0]
		name := names[0]
		if len(names) == 1 {
			prefix = "builtin"
		} else {
			name = names[1]
		}

		m, ok := op_menus[prefix]
		if !ok {
			m = make([]*fyne.MenuItem, 0)
		}

		newitem := fyne.NewMenuItem(name, func() { run_and_update(op.Name) })
		m = append(m, newitem)
		op_menus[prefix] = m
	}

	menus := make([]*fyne.Menu, 2)
	menus[0] = file_menu
	menus[1] = user_menu
	for k := range op_menus {
		v := op_menus[k]
		new_menu := fyne.NewMenu(k, v...)
		menus = append(menus, new_menu)
	}

	main_menu := fyne.NewMainMenu(menus...)

	favorites := make([]fyne.CanvasObject, 0)
	for k := range cfg.Favorites {
		v := cfg.Favorites[k]
		btn := widget.NewButton(k,
			func() {
				run_and_update(v)
			})
		favorites = append(favorites, btn)
	}

	quick_bar := container.NewVBox(favorites...)
	ui := container.New(layout.NewBorderLayout(nil, footer, nil, quick_bar), footer, list, quick_bar)
	mainWindow.SetMainMenu(main_menu)

	mainWindow.SetContent(ui)
	icon := fyne.NewStaticResource("calcicon", iconpng)
	mainWindow.SetIcon(icon)

	mainWindow.Resize(fyne.NewSize(300, 400))
	mainWindow.Canvas().Focus(myentry)

	mainWindow.ShowAndRun()

	close(shutdown)
	shutdown_complete.Wait()
}
