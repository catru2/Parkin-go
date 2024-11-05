package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	id          int
	parkingTime time.Duration
	space       int
}

func NewCar(id int) *Car {
	return &Car{
		id:          id,
		parkingTime: time.Duration(rand.Intn(10)+10) * time.Second,
		space:       0,
	}
}

func (c *Car) Enter(p *Parking) int {
	p.GetSpaces() <- c.GetId()
	p.GetEntrance().Lock()
	defer p.GetEntrance().Unlock()

	spacesArray := p.GetSpacesArray()

	fmt.Printf("Auto %d ha entrado. Espacios ocupados: %d\n", c.GetId(), len(p.GetSpaces()))

	for i := 0; i < len(spacesArray); i++ {
		if !spacesArray[i] {
			spacesArray[i] = true
			c.space = i
			p.SetSpacesArray(spacesArray)
			return i // Retorna el espacio asignado
		}
	}
	return -1
}

func (c *Car) Leave(p *Parking) {
	p.GetEntrance().Lock()
	defer p.GetEntrance().Unlock()

	<-p.GetSpaces()

	spacesArray := p.GetSpacesArray()
	spacesArray[c.space] = false
	p.SetSpacesArray(spacesArray)

	fmt.Printf("Auto %d ha salido. Espacios ocupados: %d\n", c.GetId(), len(p.GetSpaces()))
}

func (c *Car) Park(p *Parking, wg *sync.WaitGroup) {
	defer wg.Done()

	// Se maneja el tiempo de espera antes de salir
	c.Enter(p)
	time.Sleep(c.parkingTime)
	c.Leave(p)
}

func (c *Car) GetId() int {
	return c.id
}

func (c *Car) GetSpace() int {
	return c.space
}

func (c *Car) GetParkingTime() time.Duration {
	return c.parkingTime
}
