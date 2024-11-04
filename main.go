package main

import (
    "math/rand"
    "time"
    "simulador/src/app"
    "simulador/src/domain"
    "simulador/src/infrastructure/interfaz"
)

func main() {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    estacionamiento := domain.NuevoEstacionamiento(20)
    controlador := app.NuevoControlador(estacionamiento, rnd)
    
    // Inicializar la interfaz gráfica
    ui := interfaz.NuevaInterfaz(controlador)
    
    // Iniciar la simulación de llegada de vehículos en una goroutine separada
    go controlador.IniciarSimulacion()
    
    // Iniciar la interfaz gráfica en el hilo principal
    ui.Iniciar()
}

