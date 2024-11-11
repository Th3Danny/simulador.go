package app

import (
	"math/rand"
	"simulador/src/domain"
	"time"
)


func EjecutarSimulacion(estacionamiento *domain.Estacionamiento) {
	
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	controlador := NuevoControlador(estacionamiento, rnd)
	
	estacionamiento.AgregarObservador(controlador)
	
	go controlador.IniciarSimulacion()
	
	select {} // Mantiene el programa corriendo
}
