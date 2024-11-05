package services

import (
	"SimuladorGo/models"
	"fmt"
	"math/rand"
	"time"
)

type ParkingService struct {
	parking *models.Parking
}

func NewParkingService(parking *models.Parking) *ParkingService {
	return &ParkingService{
		parking: parking,
	}
}

// simula el uso de espacios de estacionamiento
func (ps *ParkingService) SimulateParking(updateUI func([]bool)) {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

			space, success := ps.parking.Enter()
			if success {
				fmt.Printf("Vehículo entró al espacio %d\n", space)
	
				// tiempo del vehículo estacionado
				parkingDuration := time.Duration(rand.Intn(2)+1) * time.Second
				go func(space int) {
					time.Sleep(parkingDuration)
					ps.parking.Leave(space)
					fmt.Printf("Vehículo salió del espacio %d\n", space)
					updateUI(ps.parking.GetOccupiedSpaces())
				}(space)

				updateUI(ps.parking.GetOccupiedSpaces())
			} else {
				fmt.Println("Estacionamiento lleno, vehículo esperando...")
			}
		}
	}()
}
