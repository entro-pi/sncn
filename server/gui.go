package main

import (
    "strconv"
    "strings"
    "log"
    "os"
    "fmt"
    "io/ioutil"
//    "bytes"
    "github.com/go-yaml/yaml"
//    "encoding/json"
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
	createRoomPopupUn, err := twoBuilder.GetObject("createRoom")
	if err != nil {
		panic(err)
	}
	createRoomPopup := createRoomPopupUn.(*gtk.Window)
	createRoomUn, err := twoBuilder.GetObject("createRoomButton")
	if err != nil {
		panic(err)
	}
	createRoom := createRoomUn.(*gtk.Button)
	createRoom.Connect("clicked", func() {
		createRoomPopup.Show()
		fmt.Println("create room button clicked")
	})
	createRoomCreateUn, err := twoBuilder.GetObject("createRoomCreate")
	if err != nil {
		panic(err)
	}
	var roomYaml []byte
	createRoomCreate := createRoomCreateUn.(*gtk.Button)
	createRoomCreate.Connect("clicked", func() {
		//create the yaml file here
		inspectUn, err := twoBuilder.GetObject("inspectMess")
		if err != nil {
			panic(err)
		}
		inspect := inspectUn.(*gtk.Label)
		fmt.Println("\033[38:2:0:200:0mCREATED ROOM\033[0m")
		//Change this so it will change the contents of
		//the popup when clicked
		inspect.SetText("ROOM CREATED")
		applyVnumUn, err := twoBuilder.GetObject("applyVnum")
		if err != nil {
			panic(err)
		}
		applyVnum := applyVnumUn.(*gtk.CheckButton)
		if applyVnum.GetActive() {
			fmt.Println("Make vnum yaml")
		}
		applyDescUn, err := twoBuilder.GetObject("applyDesc")
		if err != nil {
			panic(err)
		}
		applyDesc := applyDescUn.(*gtk.CheckButton)
		if applyDesc.GetActive() {
			fmt.Println("Make desc yaml")
		}
		applyExitUn, err := twoBuilder.GetObject("applyExit")
		if err != nil {
			panic(err)
		}
		applyExit := applyExitUn.(*gtk.CheckButton)
		if applyExit.GetActive() {
			fmt.Println("Make exit yaml")
		}
		applyExitSpecUn, err := twoBuilder.GetObject("applySpecExit")
		if err != nil {
			panic(err)
		}
		applyExitSpec := applyExitSpecUn.(*gtk.CheckButton)
		if applyExitSpec.GetActive() {
			fmt.Println("Make special exit yaml")
		}

	})
	exitCreateRoomUn, err := twoBuilder.GetObject("exitCreateRoom")
	if err != nil {
		panic(err)
	}
	exitCreateRoom := exitCreateRoomUn.(*gtk.Button)
	exitCreateRoom.Connect("clicked", func() {
		createRoomPopup.Close()
	})


	createPlayerUn, err := twoBuilder.GetObject("createPlayerButton")
	if err != nil {
		panic(err)
	}
	createPopupUn, err := twoBuilder.GetObject("createPlayer")
	if err != nil {
		panic(err)
	}
	createPopup := createPopupUn.(*gtk.Window)

	createPlayer := createPlayerUn.(*gtk.Button)
	createPlayer.Connect("clicked", func() {
		createPopup.Show()
	})
	okCreatePlayerUn, err := twoBuilder.GetObject("okCreatePlayer")
	if err != nil {
		panic(err)
	}
	okCreatePlayer := okCreatePlayerUn.(*gtk.Button)
	okCreatePlayer.Connect("clicked", func () {
		nameUn, err := twoBuilder.GetObject("createPlayerName")
		if err != nil {
			panic(err)
		}
		name := nameUn.(*gtk.Entry)
		passUn, err := twoBuilder.GetObject("createPlayerPass")
		if err != nil {
			panic(err)
		}
		pass := passUn.(*gtk.Entry)
		Name, err := name.GetText()
		Pass, err := pass.GetText()
		fmt.Println("CREATE A NEW USER : ",Name," :: PASS :: ",Pass)
		Name = strings.ToUpper(Name)
		inspectUn, err := twoBuilder.GetObject("inspectMess")
		if err != nil {
			panic(err)
		}
		inspect := inspectUn.(*gtk.Label)
		if _, err := os.Stat("../pot/pfiles/"+Name+".yaml"); err == nil {
			fmt.Println("\033[38:2:200:0:0mPLAYERFILE EXISTS\033[0m")
			//Change this so it will change the contents of
			//the popup when clicked
			inspect.SetText("ERROR\nPLAYERFILE EXISTS")
			//createPopup.SetText("Playerfile exists!")
		}else {
			playNew := InitPlayer(Name, Pass)
			yamlPlay, err := yaml.Marshal(playNew)
			if err != nil {
				panic(err)
			}
			file, err := os.OpenFile("../pot/pfiles/"+Name+".yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(yamlPlay)
			readYamlFile(Name)
			createPopup.Close()
			inspect.SetText("Playerfile "+Name+" created")
			//get the values of the player and do the thing
		}
	})

	cancelCreatePlayerUn, err := twoBuilder.GetObject("cancelCreatePlayer")
	if err != nil {
		panic(err)
	}
	cancelCreatePlayer := cancelCreatePlayerUn.(*gtk.Button)
	cancelCreatePlayer.Connect("clicked", func () {
		//cancel the thingy
		createPopup.Close()
	})


	sendUn, err := twoBuilder.GetObject("Send")
	if err != nil {
		panic(err)
	}

	send := sendUn.(*gtk.Button)

	send.Connect("pressed", func() {
		postingUn, err := twoBuilder.GetObject("postingWin")
		if err != nil {
			panic(err)
		}
		posting := postingUn.(*gtk.ScrolledWindow)
		spinUn, err := twoBuilder.GetObject("spin")
		if err != nil {
			panic(err)
		}
		spinner := spinUn.(*gtk.Spinner)
	//	spinner.SetVisible(true)
		spinner.Start()
		posting.ShowAll()
	})

	send.Connect("clicked", func() {
		smallUn, err := twoBuilder.GetObject("smalltalkWin")
		if err != nil {
			panic(err)
		}
		small := smallUn.(*gtk.ScrolledWindow)
		inputUn, err := twoBuilder.GetObject("postBuf")
		if err != nil {
			panic(err)
		}
		input := inputUn.(*gtk.TextBuffer)
		start, end := input.GetBounds()
		inputText, err := input.GetText(start, end, false)
		if err != nil {
			panic(err)
		}
		tellBool := false
		inputText = strings.ReplaceAll(inputText, "\n", "")
	        tellToArray := strings.Split(inputText, "@")
	        if len(tellToArray) > 1 {
	                tellBool = true
	        }
		postingUn, err := twoBuilder.GetObject("postingWin")
		if err != nil {
			panic(err)
		}
		posting := postingUn.(*gtk.ScrolledWindow)
		spinUn, err := twoBuilder.GetObject("spin")
		if err != nil {
			panic(err)
		}
		spinner := spinUn.(*gtk.Spinner)

		go func() {
			doGUIInput(play, inputText)
			fill(play, twoBuilder, tellBool)
			small.ShowAll()
			input.SetText("")
			spinner.Stop()
		//	spinner.SetVisible(false)
			posting.ShowAll()
		}()
	})
	invUn, err := twoBuilder.GetObject("listTells")
	if err != nil {
		panic(err)
	}
	inv := invUn.(*gtk.Button)
	inv.Connect("clicked", func (button *gtk.Button) {
		boxUn, err := twoBuilder.GetObject("smalltalkWin")
		if err != nil {
			panic(err)
		}
		box := boxUn.(*gtk.ScrolledWindow)
		//if box.GetVisible() {
			box.SetVisible(false)
		//}else {
		//	box.SetVisible(true)
		//}
	})
	equipUn, err := twoBuilder.GetObject("broadMain1")
	if err != nil {
		panic(err)
	}
	equip := equipUn.(*gtk.Button)
	equip.Connect("clicked", func () {
		box1Un, err := twoBuilder.GetObject("smalltalkWin")
		if err != nil {
			panic(err)
		}
		box1 := box1Un.(*gtk.ScrolledWindow)
		if box1.GetVisible() {
			box1.SetVisible(false)
		}
		box1.ShowAll()
	})
	wind := appWindow.(*gtk.ApplicationWindow)
	wind.Fullscreen()
	wind.SetResizable(false)
	wind.SetPosition(gtk.WIN_POS_CENTER)
	tellsUn, err := twoBuilder.GetObject("tellsMain")
	if err != nil {
		panic(err)
	}
	tells := tellsUn.(*gtk.Button)
	tells.Connect("clicked", func () {
		fill(play, twoBuilder, true)
		smallUn, err := twoBuilder.GetObject("smalltalkWin")
		if err != nil {
			panic(err)
		}
		small := smallUn.(*gtk.ScrolledWindow)
		small.ShowAll()
		eqGridUn, err := twoBuilder.GetObject("equipmentWin")
		if err != nil {
			panic(err)
		}
		eqGrid := eqGridUn.(*gtk.ScrolledWindow)
		if eqGrid.GetVisible() {
			eqGrid.SetVisible(false)
		}
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
		eqGridUn, err := twoBuilder.GetObject("equipmentWin")
		if err != nil {
			panic(err)
		}
		eqGrid := eqGridUn.(*gtk.ScrolledWindow)
		if eqGrid.GetVisible() {
			eqGrid.SetVisible(false)
		}else {
			eqGrid.SetVisible(true)
			fillEq(play, twoBuilder)
			eqGrid.ShowAll()
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
		fill(play, twoBuilder, false)
		smallUn, err := twoBuilder.GetObject("smalltalkWin")
		if err != nil {
			panic(err)
		}
		eqGridUn, err := twoBuilder.GetObject("equipmentWin")
		if err != nil {
			panic(err)
		}
		eqGrid := eqGridUn.(*gtk.ScrolledWindow)
		if eqGrid.GetVisible() {
			eqGrid.SetVisible(false)
		}
		small := smallUn.(*gtk.ScrolledWindow)
		small.ShowAll()
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
	disp, err := screen.GetDisplay()
	if err != nil {
		panic(err)
	}
	windowUn, err := twoBuilder.GetObject("mainwindow")
	if err != nil {
		panic(err)
	}
	windowApp := windowUn.(*gtk.ApplicationWindow)
	window, err := windowApp.GetWindow()
	if err != nil {
		panic(err)
	}
	moni, err := disp.GetMonitorAtWindow(window)
	if err != nil {
		panic(err)
	}
	fillTree(twoBuilder)
	fillList(twoBuilder)
	geo := moni.GetGeometry()
	height := geo.GetHeight()
	width := geo.GetWidth()
	wind.SetDefaultSize(width, height)
	wind.Fullscreen()
        wind.Show()
	application.AddWindow(wind)

}

const (
	COLUMN_SLOT = 0
	COLUMN_NAME = 1
	COLUMN_ITEM = 2
	COLUMN_VALUE = 3
	COLUMN_LONGNAME = 4
	COLUMN_NUMBER = 5
)

func readYamlFile(Name string) map[interface{}]interface{} {
	file, err := os.OpenFile("../pot/pfiles/"+Name+".yaml", os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	playList := make(map[interface{}]interface{})
//	playList := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), playList)
	if err != nil {
		panic(err)
	}
	if playList["name"] != nil {
		fmt.Println(playList["name"])
	}
	return playList
}
func createColumn(twee *gtk.TreeView, val string, constant int) *gtk.TreeViewColumn {
	var renderer *gtk.CellRenderer

	col, err := gtk.TreeViewColumnNew()
	if err != nil {
		panic(err)
	}
	col.SetTitle(val)
	col.AddAttribute(renderer, val, constant)
	col.SetVisible(true)
	return col

}
func createColumnPackStart(twee *gtk.TreeView, val string, value string, constant int) (*gtk.TreeViewColumn) {

	col, err := gtk.TreeViewColumnNew()
	if err != nil {
		panic(err)
	}
/*	renderer, err := gtk.CellRendererTextNew()
	if err != nil {
		panic(err)
	}*/
//	col.PackStart(renderer, true)
//	renderer.Set("visible", true)
//	renderer.Set("text", value)
	col.SetTitle(val)
//	col.AddAttribute(renderer, col.GetTitle(), constant)
	col.SetVisible(true)
	return col

}
func labelColumns(twee *gtk.TreeView, value string, constant int, col *gtk.TreeViewColumn) (*gtk.TreeViewColumn) {

	renderer, err := gtk.CellRendererTextNew()
	if err != nil {
		panic(err)
	}
	col.PackStart(renderer, true)
	renderer.Set("visible", true)
//	renderer.Set("text", value)
	
	col.AddAttribute(renderer, "text", constant)
	col.SetVisible(true)
	return col

}

func fillList(twoBuilder *gtk.Builder) {

	tweeUn, err := twoBuilder.GetObject("twee")
	if err != nil {
		panic(err)
	}
	twee := tweeUn.(*gtk.TreeView)
	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_FLOAT, glib.TYPE_STRING, glib.TYPE_INT)
	if err != nil {
		panic(err)
	}

/*	listStoreUn, err := twoBuilder.GetObject("liststore1")
	if err != nil {
		panic(err)
	}
	listStore, err := gtk.ListStoreNew()
	if err != nil {
		panic(err)
	}*/
	zeroColumn := createColumnPackStart(twee, "Slot", "", COLUMN_SLOT)
	twee.AppendColumn(zeroColumn)
	firstColumn := createColumnPackStart(twee, "Name", "Nyancat", COLUMN_NAME)
	twee.AppendColumn(firstColumn)
	secondColumn := createColumnPackStart(twee, "Item", "4000", COLUMN_ITEM)
	twee.AppendColumn(secondColumn)
	thirdColumn := createColumnPackStart(twee, "Value", "1.0", COLUMN_VALUE)
	twee.AppendColumn(thirdColumn)
	fourthColumn := createColumnPackStart(twee, "LongName", "A poptart kitten nyans along happily", COLUMN_LONGNAME)
	twee.AppendColumn(fourthColumn)
	fifthColumn := createColumnPackStart(twee, "Number", "0", COLUMN_NUMBER)
	twee.AppendColumn(fifthColumn)
	pos := listStore.Append()
	labelColumns(twee, "Slot", COLUMN_SLOT, zeroColumn)
	labelColumns(twee, "Rose", COLUMN_NAME, firstColumn)
	labelColumns(twee, "4001", COLUMN_ITEM, secondColumn)
	labelColumns(twee, "5.0", COLUMN_VALUE, thirdColumn)
	labelColumns(twee, "A wilting red rose.", COLUMN_LONGNAME, fourthColumn)
	labelColumns(twee, "0", COLUMN_NUMBER, fifthColumn)
	err = listStore.SetValue(pos, 0, "Head")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 1, "nyancat")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 2, 4000)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 3, 1.0)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 4, "nyaaaaaaacat")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 5, 1)
	if err != nil {
		panic(err)
	}
	pos = listStore.Append()
	err = listStore.SetValue(pos, 0, "Face")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 1, "rose")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 2, 4001)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 3, 50.0)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 4, "A wilting red rose")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, 5, 1)
	if err != nil {
		panic(err)
	}
	listStore.Append()



	twee.SetModel(listStore)
	twee.SetReorderable(false)
	twee.SetVisible(true)
	twee.Show()
}
func fillTree(twoBuilder *gtk.Builder) {

	tweeUn, err := twoBuilder.GetObject("twee1")
	if err != nil {
		panic(err)
	}
	twee := tweeUn.(*gtk.TreeView)
	listStore, err := gtk.TreeStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_FLOAT, glib.TYPE_STRING, glib.TYPE_INT)
	if err != nil {
		panic(err)
	}

/*	listStoreUn, err := twoBuilder.GetObject("liststore1")
	if err != nil {
		panic(err)
	}
	listStore, err := gtk.ListStoreNew()
	if err != nil {
		panic(err)
	}*/
	zeroColumn := createColumnPackStart(twee, "", "", COLUMN_SLOT)
	twee.AppendColumn(zeroColumn)
	firstColumn := createColumnPackStart(twee, "Name", "Nyancat", COLUMN_NAME)
	twee.AppendColumn(firstColumn)
	secondColumn := createColumnPackStart(twee, "Item", "4000", COLUMN_ITEM)
	twee.AppendColumn(secondColumn)
	thirdColumn := createColumnPackStart(twee, "Value", "1.0", COLUMN_VALUE)
	twee.AppendColumn(thirdColumn)
	fourthColumn := createColumnPackStart(twee, "LongName", "A poptart kitten nyans along happily", COLUMN_LONGNAME)
	twee.AppendColumn(fourthColumn)
	fifthColumn := createColumnPackStart(twee, "Number", "1", COLUMN_NUMBER)
	twee.AppendColumn(fifthColumn)
	top := listStore.Append(nil)
	labelColumns(twee, "", COLUMN_SLOT, zeroColumn)
	labelColumns(twee, "Rose", COLUMN_NAME, firstColumn)
	labelColumns(twee, "4001", COLUMN_ITEM, secondColumn)
	labelColumns(twee, "5.0", COLUMN_VALUE, thirdColumn)
	labelColumns(twee, "A wilting red rose.", COLUMN_LONGNAME, fourthColumn)
	labelColumns(twee, "1", COLUMN_NUMBER, fifthColumn)
	err = listStore.SetValue(top, COLUMN_NAME, "portable hole")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(top, COLUMN_ITEM, 4002)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(top, COLUMN_VALUE, 500.0)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(top, COLUMN_LONGNAME, "An atypical pocket of spacetime.")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(top, COLUMN_NUMBER, 1)
	if err != nil {
		panic(err)
	}
	pos := listStore.Insert(top, 0)
	err = listStore.SetValue(pos, COLUMN_NAME, "nyancat")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_ITEM, 4000)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_VALUE, 1.0)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_LONGNAME, "nyaaaaaaacat")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_NUMBER, 1)
	if err != nil {
		panic(err)
	}
	pos = listStore.Insert(top, 0)
	err = listStore.SetValue(pos, COLUMN_NAME, "rose")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_ITEM, 4001)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_VALUE, 50.0)
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_LONGNAME, "A wilting red rose")
	if err != nil {
		panic(err)
	}
	err = listStore.SetValue(pos, COLUMN_NUMBER, 1)
	if err != nil {
		panic(err)
	}

//	listStore.Insert(top, 0)



	twee.SetModel(listStore)
	twee.SetReorderable(false)
	twee.SetVisible(true)
	twee.Show()
}

func fill(play Player, twoBuilder *gtk.Builder, tellorbroad bool)  {
	var broadcastContainer []string
	var buttonContainer []*gtk.Button
	if tellorbroad {
//		play := InitPlayer("WEASEL", "lol")
		broadcastContainer = drawPlainTells(play)
	}else {
//		play := InitPlayer("WEASEL", "lol")
		broadcastContainer = drawPlainBroadcasts(play)
	}
	if len(broadcastContainer) >= 40 {
		broadcastContainer = broadcastContainer[len(broadcastContainer)-40:len(broadcastContainer)]
	}
	for i := 0;i < len(broadcastContainer);i++ {
		fmt.Println(i)
		broad := assembleBroadButtonWithMessage(strconv.Itoa(i), broadcastContainer[i], twoBuilder)
		buttonContainer = append(buttonContainer, broad)
	}

	smallUn, err := twoBuilder.GetObject("smalltalkGrid")
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
	small.SetRowHomogeneous(true)
	small.SetColumnHomogeneous(true)
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
	fromLabel.SetText("<-"+sender)

	messageLabel, err := gtk.LabelNew(name+"message")
	if err != nil {
		panic(err)
	}
	mess := strings.Split(message, "::=")[1]
	messHolder := ""
	addNewLine := false
	since := 0
	count := 0
	for i := 0;i < len(mess);i++ {
		count++
		if count == 12 {
			addNewLine = true
		}
		if addNewLine && mess[i] == ' ' {
			messHolder += string(mess[i])+"\n"
			addNewLine = false
			count = 0
		}else if addNewLine && mess[i] != ' ' {
			since++
		}else if since == 5 {
			//since we haven't gotten a space in five
			//characters, break the line anyway
			messHolder += string(mess[i])+"\n"
			addNewLine = false
			since = 0
			count = 0
		}else {
			messHolder += string(mess[i])
		}
	}
	fmt.Print(messHolder)
	messageLabel.SetText(messHolder)

	fromFieldLabel, err := gtk.LabelNew(name+"field")
	if err != nil {
		panic(err)
	}
	timeStamp := strings.Split(message, "::TIMESTAMP::")[1]
	fromFieldLabel.SetText(timeStamp)
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
		mess := strings.Split(message, "::=")[1]
		messHolder := ""
		addNewLine := false
		since := 0
		count := 0
		for i := 0;i < len(mess);i++ {
			count++
			if count == 40 {
				addNewLine = true
			}
			if addNewLine && mess[i] == ' ' {
				messHolder += string(mess[i])+"\n"
				addNewLine = false
				count = 0
			}else if addNewLine && mess[i] != ' ' {
				messHolder += string(mess[i])
				since++
			}else if since == 5 {
				//since we haven't gotten a space in five
				//characters, break the line anyway
				messHolder += string(mess[i])+"\n"
				addNewLine = false
				since = 0
				count = 0
			}else {
				messHolder += string(mess[i])
			}
		}
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

		inspectTime.SetText(timeStamp)

		inspectWho.SetText("<-"+sender)

		inspect.SetText(messHolder)
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


func LaunchGUI(fileChange chan bool) {
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
	    twoBuilder, err := gtk.BuilderNewFromFile("server.glade")
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
			if len(userCaps) > 3 && len(pass) > 3 {
        		        play := InitPlayer(user, pass)
				whoList := who(play.Name)
	                	go func() { actOn(play, fileChange, whoList)}()
				launch(play, application, twoBuilder)
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