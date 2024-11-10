package interfaz

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"simulador/src/app"
)

// Interfaz ahora implementa Observador
type Interfaz struct {
	controlador      *app.ControladorSimulacion
	espacios         []*canvas.Image
	vehiculosEsperar []int
}

// NuevaInterfaz crea una nueva instancia de la interfaz
func NuevaInterfaz(controlador *app.ControladorSimulacion) *Interfaz {
	return &Interfaz{
		controlador: controlador,
	}
}

func (i *Interfaz) Actualizar(ocupacion []int, vehiculosEnEspera []int) {
    fmt.Printf("Ocupación actual: %v\n", ocupacion)
    fmt.Printf("Vehículos en espera: %v\n", vehiculosEnEspera)

    i.vehiculosEsperar = vehiculosEnEspera

    // Actualizar los espacios de estacionamiento
    for j := 0; j < len(ocupacion); j++ {
        var img *canvas.Image
        if ocupacion[j] != 0 {  // Si hay un vehículo en el espacio
            img = canvas.NewImageFromFile("assets/car.png")  // Imagen del vehículo
            img.SetMinSize(fyne.NewSize(80, 50))
        } else {
            img = canvas.NewImageFromFile("assets/estacionamineto.jpg")  // Imagen del espacio vacío
            img.SetMinSize(fyne.NewSize(80, 50))
        }

        // Asignar la imagen al espacio correspondiente
        i.espacios[j] = img
        i.espacios[j].Refresh()  // Refresca la imagen del espacio
    }

    // Actualizar el contenedor completo (la grilla que contiene todos los espacios)
    grid := container.NewGridWithColumns(5)
    for _, img := range i.espacios {
        grid.Add(img)  // Agrega cada imagen a la grilla
    }

    // Actualizar la UI
    grid.Refresh()

    // Opcional: Imprimir los vehículos en espera
    fmt.Println("Vehículos en espera:", vehiculosEnEspera)
    for _, vehiculoID := range vehiculosEnEspera {
        fmt.Printf("Vehículo %d está esperando para entrar\n", vehiculoID)
    }
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

    // Suscribir la interfaz como observador
    i.controlador.AgregarObservador(i)

    // En el método Iniciar() de la interfaz, suscribimos la interfaz al canal
    go func() {
        for {
            // Espera la señal para actualizar
            <-i.controlador.NotificarCanal
            i.controlador.NotificarObservadores()
            grid.Refresh()  // Actualiza el contenido del grid para reflejar los cambios
        }
    }()

    ventana.SetContent(container.NewVBox(
        widget.NewLabel("Estado del Estacionamiento"),
        grid,
    ))
    ventana.Resize(fyne.NewSize(400, 600))
    ventana.ShowAndRun()
}
