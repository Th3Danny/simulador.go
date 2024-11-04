package main

import (
    "simulador/src/app"
    "simulador/src/domain"
    "math/rand"
    "time"
    "simulador/src/infrastructure/interfaz"
)

func main() {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    estacionamiento := domain.NuevoEstacionamiento(20)
    controlador := app.NuevoControlador(estacionamiento, rnd)

 
    ui := interfaz.NuevaInterfaz(controlador)

    // Iniciar la simulación en la gorutine
    go controlador.IniciarSimulacion()


    ui.Iniciar()
}
