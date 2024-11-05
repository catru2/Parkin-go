package characters

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type ParkingVisual struct {
	gate *canvas.Rectangle
}

func NewParkingVisual() *ParkingVisual {
	gate := canvas.NewRectangle(color.RGBA{245, 243, 39, 250})
	gate.Resize(fyne.NewSize(10, 100))
	gate.Move(fyne.NewPos(195, 300))

	return &ParkingVisual{
		gate: gate,
	}
}

func (pv *ParkingVisual) GetGate() *canvas.Rectangle {
	return pv.gate
}

func (pv *ParkingVisual) ShowExitQueue(carsContainer *fyne.Container, carImage *canvas.Image) {
	// Mover la imagen del auto a la posición inicial de la cola de salida
	carImage.Move(fyne.NewPos(205, 350))
	carsContainer.Add(carImage)
	carsContainer.Refresh()
}
