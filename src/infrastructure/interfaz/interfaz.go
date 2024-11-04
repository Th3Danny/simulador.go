package interfaz

import (
    fyneApp "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2"
    "simulador/src/app"
    "image/color"
    "fmt"
    "time"
)

type Interfaz struct {
    controlador *app.ControladorSimulacion
    espacios    []*canvas.Rectangle
}

func NuevaInterfaz(controlador *app.ControladorSimulacion) *Interfaz {
    return &Interfaz{controlador: controlador}
}

func (i *Interfaz) Iniciar() {
    aplicacion := fyneApp.New()
    ventana := aplicacion.NewWindow("Simulador de Estacionamiento")

    i.espacios = make([]*canvas.Rectangle, 20)
    grid := container.NewGridWithColumns(5)

    for j := 0; j < 20; j++ {
        rect := canvas.NewRectangle(color.Gray{Y: 200})
        rect.SetMinSize(fyne.NewSize(50, 50))
        i.espacios[j] = rect
        grid.Add(rect)
    }

    go func() {
        for {
            time.Sleep(500 * time.Millisecond)
            ocupacion := i.controlador.EstacionamientoOcupado()
            fmt.Printf("Ocupación actual: %d\n", ocupacion) // Mensaje de depuración
    
            for j := 0; j < 20; j++ {
                if j < ocupacion {
                    i.espacios[j].FillColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
                } else {
                    i.espacios[j].FillColor = color.Gray{Y: 200}
                }
                i.espacios[j].Refresh()
            }
        }
    }()
    
    
    ventana.SetContent(container.NewVBox(
        widget.NewLabel("Estado del Estacionamiento"),
        grid,
    ))
    ventana.Resize(fyne.NewSize(400, 600))
    ventana.ShowAndRun()
}
