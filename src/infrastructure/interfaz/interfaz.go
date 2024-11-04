package interfaz

import (
    fyneApp "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2"
    "simulador/src/app"
    "fmt"
    "time"
)

type Interfaz struct {
    controlador *app.ControladorSimulacion
    espacios     []*canvas.Image 
}

func NuevaInterfaz(controlador *app.ControladorSimulacion) *Interfaz {
    return &Interfaz{controlador: controlador}
}

func (i *Interfaz) Iniciar() {
    aplicacion := fyneApp.New()
    ventana := aplicacion.NewWindow("Simulador de Estacionamiento")

    // Inicializar los espacios para los vehículos
    i.espacios = make([]*canvas.Image, 20)
    grid := container.NewGridWithColumns(5)

    
    for j := 0; j < 20; j++ {
       
        img := canvas.NewImageFromFile("assets/estacionamineto.jpg") 
        img.SetMinSize(fyne.NewSize(80, 50)) 
        i.espacios[j] = img
        grid.Add(img) // Agregar el espacio al grid
    }

    go func() {
        for {
            time.Sleep(500 * time.Millisecond) // Controla la frecuencia de actualización
            ocupacion := i.controlador.EstacionamientoOcupado()
            fmt.Printf("Ocupación actual: %d\n", ocupacion) 

            // Actualizar el estado de los espacios
            for j := 0; j < 20; j++ {
                if j < ocupacion {
                   
                    imgVehiculo := canvas.NewImageFromFile("assets/car.png") 
                    imgVehiculo.SetMinSize(fyne.NewSize(80, 40)) 
                    i.espacios[j] = imgVehiculo // Actualizar el espacio con la imagen del vehículo
                } else {
                    // Mantener la imagen de fondo para el espacio libre
                    imgEspacio := canvas.NewImageFromFile("assets/estacionamineto.jpg") 
                    i.espacios[j] = imgEspacio
                }
                grid.Objects[j] = i.espacios[j] 
                i.espacios[j].Refresh() 
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
