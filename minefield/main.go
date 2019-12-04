package main

import (
    "strings"
    "log"
    "os"
    "fmt"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
//    "github.com/gotk3/gotk3/gdk"
)


func main() {
    // Create Gtk Application, change appID to your application domain name reversed.
    const appID = "org.gtk.sncn"
    application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
    // Check to make sure no errors when creating Gtk Application
    if err != nil {
        log.Fatal("Could not create application.", err)
    }

    // Application signals available
    // startup -> sets up the application when it first starts
    // activate -> shows the default first window of the application (like a new document). This corresponds to the application being launched by the desktop environment.
    // open -> opens files and shows them in a new window. This corresponds to someone trying to open a document (or documents) using the application from the file browser, or similar.
    // shutdown ->  performs shutdown tasks
    // Setup activate signal with a closure function.
    application.Connect("activate", func() {
	    twoBuilder, err := gtk.BuilderNewFromFile("twobutton.glade")
	    if err != nil {
		panic(err)
		}
	if err == nil {
		view, err := twoBuilder.GetObject("view1")
		yesButton, err := twoBuilder.GetObject("b1")
		if err != nil {
			panic(err)
		}
		yes := yesButton.(*gtk.Button)
		yes.Connect("clicked", func (btn *gtk.Button) {
			text := view.(*gtk.TextView)
			draw, err := text.GetBuffer()
			if err != nil {
				panic(err)
			}
			draw.SetText("ACK")
			fmt.Println("b1 clicked")
			os.Exit(1)
		})
		noButton, err := twoBuilder.GetObject("b2")
		if err != nil {
			panic(err)
		}
		no := noButton.(*gtk.Button)
		no.Connect("clicked", func (btn *gtk.Button) {
			text := view.(*gtk.TextView)
			draw, err := text.GetBuffer()
			if err != nil {
				panic(err)
			}
			draw.SetText("ACK")
			fmt.Println("b1 clicked")
		})
		userField, err := twoBuilder.GetObject("loginbuffer")
		if err != nil {
			panic(err)
		}
		passField, err := twoBuilder.GetObject("passwordbuffer")
		if err != nil {
			panic(err)
		}

		user := userField.(*gtk.TextBuffer)
		pass := passField.(*gtk.TextBuffer)
		limit := 25
		user.Connect("insert-text", func (textBuf *gtk.TextBuffer) {
			start, end := textBuf.GetBounds()
			text, err := textBuf.GetText(start, end, true)
			if len(strings.Split(text, "\n")) > 1 {
				textBuf.SetText(strings.Split(text, "\n")[0])
			}
			if err != nil {
				fmt.Printf("", err)
			}
			err = nil
			if len(text) >= limit {
				textBuf.SetText(text[len(text)-1:])
			}
		})
		pass.Connect("insert-text", func (textBuf *gtk.TextBuffer) {
			start, end := textBuf.GetBounds()
			text, err := textBuf.GetText(start, end, true)
			if len(strings.Split(text, "\n")) > 1 {
				textBuf.SetText(strings.Split(text, "\n")[0])
			}
			if err != nil {
				fmt.Printf("", err)
			}
			err = nil
			if len(text) >= limit {
				textBuf.SetText(text[len(text)-1:])
			}

		})


	}
//	twoBuilder.ConnectSignals(signals)

        // Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("mainwindow")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }

	wind := appWindow.(*gtk.ApplicationWindow)

	wind.SetDefaultSize(400, 400)
	wind.SetResizable(false)
	wind.SetPosition(gtk.WIN_POS_CENTER)
	windowWidget, err := wind.GetStyleContext()
	if err != nil {
		panic(err)
	}

	css, err := gtk.CssProviderNew()
	if err != nil {
		panic(err)
	}

	css.LoadFromPath("design.css")
	user, err := twoBuilder.GetObject("user")
	userView := user.(*gtk.TextView)
	userView.SetWrapMode(gtk.WRAP_NONE)
	if err != nil {
		panic(err)
	}
	pass, err := twoBuilder.GetObject("pass")
	passView := pass.(*gtk.TextView)
	passView.SetWrapMode(gtk.WRAP_NONE)
	if err != nil {
		panic(err)
	}
	screen, err := windowWidget.GetScreen()
	if err != nil {
		panic(err)
	}
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// Set ApplicationWindow Properties
        wind.Show()
	application.AddWindow(wind)
    })
    // Run Gtk application
    application.Run(os.Args)
}
