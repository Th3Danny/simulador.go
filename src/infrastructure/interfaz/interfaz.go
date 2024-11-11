package interfaz

import (
    "fmt"
    "fyne.io/fyne/v2"
    fyneApp "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "simulador/src/app"
    "simulador/src/domain"
)

type Interfaz struct {
    controlador    *app.ControladorSimulacion
    espacios       []*canvas.Image
    notificaciones chan struct{} // Canal para notificaciones
}

func NuevaInterfaz(controlador *app.ControladorSimulacion) *Interfaz {
    return &Interfaz{
        controlador:    controlador,
        espacios:       make([]*canvas.Image, 20),
        notificaciones: make(chan struct{}, 1), // Canal con capacidad 1 para evitar bloqueos
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

    ventana.SetContent(container.NewVBox(
        widget.NewLabel("Estado del Estacionamiento"),
        grid,
    ))

    ventana.Resize(fyne.NewSize(400, 600))
    ventana.Show()

    // Goroutine para manejar notificaciones de actualización
    go func() {
        for range i.notificaciones {
            fmt.Println("Notificación recibida en goroutine de actualización")
            i.actualizarEstado(grid)
        }
    }()

    aplicacion.Run()
}

func (i *Interfaz) actualizarEstado(grid *fyne.Container) {
    ocupacion := i.controlador.EstacionamientoOcupado()
    fmt.Printf("Ocupación actual: %d\n", ocupacion)

    for j := 0; j < 20; j++ {
        var nuevaImagen *canvas.Image
        if j < ocupacion {
            nuevaImagen = canvas.NewImageFromFile("assets/car.png")
            nuevaImagen.SetMinSize(fyne.NewSize(80, 40))
        } else {
            nuevaImagen = canvas.NewImageFromFile("assets/estacionamineto.jpg")
            nuevaImagen.SetMinSize(fyne.NewSize(80, 50))
        }

        // Reemplaza la imagen y refresca
        i.espacios[j] = nuevaImagen
        grid.Objects[j] = nuevaImagen
        nuevaImagen.Refresh()
    }

    grid.Refresh() // Refresca el contenedor que contiene las imágenes
}

func (i *Interfaz) ActualizarEstadoEstacionamiento() {
    fmt.Println("Notificación de actualización enviada")
    select {
    case i.notificaciones <- struct{}{}:
        fmt.Println("Señal enviada a notificaciones")
    default:
        fmt.Println("Canal lleno, evitando bloqueo")
    }
}

// Implementa el método requerido por la interfaz `domain.Observador`
func (i *Interfaz) Actualizar(vehiculo *domain.Vehiculo) {
    // Lógica para manejar la actualización cuando se genera un nuevo vehículo
    fmt.Printf("Notificación recibida para el vehículo con ID %d\n", vehiculo.ID)
    i.actualizarEspacioVehiculo(vehiculo)

    // Notificar a la interfaz
    select {
    case i.notificaciones <- struct{}{}:
        fmt.Println("Notificación enviada al canal para actualizar la interfaz.")
    default:
        fmt.Println("Canal lleno, no se puede enviar la notificación.")
    }
}

// Método auxiliar para actualizar el espacio de un vehículo
func (i *Interfaz) actualizarEspacioVehiculo(vehiculo *domain.Vehiculo) {
    // Aquí se puede agregar la lógica para manejar el vehículo en la interfaz.
    // Este es un ejemplo que muestra cómo podrías manejarlo:
    if vehiculo.ID < len(i.espacios) {
        nuevaImagen := canvas.NewImageFromFile("assets/car.png")
        nuevaImagen.SetMinSize(fyne.NewSize(80, 40))
        i.espacios[vehiculo.ID] = nuevaImagen
        nuevaImagen.Refresh()
        fmt.Printf("Vehículo con ID %d actualizado en la interfaz.\n", vehiculo.ID)
    } else {
        fmt.Printf("El ID del vehículo %d está fuera del rango de los espacios.\n", vehiculo.ID)
    }
}
