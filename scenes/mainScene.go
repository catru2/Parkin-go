package scenes

import (
	"concurrent-parking/characters"
	"concurrent-parking/models"
	"fmt"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

var carsContainer = container.NewWithoutLayout()

func (s *MainScene) Show() {
	// Fondo del estacionamiento
	parkingBackground := canvas.NewRectangle(color.RGBA{240, 240, 240, 255})
	parkingBackground.Resize(fyne.NewSize(600, 650))
	parkingBackground.Move(fyne.NewPos(0, 0))
	carsContainer.Add(parkingBackground)

	// Límite o contorno del estacionamiento
	rectangle := canvas.NewRectangle(color.Transparent)
	rectangle.StrokeWidth = 2
	rectangle.StrokeColor = color.RGBA{0, 0, 0, 255}
	rectangle.Resize(fyne.NewSize(500, 620))
	rectangle.Move(fyne.NewPos(50, 10))
	carsContainer.Add(rectangle)

	// Entrada al estacionamiento (la puerta)
	gate := canvas.NewRectangle(color.RGBA{0, 128, 0, 255}) // Cambia de color para resaltar la puerta
	gate.Resize(fyne.NewSize(10, 100))
	gate.Move(fyne.NewPos(40, 300)) // Cambiar posición para alinearlo con el contorno
	carsContainer.Add(gate)

	s.window.SetContent(carsContainer)
}

func (s *MainScene) Run() {
	p := models.NewParking(make(chan int, 10), &sync.Mutex{})
	poissonDist := models.NewPoissonDist()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Crear auto y configurar su imagen
			car := models.NewCar(id)
			carCharacter := characters.NewCarCharacter()
			carImage := carCharacter.GetImage()
			carImage.Resize(fyne.NewSize(40, 20))
			carImage.Move(fyne.NewPos(-20, 310))
			carsContainer.Add(carImage)
			carsContainer.Refresh()

			// Animación de entrada
			carCharacter.ParkAnimation(carsContainer)

			// Lógica de asignación de espacio
			space := car.Enter(p)
			if space >= 0 {
				carCharacter.AnimateEntry(carsContainer, space)
			}

			// Tiempo de espera
			time.Sleep(car.GetParkingTime())

			// Animación de salida
			carsContainer.Remove(carImage) // Eliminar imagen de entrada
			carCharacter.AnimateExit(carsContainer)
			car.Leave(p)
		}(i)

		// Tiempo aleatorio de espera usando Poisson
		randPoissonNumber := poissonDist.Generate(float64(2))
		time.Sleep(time.Second * time.Duration(randPoissonNumber))
	}

	wg.Wait()
	fmt.Println("Proceso completado")
}
