package main

import (
	"design_pattern/factory"
)

func main() {

	var loadConnection factory.ConnectionFactory

	loadConnection = &factory.MyConnection{}

	wifi := loadConnection.GetConnection("192.168.0.1")

	wifi.Process()
	wifi.Dose()

	// var createDialog factory.Dialog

	// createDialog = &factory.WindowDialog{}

	// windows := createDialog.CreateButton("WINDOWS")

	// windows.Render()
	// windows.OnClick()

	// //web dialog
	// createDialog = &factory.WebDialog{}

	// web := createDialog.CreateButton("HTML")
	// web.Render()
	// web.OnClick()

}
