package factory

import "fmt"

type DataConnection interface {
	Process()
	Dose()
}

type WIFIConnection struct{}

func (w *WIFIConnection) Process() {
	fmt.Println("Wifi connection on proccess")
}

func (w *WIFIConnection) Dose() {
	fmt.Println("Wifi connection on dose")
}

type HotspotConnection struct{}

func (h *HotspotConnection) Process() {
	fmt.Println("Hotspot connection on proccess")
}

func (h *HotspotConnection) Dose() {
	fmt.Println("Hotspot connection on dose")
}

type ConnectionFactory interface {
	GetConnection(ip_address string) DataConnection
}

type MyConnection struct{}

func (h *MyConnection) GetConnection(ip_address string) DataConnection {

	if ip_address == "192.168.0.1" {
		return &WIFIConnection{}
	}

	if ip_address == "192.168.0.2" {
		return &HotspotConnection{}
	}

	return nil

}

// interface
// type Button interface {
// 	Render()
// 	OnClick()
// }

// type WindowButton struct{}

// func (w *WindowButton) Render() {
// 	fmt.Println("Windows button on render active")
// }

// func (w *WindowButton) OnClick() {
// 	fmt.Println("Windows button on click active")
// }

// type HTMLButton struct{}

// func (h *HTMLButton) Render() {
// 	fmt.Println("HTML button on render active")
// }

// func (h *HTMLButton) OnClick() {
// 	fmt.Println("HTML button on click active")
// }

// type Dialog interface {
// 	CreateButton(tipe string) Button
// }

// type WindowDialog struct{}

// func (w *WindowDialog) CreateButton(tipe string) Button {

// 	if tipe == "HTML" {
// 		return &HTMLButton{}
// 	}

// 	if tipe == "WINDOWS" {
// 		return &WindowButton{}
// 	}

// 	return nil
// }

// type WebDialog struct{}

// func (w *WebDialog) CreateButton(tipe string) Button {

// 	if tipe == "HTML" {
// 		return &HTMLButton{}
// 	}

// 	if tipe == "WINDOWS" {
// 		return &WindowButton{}
// 	}

// 	return nil

// }
