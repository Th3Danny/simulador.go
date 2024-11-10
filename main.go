package main

import (
    "simulador/src/app"
    "simulador/src/infrastructure/interfaz"
    "simulador/src/domain"
    "math/rand"
    "time"
)

func main() {
    estacionamiento := domain.NuevoEstacionamiento(20)
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    controlador := app.NuevoControlador(estacionamiento, rnd)

    // Crear la interfaz y pasarle el controlador
    ui := interfaz.NuevaInterfaz(controlador)

    // Iniciar la simulación en una goroutine
    go controlador.IniciarSimulacion()

    // Iniciar la interfaz gráfica
    ui.Iniciar()
}

