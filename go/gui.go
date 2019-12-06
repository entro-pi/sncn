package main

import (
    "time"
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

func launch(play Player, application *gtk.Application, twoBuilder *gtk.Builder) {
	// Create ApplicationWindow
        appWindow, err := twoBuilder.GetObject("maininterface")
        if err != nil {
            log.Fatal("Could not create application window.", err)
        }
	exitUn, err := twoBuilder.GetObject("exitMain")
	postUn, err := twoBuilder.GetObject("postBuf")
	if err != nil {
		panic(err)
	}
	post := postUn.(*gtk.TextBuffer)
	start, end := post.GetBounds()
	tagUn, err := twoBuilder.GetObject("greenTex")
	if err != nil {
		panic(err)
	}
	tag := tagUn.(*gtk.TextTag)
	post.ApplyTag(tag, start, end)
	post.Connect("insert-text", func () {
		start, end := post.GetBounds()

		post.ApplyTag(tag, start, end)
	})
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
	inv.Connect("clicked", func (button *gtk.Button) {
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
	equipUn, err := twoBuilder.GetObject("equipMain")
	if err != nil {
		panic(err)
	}
	equip := equipUn.(*gtk.Button)
	equip.Connect("clicked", func () {
		box1Un, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		box1 := box1Un.(*gtk.Grid)
		if box1.GetVisible() {
			box1.SetVisible(false)
		}else {
			box1.SetVisible(true)
		}
	})
	wind := appWindow.(*gtk.ApplicationWindow)
	wind.SetDefaultSize(1920, 1080)
	wind.SetResizable(false)
	wind.SetPosition(gtk.WIN_POS_CENTER)
	tellsUn, err := twoBuilder.GetObject("tellsMain")
	if err != nil {
		panic(err)
	}
	tells := tellsUn.(*gtk.Button)
	tells.Connect("clicked", func () {
		fill(twoBuilder, true)
		wind.SetDefaultSize(1920, 1080)
//		paintOver(twoBuilder)
/*		butt := assembleBroadButton("0")
		smallUn, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		small := smallUn.(*gtk.Grid)
		numBroad++
		colCount++
		small.Add(butt)
		//small.InsertColumn(colCount)
		wind.ShowAll()
		if colCount == 4 {
			colCount = 1
			small.InsertRow(rowCount)
			rowCount++
		}
		
		box1Un, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		box1 := box1Un.(*gtk.Grid)
		if box1.GetVisible {
			box1.SetVisible(false)
		}else {
			box1.SetVisible(true)
		}*/
	})
	broadUn, err := twoBuilder.GetObject("broadMain")
	if err != nil {
		panic(err)
	}
	broad := broadUn.(*gtk.Button)
	broad.Connect("clicked", func () {
		fill(twoBuilder, false)
		wind.SetDefaultSize(1920, 1080)
		/*
		butt := assembleBroadButton("0")
		smallUn, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		small := smallUn.(*gtk.Grid)
		numBroad++
		colCount++
		small.Add(butt)
		//small.InsertColumn(colCount)
		wind.ShowAll()
		if colCount == 4 {
			colCount = 1
			small.InsertRow(rowCount)
			rowCount++
		}*/

		/*
		box1Un, err := twoBuilder.GetObject("smalltalkgrid")
		if err != nil {
			panic(err)
		}
		box1 := box1Un.(*gtk.Grid)
		if box1.GetVisible {
			box1.SetVisible(false)
		}else {
			box1.SetVisible(true)
		}*/
	})
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

func fill(twoBuilder *gtk.Builder, tellorbroad bool)  {
	var broadcastContainer []string
	var buttonContainer []*gtk.Button
	if tellorbroad {
		play := InitPlayer("WEASEL", "lol")
		broadcastContainer = drawPlainTells(play)
	}else {
		play := InitPlayer("WEASEL", "lol")
		broadcastContainer = drawPlainBroadcasts(play)
	}
	if len(broadcastContainer) >= 20 {
		broadcastContainer = broadcastContainer[len(broadcastContainer)-20:len(broadcastContainer)]
	}
	for i := 0;i < len(broadcastContainer);i++ {
		fmt.Println(i)
		broad := assembleBroadButtonWithMessage(strconv.Itoa(i), broadcastContainer[i], twoBuilder)
		buttonContainer = append(buttonContainer, broad)
	}

	smallUn, err := twoBuilder.GetObject("smalltalkgrid")
	if err != nil {
		panic(err)
	}
	
	small := smallUn.(*gtk.Grid)
	numInRow := 4
	for i := 0;i < 12;i++ {
		small.RemoveRow(0)
	}
	row := 0
	numCount := 0
	for i := 0;i < len(buttonContainer);i++ {
		if numCount < numInRow {
//			small.Add(buttonContainer[i])
			small.Attach(buttonContainer[i], numCount, row, 1, 1)
			numCount++
			small.ShowAll()
//			fmt.Println("Num in row", numCount)
		}else {
//			small.InsertRow(row)
			small.Attach(buttonContainer[i], numCount, row, 1, 1)
			row++
			numCount = 0
//			fmt.Println("row", row)
			small.ShowAll()
		}
	}
	small.ShowAll()

}
func SetupBroadcastWindow(twoBuilder *gtk.Builder) {
	inspectUn, err := twoBuilder.GetObject("inspect")
	if err != nil {
		panic(err)
	}
	inspect := inspectUn.(*gtk.Box)
	button, err := gtk.ButtonNew()
	if err != nil {
		panic(err)
	}
	newBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		panic(err)
	}
	newLabel, err := gtk.LabelNew("doot")
	if err != nil {
		panic(err)
	}
	newLabel.SetText("BOOPS")
	newBox.Add(button)
	newBox.Add(newLabel)
	boxCtx, err := newBox.GetStyleContext()
	if err != nil {
		panic(err)
	}
	boxCtx.AddClass("cel")
	newBox.PackEnd(button, true, true, 1)
	inspect.Add(newBox)

}


func assembleBroadButton(name string) *gtk.Button {
	newBroadcast, err := gtk.ButtonNew()
	if err != nil {
		panic(err)
	}

	newBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		panic(err)
	}

	timeDateLabel, err := gtk.LabelNew(name+"timedate")
	if err != nil {
		panic(err)
	}

	messageLabel, err := gtk.LabelNew(name+"message")
	if err != nil {
		panic(err)
	}

	fromFieldLabel, err := gtk.LabelNew(name+"field")
	if err != nil {
		panic(err)
	}
	newBox.PackEnd(fromFieldLabel, false, false, 1)

	buttStyle, err := newBroadcast.GetStyleContext()
	if err != nil {
		panic(err)
	}
	buttStyle.AddClass("cel")
	buttStyle.AddClass("cell:hover")

	TDStyle, err := timeDateLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	TDStyle.AddClass("header")

	messStyle, err := messageLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	messStyle.AddClass("contents")

	fromFieldStyle, err := fromFieldLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	fromFieldStyle.AddClass("footer")

	newBox.Add(timeDateLabel)
	newBox.Add(messageLabel)
	newBox.Add(fromFieldLabel)

	newBroadcast.Add(newBox)

	return newBroadcast

}
func GetSender(message string) string {
	sender := strings.Split(message, "::SENDER::")[1]
	return sender

}
func assembleBroadButtonWithMessage(name string, message string, twoBuilder *gtk.Builder) *gtk.Button {
	newBroadcast, err := gtk.ButtonNew()
	if err != nil {
		panic(err)
	}

	newBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		panic(err)
	}

	fromLabel, err := gtk.LabelNew(name+"from")
	if err != nil {
		panic(err)
	}
	sender := GetSender(message)
	fromLabel.SetText("@"+sender)

	messageLabel, err := gtk.LabelNew(name+"message")
	if err != nil {
		panic(err)
	}
	mess := strings.Split(message, "::=")[1]
	messageLabel.SetText(mess)

	fromFieldLabel, err := gtk.LabelNew(name+"field")
	if err != nil {
		panic(err)
	}
	fromFieldLabel.SetText(time.Now().Weekday().String())
	newBox.PackEnd(fromFieldLabel, false, false, 1)

	buttStyle, err := newBroadcast.GetStyleContext()
	if err != nil {
		panic(err)
	}
	buttStyle.AddClass("cel")
	buttStyle.AddClass("cell:hover")

	TDStyle, err := fromLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	TDStyle.AddClass("header")

	messStyle, err := messageLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	messStyle.AddClass("contents")

	fromFieldStyle, err := fromFieldLabel.GetStyleContext()
	if err != nil {
		panic(err)
	}
	fromFieldStyle.AddClass("footer")

	newBox.Add(fromLabel)
	newBox.Add(messageLabel)
	newBox.Add(fromFieldLabel)

	newBroadcast.Add(newBox)


	newBroadcast.Connect("clicked", func (button *gtk.Button) {
		//fmt.Println("GETTING LABEL")
		inspectUn, err := twoBuilder.GetObject("inspectMess")
		if err != nil {
			panic(err)
		}
		inspect := inspectUn.(*gtk.Label)
		inspectWhoUn, err := twoBuilder.GetObject("inspectWho")
		if err != nil {
			panic(err)
		}
		inspectWho := inspectWhoUn.(*gtk.Label)

		inspectTimeUn, err := twoBuilder.GetObject("inspectTime")
		if err != nil {
			panic(err)
		}
		inspectTime := inspectTimeUn.(*gtk.Label)

		hours := strconv.Itoa(time.Now().Hour())
		minutes := strconv.Itoa(time.Now().Minute())

		inspectTime.SetText(time.Now().Weekday().String()+hours+minutes)

		inspectWho.SetText("@"+sender)

		inspect.SetText(mess)
		inTctx, err := inspectTime.GetStyleContext()
		if err != nil {
			panic(err)
		}
		inWctx, err := inspectWho.GetStyleContext()
		if err != nil {
			panic(err)
		}
		inctx, err := inspect.GetStyleContext()
		if err != nil {
			panic(err)
		}
		inTctx.AddClass("inspectIn")
		inWctx.AddClass("inspectIn")
		inctx.AddClass("inspectIn")
	})
	return newBroadcast

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
				launch(InitPlayer(user, pass), application, twoBuilder)
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

	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	// Set ApplicationWindow Properties
        wind.Show()
	application.AddWindow(wind)
    })
    var placeholder []string
    // Run Gtk application
    application.Run(placeholder)
}
