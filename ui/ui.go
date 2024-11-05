package ui

import (
	"SimuladorGo/models"
	"SimuladorGo/services"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simulador de Estacionamiento")

	parking := models.NewParking(models.MaxParkingSpaces)

	vehicleService := services.NewVehicleService(parking)

	messageLog := widget.NewLabel("")

	parkingContainer := container.New(layout.NewGridLayout(5))

	updateUI := func(occupiedSpaces []bool) {
		parkingContainer.Objects = nil
		for i := 0; i < models.MaxParkingSpaces; i++ {
			if occupiedSpaces[i] {
				carImage := canvas.NewImageFromFile("assets/car1.png")
				carImage.Resize(fyne.NewSize(40, 40))
				parkingContainer.Add(carImage)
			} else {
				parkingContainer.Add(canvas.NewText("Espacio vacío", nil))
			}
		}
		parkingContainer.Refresh()

		messageLog.SetText("Vehículo saliendo del espacio...")
	}

	vehicleService.StartSimulation(1*time.Second, updateUI)

	myWindow.SetContent(container.NewVBox(parkingContainer, messageLog))
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}
