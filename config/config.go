package config

import "log"

const (
	MaxParkingSpaces   = 20
	VehicleArrivalRate = 100
)

func Initialize() {
	log.Println("Inicializando configuraciones del simulador...")
}
