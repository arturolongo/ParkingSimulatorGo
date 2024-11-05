package services

import (
	"SimuladorGo/models"
	"fmt"
	"math/rand"
	"time"
)

type VehicleService struct {
	parking       *models.Parking
	vehicleNumber int 
}

func NewVehicleService(parking *models.Parking) *VehicleService {
	return &VehicleService{
		parking:       parking,
		vehicleNumber: 0,
	}
}

func (vs *VehicleService) StartSimulation(arrivalRate time.Duration, updateUI func([]bool)) {
	go func() {
		for {
			time.Sleep(arrivalRate)

			vs.vehicleNumber++

			// Intenta que un vehículo entre al estacionamiento
			space, success := vs.parking.Enter()
			if success {
				vehicleID := vs.vehicleNumber
				fmt.Printf("Vehículo %d entrando al espacio %d\n", vehicleID, space)

				parkingDuration := time.Duration(rand.Intn(2)+1) * time.Second
				go func(space int, vehicleID int) {
					time.Sleep(parkingDuration)
					vs.parking.Leave(space)
					fmt.Printf("Vehículo %d saliendo del espacio %d\n", vehicleID, space) // Mensaje en la consola
					updateUI(vs.parking.GetOccupiedSpaces())
				}(space, vehicleID)

				updateUI(vs.parking.GetOccupiedSpaces())
			}
		}
	}()
}
