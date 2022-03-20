package windigoui

import (
	"runtime"

	"github.com/Elren44/qr_code_gui/qrgen"
	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
)

func RunWindigo() {
	runtime.LockOSThread()

	myWindow := NewMyWindow() // instantiate
	myWindow.wnd.RunAsMain()  // ...and run
}

// This struct represents our main window.
type MyWindow struct {
	wnd     ui.WindowMain
	lblName ui.Static
	txtName ui.Edit
	btnShow ui.Button
}

// Creates a new instance of our main window.
func NewMyWindow() *MyWindow {
	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("Elren-QR-Code-Generator").
			ClientArea(win.SIZE{Cx: 500, Cy: 100}),
		//IconId(101), // ID of icon resource, see resources folder
	)

	me := &MyWindow{
		wnd: wnd,
		lblName: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text("Введи текст").
				Position(win.POINT{X: 10, Y: 22}),
		),
		txtName: ui.NewEdit(wnd,
			ui.EditOpts().
				Position(win.POINT{X: 80, Y: 20}).
				Size(win.SIZE{Cx: 350}),
		),
		btnShow: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Генерировать").
				Position(win.POINT{X: 140, Y: 60}),
		),
	}

	me.btnShow.On().BnClicked(func() {
		//msg := fmt.Sprintf("Hello, %s!", me.txtName.Text())
		t := me.txtName.Text()
		qrgen.QrGen(t)

		//me.wnd.Hwnd().MessageBox(msg, "Saying hello", co.MB_ICONINFORMATION)
	})

	return me
}
