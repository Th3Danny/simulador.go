package main

import (
	"math/rand"
	"simulador/src/app"
	"simulador/src/domain"
	"simulador/src/infrastructure/interfaz"
	"time"
)

func main() {
	// Crear un estacionamiento con capacidad para 20 vehículos
	estacionamiento := domain.NuevoEstacionamiento(20)

	// Crear el controlador
	controlador := app.NuevoControlador(estacionamiento, rand.New(rand.NewSource(time.Now().UnixNano())))

	interfaz := interfaz.NuevaInterfaz(controlador)
	estacionamiento.AgregarObservador(interfaz)

	go controlador.IniciarSimulacion()
	interfaz.Iniciar()
}
