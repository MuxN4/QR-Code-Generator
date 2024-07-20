package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("QR Code Generator")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter data to encode")

	qrCodeContainer := container.NewVBox()

	generateButton := widget.NewButton("Generate QR Code", func() {
		data := input.Text
		if data == "" {
			dialog.ShowInformation("Error", "Please enter data to encode", myWindow)
			return
		}

		err := qrcode.WriteFile(data, qrcode.Medium, 256, "qrcode.png")
		if err != nil {
			dialog.ShowInformation("Error", "Failed to generate QR code", myWindow)
			log.Fatalf("Failed to generate QR code: %v", err)
		}

		file, err := os.Open("qrcode.png")
		if err != nil {
			dialog.ShowInformation("Error", "Failed to open generated QR code", myWindow)
			log.Fatalf("Failed to open QR code image: %v", err)
		}
		defer file.Close()

		img := canvas.NewImageFromFile("qrcode.png")
		img.FillMode = canvas.ImageFillContain
		img.SetMinSize(fyne.NewSize(256, 256))
		qrCodeContainer.Objects = []fyne.CanvasObject{img}
		qrCodeContainer.Refresh()
	})

	content := container.NewVBox(
		widget.NewLabel("QR Code Generator"),
		input,
		generateButton,
		qrCodeContainer,
	)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 400))

	myWindow.ShowAndRun()
}
