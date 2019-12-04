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

func getUserPass(twoBuilder *gtk.Builder) (string, string) {
	userBufUncast, err := twoBuilder.GetObject("loginbuffer")
	if err != nil {
		panic(err)
	}
	userBuf := userBufUncast.(*gtk.TextBuffer)
	start, end := userBuf.GetBounds()
	passBufUncast, err := twoBuilder.GetObject("passwordbuffer")
	if err != nil {
		panic(err)
	}
	passBuf := passBufUncast.(*gtk.TextBuffer)
	startP, endP := passBuf.GetBounds()
	user, err := userBuf.GetText(start, end, false)
	if err != nil {
		panic(err)
	}
	pass, err := passBuf.GetText(startP, endP, false)
	if err != nil {
		panic(err)
	}

	return user, pass

}

func launch(application *gtk.Application, twoBuilder *gtk.Builder) {
        // Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("smalltalkwindow")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }

	wind := appWindow.(*gtk.ApplicationWindow)

	wind.SetDefaultSize(1920, 1000)
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
	//	loginTitle, err := twoBuilder.GetObject("loginTitle")
	//	passTitle, err := twoBuilder.GetObject("passTitle")
		/*view, err := twoBuilder.GetObject("syn-ack")
		if err != nil {
			panic(err)
		}*/
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
			drawBoof, err := twoBuilder.GetObject("buf1")
			if err != nil {
				panic(err)
			}
			draw := drawBoof.(*gtk.TextBuffer)
			user, pass := getUserPass(twoBuilder)
			userCaps := strings.ToUpper(user)
			draw.SetText(userCaps+"-ACK")
			fmt.Print(pass)
			fmt.Println("b2 clicked")
			if userCaps == "WEASEL" && pass == "lol" {
				launch(application, twoBuilder)
			}
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
		user.Connect("insert-text", func (textBuf *gtk.TextBuffer) {
			start, end := textBuf.GetBounds()
			text, err := textBuf.GetText(start, end, true)
			if err != nil {
				fmt.Printf("", err)
			}/*
			if strings.Contains(text, "\n") {
				//start.BackwardChars(1)
				//textBuf.Delete(start, end)
				btnHold, err := twoBuilder.GetObject("b2")
				if err != nil {
					panic(err)
				}
				btn := btnHold.(*gtk.Button)
				btn.Clicked()
			}*/
			if len(strings.Split(text, "\n")) > 1 {
				textBuf.SetText(strings.Split(text, "\n")[0])
			}
			err = nil
		})
		pass.Connect("insert-text", func (textBuf *gtk.TextBuffer) {

			start, end := textBuf.GetBounds()
/*			tagTable, err := textBuf.GetTagTable()
			if err != nil {
				panic(err)
			}
			greenTag, err := tagTable.Lookup("greenTag2")
			if err != nil {
				panic(err)
			}
			textBuf.ApplyTag(greenTag, start, end)
*/

			text, err := textBuf.GetText(start, end, true)
			if err != nil {
				fmt.Printf("", err)
			}
/*			if strings.Contains(text, "\n") {
//				end.BackwardChars(1)
				start, end = textBuf.GetBounds()
				//newEnd := textBuf.GetEndIter()
//				textBuf.Modified(true)
				textBuf.Delete(start, end)
				start, _ = textBuf.GetBounds()
				textBuf.Insert(start, "NOOT")
//				btnHold, err := twoBuilder.GetObject("b2")
//				if err != nil {
//					panic(err)
//				}
//				btn := btnHold.(*gtk.Button)
//				btn.Clicked()
			}*/
			if len(strings.Split(text, "\n")) > 1 {
				textBuf.SetText(strings.Split(text, "\n")[0])
			}
			err = nil

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
