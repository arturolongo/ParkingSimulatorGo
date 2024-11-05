package models

type Vehicle struct {
	ID        int
	ImagePath string
}

func NewVehicle(id int, imagePath string) *Vehicle {
	return &Vehicle{
		ID:        id,
		ImagePath: imagePath,
	}
}
