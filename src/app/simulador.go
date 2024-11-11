package app

import (
	"math/rand"
	"simulador/src/domain"
	"time"
)

// EjecutarSimulacion inicia el proceso de simulación, creando los vehículos y manejando el ciclo de entrada al estacionamiento.
func EjecutarSimulacion(estacionamiento *domain.Estacionamiento) {
	// Inicializa el generador de números aleatorios con la semilla actual del tiempo
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crea una nueva instancia del controlador y le pasa el generador aleatorio
	controlador := NuevoControlador(estacionamiento, rnd)

	// Registra el controlador como observador del estacionamiento
	estacionamiento.AgregarObservador(controlador)

	// Inicia el ciclo de la simulación en una goroutine para que sea asincrónico
	go controlador.IniciarSimulacion()

	// Aquí, si necesitas gestionar más ciclos o tareas adicionales, puedes hacerlo
	select {} // Mantiene el programa corriendo
}
