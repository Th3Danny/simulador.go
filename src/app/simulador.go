package app

import (
    "simulador/src/domain"
    "math/rand"
    "time"
)

func EjecutarSimulacion(estacionamiento *domain.Estacionamiento) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano())) // Crear el generador aleatorio local
    controlador := NuevoControlador(estacionamiento, rnd)  // Ahora acepta dos argumentos
    controlador.IniciarSimulacion()
}
