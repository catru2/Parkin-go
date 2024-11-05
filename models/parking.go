package models

import (
	"sync"
)

type Parking struct {
	spaces      chan int
	entrance    *sync.Mutex
	spacesArray [20]bool
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		spaces:      spaces,
		entrance:    entrance,
		spacesArray: [20]bool{},
	}
}

func (p *Parking) GetSpaces() chan int {
	return p.spaces
}

func (p *Parking) GetEntrance() *sync.Mutex {
	return p.entrance
}

func (p *Parking) GetSpacesArray() [20]bool {
	return p.spacesArray
}

func (p *Parking) SetSpacesArray(spacesArray [20]bool) {
	p.spacesArray = spacesArray
}

func (p *Parking) AllocateSpace() int {
	for i := 0; i < len(p.spacesArray); i++ {
		if !p.spacesArray[i] {
			p.spacesArray[i] = true
			return i // Retorna el índice del espacio asignado
		}
	}
	return -1 // -1 indica que no hay espacios disponibles
}

func (p *Parking) ReleaseSpace(spaceIndex int) {
	p.spacesArray[spaceIndex] = false
}
