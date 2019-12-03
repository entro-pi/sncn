package main

import (
    "log"
    "os"
    "fmt"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

// Simple Gtk3 Application written in go.
// This application creates a window on the application callback activate.
// More GtkApplication info can be found here -> https://wiki.gnome.org/HowDoI/GtkApplication

func b1Clicked() {
	fmt.Println("b1 clicked")
}

func b2Clicked() {
	fmt.Println("b2 clicked")
}


// you just place them in a map that names the signals, then feed the map to the builder
var signals = map[string]interface{}{
	"B1": b1Clicked,
	"B2": b2Clicked,
}


func main() {
    // Create Gtk Application, change appID to your application domain name reversed.
    const appID = "org.gtk.example"
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

	twoBuilder.ConnectSignals(signals)

        // Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("mainwindow")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }
	wind := appWindow.(*gtk.Window)
        // Set ApplicationWindow Properties
        wind.Show()
	application.AddWindow(wind)
    })
    // Run Gtk application
    application.Run(os.Args)
}
