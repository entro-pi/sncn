package main

import (
    "strconv"
    "strings"
    "log"
    "os"
    "fmt"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
//    "github.com/gotk3/gotk3/gdk"
)




func getUserPass(twoBuilder *gtk.Builder) (string, string) {
	userUncast, err := twoBuilder.GetObject("login")
	if err != nil {
		panic(err)
	}
	userEntry := userUncast.(*gtk.Entry)
	userBuf, err := userEntry.GetBuffer()
	if err != nil {
		panic(err)
	}
//	userBuf := userBufUncast.(*gtk.EntryBuffer)
	passUncast, err := twoBuilder.GetObject("pass")
	if err != nil {
		panic(err)
	}
	passEntry := passUncast.(*gtk.Entry)
	passBuf, err := passEntry.GetBuffer()
	if err != nil {
		panic(err)
	}

//	passBuf := passBufUncast.(*gtk.EntryBuffer)
	user, err := userBuf.GetText()
	if err != nil {
		panic(err)
	}
	pass, err := passBuf.GetText()
	if err != nil {
		panic(err)
	}
	user = strings.ToUpper(user)
	return user, pass

}

func launch(application *gtk.Application, twoBuilder *gtk.Builder) {
        // Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("maininterface")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }
	exitUn, err := twoBuilder.GetObject("exitMain")
	if err != nil {
		panic(err)
	}
	exit := exitUn.(*gtk.Button)
	exit.Connect("clicked", func () {
		os.Exit(1)
	})
	invUn, err := twoBuilder.GetObject("invMain")
	if err != nil {
		panic(err)
	}
	inv := invUn.(*gtk.Button)
	inv.Connect("clicked", func () {
		boxUn, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		box := boxUn.(*gtk.Grid)
		box.SetVisible(false)
	})
	broadUn, err := twoBuilder.GetObject("broadMain")
	if err != nil {
		panic(err)
	}
	broad := broadUn.(*gtk.Button)
	broad.Connect("clicked", func () {
		boxUn, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		box := boxUn.(*gtk.Grid)
		if box.GetVisible() {
			box.SetVisible(false)
		}else {
			box.SetVisible(true)
		}
	})

	wind := appWindow.(*gtk.ApplicationWindow)
	wind.SetDefaultSize(1920, 1080)
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
	screen, err := windowWidget.GetScreen()
	if err != nil {
		panic(err)
	}
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// Set ApplicationWindow Properties
        wind.Show()
	application.AddWindow(wind)


}


func LaunchGUI() {
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
	//	loginTitle, err := twoBuilder.GetObject("loginTitle")
	//	passTitle, err := twoBuilder.GetObject("passTitle")
		view, err := twoBuilder.GetObject("syn-ack")
		if err != nil {
			panic(err)
		}
		drawField := view.(*gtk.TextView)
		draw, err := drawField.GetBuffer()
		if err != nil {
			panic(err)
		}
		yesButton, err := twoBuilder.GetObject("b1")
		if err != nil {
			panic(err)
		}
		yes := yesButton.(*gtk.Button)
		yes.Connect("clicked", func (btn *gtk.Button) {
			os.Exit(1)
		})
		noButton, err := twoBuilder.GetObject("b2")
		if err != nil {
			panic(err)
		}
		no := noButton.(*gtk.Button)
		no.Connect("clicked", func (btn *gtk.Button) {
			user, pass := getUserPass(twoBuilder)
			userCaps := strings.ToUpper(user)
			draw.SetText(userCaps+"-ACK")
			fmt.Print(pass)
			fmt.Println("b2 clicked")
			if userCaps == "WEASEL" && pass == "lol" {
				launch(application, twoBuilder)
			}
		})



	}

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
	screen, err := windowWidget.GetScreen()
	if err != nil {
		panic(err)
	}

	rows := 7
	cols := 4
	count := 0
	for r := 0;r < (rows);r++ {
		for c := 0;c < (cols);c++ {
			messageName := fmt.Sprint("message"+strconv.Itoa(count))
			messageUncast, err := twoBuilder.GetObject(messageName)
			if err != nil {
				panic(err)
			}
			message := messageUncast.(*gtk.Label)
			message.SetText("dorp dorp dorp dorp!")
			count++
		}
	}
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// Set ApplicationWindow Properties
        wind.Show()
	application.AddWindow(wind)
    })
    var placeholder []string
    // Run Gtk application
    application.Run(placeholder)
}
