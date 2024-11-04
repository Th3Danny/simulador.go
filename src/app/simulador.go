package app

import (
    "simulador/src/domain"
    "math/rand"
    "time"
)

func EjecutarSimulacion(estacionamiento *domain.Estacionamiento) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano())) 
    controlador := NuevoControlador(estacionamiento, rnd)  
    controlador.IniciarSimulacion()
}
