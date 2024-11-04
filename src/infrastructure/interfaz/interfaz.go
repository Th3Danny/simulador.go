package interfaz

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"simulador/src/app"
	"time"
)

type Interfaz struct {
	controlador      *app.ControladorSimulacion
	espacios         []*canvas.Image 
	vehiculosEsperar []int
}

func NuevaInterfaz(controlador *app.ControladorSimulacion) *Interfaz {
	return &Interfaz{controlador: controlador}
}

func (i *Interfaz) Iniciar() {
	aplicacion := fyneApp.New()
	ventana := aplicacion.NewWindow("Simulador de Estacionamiento")

	i.espacios = make([]*canvas.Image, 20)
	grid := container.NewGridWithColumns(5)

	// Inicializar los espacios
	for j := 0; j < 20; j++ {
		img := canvas.NewImageFromFile("assets/estacionamineto.jpg") // Espacio libre
		img.SetMinSize(fyne.NewSize(80, 50))
		i.espacios[j] = img
		grid.Add(img)
	}

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ocupacion := i.controlador.EstacionamientoOcupado()
			fmt.Printf("Ocupación actual: %d\n", ocupacion)

			// Actualizar el estado de los espacios
			for j := 0; j < 20; j++ {
				if j < ocupacion {
					// Cambiar a la imagen del vehículo
					i.espacios[j] = canvas.NewImageFromFile("assets/car.png")
					i.espacios[j].SetMinSize(fyne.NewSize(80, 40))
				} else {
					// Mantener la imagen de fondo para el espacio libre
					i.espacios[j] = canvas.NewImageFromFile("assets/estacionamineto.jpg")
				}
				grid.Objects[j] = i.espacios[j]
				i.espacios[j].Refresh()
			}

			// Mostrar vehículos en espera
			if len(i.controlador.VehiculosEnEspera) > 0 {
				fmt.Printf("Vehículos en espera: %v\n", i.controlador.VehiculosEnEspera)
				for _, vehiculoID := range i.controlador.VehiculosEnEspera {
					fmt.Printf("Vehículo %d está esperando para entrar\n", vehiculoID)
				}
			}

			
			grid.Refresh()
		}
	}()

	ventana.SetContent(container.NewVBox(
		widget.NewLabel("Estado del Estacionamiento"),
		grid,
	))
	ventana.Resize(fyne.NewSize(400, 600))
	ventana.ShowAndRun()
}
