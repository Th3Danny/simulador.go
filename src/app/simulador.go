package app

import (
    "math/rand"
    "time"
    "simulador/src/domain"
)

func EjecutarSimulacion(estacionamiento *domain.Estacionamiento) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano())) // Inicializa el generador de números aleatorios
    controlador := NuevoControlador(estacionamiento, rnd)   // Pasa el generador al controlador
    controlador.IniciarSimulacion()
}
