package main

import (
    "time"
    "strconv"
    "strings"
    "log"
    "os"
    "fmt"
    "math"
    "io/ioutil"
    "github.com/go-yaml/yaml"
    "unsafe"
//    "github.com/gotk3/gotk3/cairo"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/go-gl/gl/v4.1-core/gl"
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

func register(application *gtk.Application, twoBuilder *gtk.Builder) bool {
	registerWindowUn, err := twoBuilder.GetObject("createWindow")
	if err != nil {
		panic(err)
	}
	registerWindow := registerWindowUn.(*gtk.Window)
//	registerWindow.Fullscreen()
	registerWindow.ShowAll()
	registerWindow.SetDefaultSize(400, 200)
	registerWindow.Fullscreen()
	registerWindow.GrabFocus()
        registerWindow.Show()
	exitUn, err := twoBuilder.GetObject("exitCreate")
	if err != nil {
		panic(err)
	}
	created := false
	didAThing := false
	exit := exitUn.(*gtk.Button)
	exit.Connect("clicked", func() {
		created = false
		didAThing = true
	})
	createUn, err := twoBuilder.GetObject("create")
	if err != nil {
		panic(err)
	}
	create := createUn.(*gtk.Button)
	create.Connect("clicked", func() {
		created = true
		didAThing = true
	})

//	fmt.Println("Register launched")
//	start := time.Now()
/*	for {
		if didAThing {
			break
		}
		registerWindow.Show()
//		registerWindow.GrabFocus()
	}*/
	return created
}

func launch(play Player, application *gtk.Application, twoBuilder *gtk.Builder) {

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
	//To populate the text window programmatically we do the following
	logMainUn, err := twoBuilder.GetObject("mainLog")
	if err != nil {
		panic(err)
	}
	logMain := logMainUn.(*gtk.Label)
	/*buildString := ""
	for i := 0;i < 45;i++ {
		buildString += strconv.Itoa(i)+"\n"
	}
	logMain.SetText(buildString)
*/	
	
	logMain.SetText("Connection interrupted.\nLock destination and engage to view.")
	logMain.AddTickCallback(round, uintptr(unsafe.Pointer(&play)))

	zoomInUn, err := twoBuilder.GetObject("zoomIn")
	if err != nil {
		panic(err)
	}
	zoomIn := zoomInUn.(*gtk.Button)
	zoomOutUn, err := twoBuilder.GetObject("zoomOut")
	if err != nil {
		panic(err)
	}
	zoomOut := zoomOutUn.(*gtk.Button)
	
	mapBoxUn, err := twoBuilder.GetObject("mapBox")
	if err != nil {
		panic(err)
	}
	numCol := 7
	numRow := 10//TODO
	//rooms := walkRooms
	boxMap := make([]*gtk.Button, numCol*numRow)
	mapBox := mapBoxUn.(*gtk.Box)
	mapGrid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}
	mapBox.Add(mapGrid)
	col := 1
	row := 1
	for i := 0;i < numCol * numRow;i++ {
		roomButton, err := gtk.ButtonNewWithLabel("0000")
		if err != nil {
			panic(err)
		}
		roomButton.SetHExpand(true)
		roomButton.SetVExpand(true)
		boxMap[i] = roomButton
		mapGrid.Attach(boxMap[i], col, row, 1, 1)
		if col == numCol {
			col = 1
			row++
		}else {
			col++
		}
	}
	zoomIn.Connect("clicked", func() {

	})
	zoomOut.Connect("clicked", func(button *gtk.Button) {
	})
	rootClicked := false
	rootPos := 0
//	world := populateWorld()
	//initiate the world by walking all rooms we currently have access to
	mappedWorld := walkRooms(play.CurrentRoom)
	//now loop over the in place map
	for i := 0;i < len(boxMap)-2;i++ {
		if i == (len(boxMap) / 2) {
			rootPos = i-4
			boxMap[i-4].SetLabel("ROOT")
			boxStyle, err := boxMap[i-4].GetStyleContext()
			if err != nil {
				panic(err)
			}
			boxStyle.AddClass("mapgOnBl")
/*		}
		lab, err := boxMap[rootPos].GetLabel()
		if err != nil {
			panic(err)
		}
		if lab == "ROOT" {
		*/	boxMap[rootPos].SetLabel(mappedWorld[play.CurrentRoom.Vnums].Vnums)
			if play.CurrentRoom.ExitMap["North"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["North"])
				boxMap[rootPos-numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos-numCol].GetStyleContext()
				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")
			}
			if play.CurrentRoom.ExitMap["South"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["South"])
				boxMap[rootPos+numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos+numCol].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["East"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["East"])
				boxMap[rootPos+1].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos+1].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["West"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["West"])
				boxMap[rootPos-1].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos-1].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["NorthEast"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["NorthEast"])
				boxMap[rootPos-1-numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos-1-numCol].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["NorthWest"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["NorthWest"])
				boxMap[rootPos+1-numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos+1-numCol].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["SouthEast"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["SouthEast"])
				boxMap[rootPos+1+numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos+1+numCol].GetStyleContext()

				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
			if play.CurrentRoom.ExitMap["SouthWest"] > 1 {
				stringRoom := strconv.Itoa(play.CurrentRoom.ExitMap["SouthWest"])
				boxMap[rootPos+1-numCol].SetLabel(stringRoom)
				styCtx, err := boxMap[rootPos+1-numCol].GetStyleContext()
				if err != nil {
					panic(err)
				}
				styCtx.AddClass("mapgOnBl")

			}
		}

		if i > 1 && i < len(boxMap) {
			boxMap[i].Connect("clicked", func(butt *gtk.Button) {
				value, err := butt.GetLabel()
				if err != nil {
					panic(err)
				}
				for c := 0;c < len(boxMap);c++ {
					tocompare, err := boxMap[c].GetLabel()
					if err != nil {
						panic(err)
					}
					if tocompare == value && (tocompare != "0000" && value != "0000") {
						fmt.Println("tocompare:"+tocompare)
						roomDesc := mappedWorld[value].Desc
						inspectUn, err := twoBuilder.GetObject("inspectMess")
						if err != nil {
							panic(err)
						}
						inspect := inspectUn.(*gtk.Label)
						fmt.Println("value:"+value)
						ctx, err := boxMap[c].GetStyleContext()
						if err != nil {
							panic(err)
						}

						if !rootClicked {
							inspect.SetText(roomDesc+"\n\n\n\nassign(Destination "+value+")")

							ctx.AddClass("mapButton")
							rootClicked = true
						}else {
							inspect.SetText("\n\n\n\nmalloc(Destination)")
							rootClicked = false
							for d := 0;d < len(boxMap);d++ {
								newCtx, err := boxMap[d].GetStyleContext()
								if err != nil {
									panic(err)
								}
								newCtx.RemoveClass("mapButton")
							}
							gtk.MainIterationDo(true)
						}
					}
				}
			})
                        boxMap[i].Connect("button-press-event", func (butt *gtk.Button, ev *gdk.Event) {
                                keyEvent := gdk.EventButtonNewFromEvent(ev)
				mapRightUn, err := twoBuilder.GetObject("mapRight")
				if err != nil {
					panic(err)
				}
				mapRight := mapRightUn.(*gtk.Popover)
				if keyEvent.ButtonVal() == 1 {
                                        val, err := butt.GetLabel()
                                        if err != nil {
                                                panic(err)
                                        }
					mapRight.SetVisible(false)
                                        fmt.Println("Left click on : "+ val)
                                }
                                if keyEvent.ButtonVal() == 2 {
                                        val, err := butt.GetLabel()
                                        if err != nil {
                                                panic(err)
                                        }
                                        fmt.Println("Middle click on : "+ val)
					mapRight.SetVisible(false)
                                }
                                if keyEvent.ButtonVal() == 3 {
                                        val, err := butt.GetLabel()
                                        if err != nil {
                                                panic(err)
                                        }
                                        fmt.Println("Right click on : "+ val)
					if val != "0000" {
						mapRight.SetRelativeTo(butt)
						mapRight.SetVisible(true)
						mapRight.Show()
					}else {
						mapRight.SetVisible(false)
					}
					engageUn, err := twoBuilder.GetObject("mapRightEngage")
					if err != nil {
						panic(err)
					}
					engage := engageUn.(*gtk.Button)
					engage.Connect("clicked", func(button *gtk.Button) {
						if val != "0000" {
							play.CurrentRoom = mappedWorld[val]
						}
					})

                                }
                        })

		}
	}
	mapBox.ShowAll()


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
	inputMainUn, err := twoBuilder.GetObject("inputMain")
	if err != nil {
		panic(err)
	}
	inputMain := inputMainUn.(*gtk.Entry)
	inputMain.Connect("activate", func() {
		input, err := inputMain.GetText()
		if err != nil {
			panic(err)
		}
		doGUIInput(play, input)
		inputMain.SetText("")
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
		input := inputUn.(*gtk.Entry)
		inputText, err := input.GetText()
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
	invUn, err := twoBuilder.GetObject("invMain")
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
	//now we test the prompt

	promptUn, err := twoBuilder.GetObject("prompt")
	if err != nil {
		panic(err)
	}
	prompt := promptUn.(*gtk.Box)
	//Populate the prompt with the player's hp and mana etc...
	hpText, err := gtk.ButtonNewWithLabel(strconv.Itoa(play.Rezz))
	techText, err := gtk.ButtonNewWithLabel(strconv.Itoa(play.Tech))
	manaText, err := gtk.ButtonNewWithLabel(strconv.Itoa(play.Mana))
	if err != nil {
		panic(err)
	}
	//Now set their styles
	hpCtx, err := hpText.GetStyleContext()
	if err != nil {
		panic(err)
	}
	techCtx, err := techText.GetStyleContext()
	if err != nil {
		panic(err)
	}
	manaCtx, err := manaText.GetStyleContext()
	if err != nil {
		panic(err)
	}
	hpCtx.AddClass("button")
	techCtx.AddClass("button")
	manaCtx.AddClass("button")
	prompt.Add(hpText)
	prompt.Add(techText)
	prompt.Add(manaText)
	prompt.ShowAll()
	play.Rezz -= 1
	qAct2Un, err := twoBuilder.GetObject("statsUp")
	if err != nil {
		panic(err)
	}
	qAct2 := qAct2Un.(*gtk.Button)
	qAct2.Connect("pressed", func () {
		play.Rezz += 10
		play.Mana += 10
		play.Tech += 10
		hpText.SetLabel(strconv.Itoa(play.Rezz))
		techText.SetLabel(strconv.Itoa(play.Tech))
		manaText.SetLabel(strconv.Itoa(play.Mana))
	})
	qAct1Un, err := twoBuilder.GetObject("statsDown")
	if err != nil {
		panic(err)
	}
	qAct1 := qAct1Un.(*gtk.Button)
	qAct1.Connect("pressed", func () {
		play.Rezz -= 10
		play.Mana -= 10
		play.Tech -= 10
		hpText.SetLabel(strconv.Itoa(play.Rezz))
		techText.SetLabel(strconv.Itoa(play.Tech))
		manaText.SetLabel(strconv.Itoa(play.Mana))
	})

	equipUn, err := twoBuilder.GetObject("equipMain")
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
	smallUn, err := twoBuilder.GetObject("smalltalkWin")
	if err != nil {
		panic(err)
	}
	small := smallUn.(*gtk.ScrolledWindow)
//	fill(play, twoBuilder, true)
	
	tells.Connect("clicked", func () {
		small.Show()
	})
	broadUn, err := twoBuilder.GetObject("broadMain")
	if err != nil {
		panic(err)
	}
	broad := broadUn.(*gtk.Button)
	go fill(play, twoBuilder, false)
	broad.Connect("clicked", func () {
		small.Show()
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

//	register(moni, application, twoBuilder)
	fillTree(twoBuilder)
	fillList(twoBuilder)
	geo := moni.GetGeometry()
	heightD := geo.GetHeight()
	widthD := geo.GetWidth()
	wind.SetDefaultSize(widthD, heightD)
	wind.Fullscreen()
	rezzGLUn, err := twoBuilder.GetObject("RezzGL")
	if err != nil {
		panic(err)
	}
	rezzGL := rezzGLUn.(*gtk.GLArea)
	fmt.Println(rezzGL)
	//startDelta := time.Now()
	rezzGL.AddTickCallback(render, uintptr(0))
	rezzGL.SetAutoRender(true)
	rezzGL.Connect("create-context", func (area *gtk.GLArea) {
		gl.Init()
		area.MakeCurrent()
		fmt.Println(area.GetError())
	})
	GLCount := 1
	rezzGL.Connect("render", func (area *gtk.GLArea)  {
		GLCount = renderRezz(play, GLCount)
	})
	techGLUn, err := twoBuilder.GetObject("TechGL")
	if err != nil {
		panic(err)
	}
	techGL := techGLUn.(*gtk.GLArea)
	techGL.AddTickCallback(render, uintptr(0))
	techGL.SetAutoRender(true)
	techGL.Connect("render", func (area *gtk.GLArea) {
		GLCount = renderTech(play, GLCount)
	})
	manaGLUn, err := twoBuilder.GetObject("ManaGL")
	if err != nil {
		panic(err)
	}
	manaGL := manaGLUn.(*gtk.GLArea)
	manaGL.AddTickCallback(render, uintptr(0))
	manaGL.SetAutoRender(true)
	manaGL.Connect("render", func (area *gtk.GLArea) {
		GLCount = renderMana(play, GLCount)
	})
	sideGLUn, err := twoBuilder.GetObject("sideGL")
	if err != nil {
		panic(err)
	}
	sideGL := sideGLUn.(*gtk.GLArea)
	sideGL.AddTickCallback(render, uintptr(0))
	sideGL.Connect("render", func (area *gtk.GLArea) {
		GLCount = renderMana(play, GLCount)
	})
        wind.Show()
	application.AddWindow(wind)

}


func render(widget *gtk.Widget, frameClock *gdk.FrameClock, Userdata uintptr) bool {
	widget.QueueDraw()
	return true
}
func round(widget *gtk.Widget, frameClock *gdk.FrameClock, Userdata uintptr) bool {
	log := (*gtk.Label)(unsafe.Pointer(widget))
	play := (*Player)(unsafe.Pointer(Userdata))
	log.SetText(play.CurrentRoom.Desc)
	return true
}
func renderRezz(play Player, count int) int {
	if count >= 60 {
		count = 1
	}
	delta := play.MaxRezz - play.Rezz
	deltaSin := (float64(delta) / 60.0)
	count++
	colorRed := float32(math.Cos(float64(deltaSin)) )
		gl.ClearColor(colorRed, 0, 0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	return count
}
func renderTech(play Player, count int) int {
	if count >= 60 {
		count = 1
	}
	delta := play.MaxTech - play.Tech
	deltaSin :=  (float64(delta) / 60.0)
	count++
	colorGreen := float32(math.Cos(float64(deltaSin)) )
		gl.ClearColor(0, colorGreen, 0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	return count
}
func renderMana(play Player, count int) int {
	if count >= 60 {
		count = 1
	}
	delta := play.MaxMana - play.Mana
	deltaSin :=  (float64(delta) / 60.0)
	count++
	colorPurp := float32(math.Cos(float64(deltaSin)) )
		gl.ClearColor(colorPurp, 0, colorPurp, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	return count
}


const (
	COLUMN_SLOT = 0
	COLUMN_NAME = 1
	COLUMN_ITEM = 2
	COLUMN_VALUE = 3
	COLUMN_LONGNAME = 4
	COLUMN_NUMBER = 5
)

func showRoom(path string) string {
	var space Space
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	roomContents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(roomContents, &space)
	if err != nil {
		panic(err)
	}
	return space.Desc

}


func initColumn(value string, constant int, col *gtk.TreeViewColumn) (*gtk.TreeViewColumn) {
	healthy, failing, failed := initIcons()
//	time.Sleep(10*time.Millisecond)
        renderer, err := gtk.CellRendererPixbufNew()
        if err != nil {
                panic(err)
        }
	col.PackStart(renderer, true)
        renderer.Set("visible", true)
//        renderer.Set("text", value)
	if value == "failing"  {
		renderer.Set("pixbuf", failing)
       	}else if value == "failed" {
		renderer.Set("pixbuf", failed)
	}else {
		renderer.Set("pixbuf", healthy)
	}

//	col.AddAttribute(renderer, "text", 0)
        col.SetVisible(true)
        return col

}

func initIcons() (gdk.Pixbuf, gdk.Pixbuf, gdk.Pixbuf) {
	healthyImg, err := gtk.ImageNewFromFile("dat/healthy.png")
	if err != nil {
		panic(err)
	}
	healthy := healthyImg.GetPixbuf()
	failingImg, err := gtk.ImageNewFromFile("dat/failing.png")
	if err != nil {
		panic(err)
	}
	failing := failingImg.GetPixbuf()
	failedImg, err := gtk.ImageNewFromFile("dat/failed.png")
	if err != nil {
		panic(err)
	}
	failed := failedImg.GetPixbuf()
	if err != nil {
		panic(err)
	}
	return *healthy, *failing, *failed
}

func backAndForthProgress(startHealth int, endHealth int, twoBuilder *gtk.Builder, pos *gtk.TreeIter) *gtk.TreeIter {
	healthBarUn, err := twoBuilder.GetObject("healthBar")
	if err != nil {
		panic(err)
	}
	healthBar := healthBarUn.(*gtk.TreeView)
	healthStore, err := gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		panic(err)
	}
	healthBar.SetModel(healthStore)
//	healthy := " "
//	failing := " "
//	failed := " "

//	bar := ""
	pos = healthStore.Append()
	for i := 0;i < 100 ;i++ {
		col := healthBar.GetColumn(i)
		healthBar.RemoveColumn(col)
	}
	healthStore.Clear()
	percent := math.Floor(float64(startHealth) / float64(endHealth))
	hundredPer := int(percent * 100)
	count := 0
	for i := count;i < hundredPer;i++ {
		time.Sleep(10*time.Millisecond)
		col, err := gtk.TreeViewColumnNew()
		if err != nil {
			panic(err)
		}
		initColumn("healthy", i, col)
		healthBar.AppendColumn(col)
//		healthStore.SetValue(pos, i, healthy)
//		pos = healthStore.Append()
		count++
	}
	col, err := gtk.TreeViewColumnNew()
	if err != nil {
		panic(err)
	}
	healthBar.AppendColumn(col)
	initColumn("failing", count, col)
//	healthStore.SetValue(pos, count, failing)
//	pos = healthStore.Append()
	count++
	for i := count;i < 100;i++ {
		time.Sleep(10*time.Millisecond)
		col, err := gtk.TreeViewColumnNew()
		if err != nil {
			panic(err)
		}
		healthBar.AppendColumn(col)
		initColumn("failed", i, col)
//		healthStore.SetValue(pos, i, failed)
//		pos = healthStore.Append()
		count++
	}
	pos = healthStore.Append()
	return pos
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
	if len(broadcastContainer) >= 6 {
		broadcastContainer = broadcastContainer[len(broadcastContainer)-6:len(broadcastContainer)]
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
	numInRow := 0
	for i := 0;i < 12;i++ {
		small.RemoveRow(0)
	}
	row := 0
	numCount := 0
	broadRightUn, err := twoBuilder.GetObject("broadRight")
	if err != nil {
		panic(err)
	}
	broadRight := broadRightUn.(*gtk.Popover)
	for i := 0;i < len(buttonContainer);i++ {
                buttonContainer[i].Connect("button-press-event", func (butt *gtk.Button, ev *gdk.Event) {
                        keyEvent := gdk.EventButtonNewFromEvent(ev)
                        if keyEvent.ButtonVal() == 1 {
				broadRight.SetVisible(false)
				broadRight.Show()
                        }
                        if keyEvent.ButtonVal() == 2 {
				broadRight.SetVisible(false)
				broadRight.Show()
                        }
                        if keyEvent.ButtonVal() == 3 {
				broadRight.SetRelativeTo(butt)
				broadRight.SetVisible(true)
				broadRight.Show()
                        }
			gtk.MainIterationDo(true)

                })
		if numCount < numInRow {
//			small.Add(buttonContainer[i])
//			buttonContainer[i].SetVAlign(gtk.ALIGN_CENTER)
//			buttonContainer[i].SetVExpand(true)
			box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
			if err != nil {
				panic(err)
			}
			box.SetCenterWidget(buttonContainer[i])
			box.SetChildPacking(buttonContainer[i], true, true, 1, gtk.PACK_START)
//			box.PackEnd(buttonContainer[i], true, true, 0)
			box.SetVAlign(gtk.ALIGN_FILL)
			box.SetVExpand(true)
			small.Attach(box, numCount, row, 1, 1)
			numCount++
			small.ShowAll()
//			small.ShowAll()
//			fmt.Println("Num in row", numCount)

		}else {
//			small.InsertRow(row)
//			buttonContainer[i].SetVAlign(gtk.ALIGN_CENTER)
//			buttonContainer[i].SetVExpand(true)
			box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
			if err != nil {
				panic(err)
			}
			box.SetCenterWidget(buttonContainer[i])
			box.SetChildPacking(buttonContainer[i], true, true, 1, gtk.PACK_START)
//			box.PackEnd(buttonContainer[i], true, true, 0)
			box.SetVAlign(gtk.ALIGN_FILL)
			box.SetVExpand(true)
			small.Attach(box, numCount, row, 1, 1)
			row++
			numCount = 0
//			fmt.Println("row", row)
//			small.ShowAll()
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
		if count == 140 {
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
	messageLabel.SetJustify(gtk.JUSTIFY_CENTER)
	messageLabel.SetVAlign(gtk.ALIGN_FILL)
	messageLabel.SetVExpand(true)

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
	    twoBuilder, err := gtk.BuilderNewFromFile("client.glade")
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
		//launch the register screen
		registerButtonUn, err := twoBuilder.GetObject("register")
		if err != nil {
			panic(err)
		}
		registerButton := registerButtonUn.(*gtk.Button)
		registerButton.Connect("clicked", func() {
			register(application, twoBuilder)
		})
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
        		        play := LogPlayerIn(user, pass)

//				play := InitPlayer(user, pass)
				whoList := who(play.Name)
//	                	go actOn(play, fileChange, whoList)
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
