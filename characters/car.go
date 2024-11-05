package characters

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type CarCharacter struct {
	image     *canvas.Image
	exitImage *canvas.Image
}

func NewCarCharacter() *CarCharacter {
	image := canvas.NewImageFromURI(storage.NewFileURI("./assets/car.png"))
	exitImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/car_exit.png"))
	return &CarCharacter{
		image:     image,
		exitImage: exitImage,
	}
}

func (cc *CarCharacter) AnimateEntry(carsContainer *fyne.Container, spaceIndex int) {
	for i := 0; i < 5; i++ {
		cc.image.Move(fyne.NewPos(cc.image.Position().X+20, cc.image.Position().Y))
		time.Sleep(time.Millisecond * 200)
	}
	cc.image.Move(fyne.NewPos(290, float32(15+(spaceIndex*30))))
	carsContainer.Refresh()
}

func (cc *CarCharacter) AnimateExit(carsContainer *fyne.Container) {
	// Posicionar `exitImage` en la misma posición que `image` antes de la salida
	cc.exitImage.Move(cc.image.Position())
	carsContainer.Add(cc.exitImage)
	carsContainer.Refresh()

	// Animación de salida hacia la izquierda
	for i := 0; i < 10; i++ {
		cc.exitImage.Move(fyne.NewPos(cc.exitImage.Position().X-30, cc.exitImage.Position().Y))
		carsContainer.Refresh()
		time.Sleep(time.Millisecond * 200)
	}

	// Eliminar `exitImage` del contenedor después de la animación
	carsContainer.Remove(cc.exitImage)
	carsContainer.Refresh()
}

func (cc *CarCharacter) ParkAnimation(carsContainer *fyne.Container) {
	for i := 0; i < 7; i++ {
		cc.image.Move(fyne.NewPos(cc.image.Position().X+20, cc.image.Position().Y))
		time.Sleep(time.Millisecond * 200)
	}
	carsContainer.Refresh()
}

func (cc *CarCharacter) GetImage() *canvas.Image {
	return cc.image
}

func (cc *CarCharacter) GetExitImage() *canvas.Image {
	return cc.exitImage
}
