package models

import (
	"sync"
)

const MaxParkingSpaces = 20

type Parking struct {
	spaces []bool       
	mu     sync.Mutex   
}

func NewParking(maxSpaces int) *Parking {
	return &Parking{
		spaces: make([]bool, maxSpaces),
	}
}

func (p *Parking) Enter() (int, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Buscar el primer espacio libre
	for i, occupied := range p.spaces {
		if !occupied {
			p.spaces[i] = true // Ocupar el espacio
			return i, true
		}
	}

	return -1, false
}

// Libera un espacio en el estacionamiento
func (p *Parking) Leave(space int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Liberar el espacio
	if space >= 0 && space < len(p.spaces) {
		p.spaces[space] = false
	}
}

func (p *Parking) GetOccupiedSpaces() []bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	occupiedCopy := make([]bool, len(p.spaces))
	copy(occupiedCopy, p.spaces)
	return occupiedCopy
}
