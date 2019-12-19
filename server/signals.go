package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)


func signals(twoBuilder *gtk.Builder) {
	listTellsUn, err := twoBuilder.GetObject("listTells")
	if err != nil {
		panic(err)
	}
	listTells := listTellsUn.(*gtk.Button)

	listTells.Connect("clicked", func() {
		fmt.Println("Button was pressed!")
	})
}
