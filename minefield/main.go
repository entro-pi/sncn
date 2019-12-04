package main

import (
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
			draw.SetText("YES")
			fmt.Println("b1 clicked")
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
			draw.SetText("no")
			fmt.Println("b1 clicked")
		})
	}
//	twoBuilder.ConnectSignals(signals)

        // Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("mainwindow")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }

	wind := appWindow.(*gtk.Window)

	wind.SetDefaultSize(400, 400)
	windowWidget, err := wind.GetStyleContext()
	if err != nil {
		panic(err)
	}

	css, err := gtk.CssProviderNew()
	if err != nil {
		panic(err)
	}

	css.LoadFromPath("design.css")
	
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
