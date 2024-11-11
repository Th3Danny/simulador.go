package app

import (
	"fmt"
	"simulador/src/domain"
	"sync"
	"time"
	"math/rand"
)

type ControladorSimulacion struct {
	estacionamiento   *domain.Estacionamiento
	rnd               *rand.Rand
	log               []string
	mutex             sync.Mutex
	VehiculosEnEspera []int
}

// NuevoControlador crea una nueva instancia del controlador de simulación.
func NuevoControlador(estacionamiento *domain.Estacionamiento, rnd *rand.Rand) *ControladorSimulacion {
	return &ControladorSimulacion{
		estacionamiento:   estacionamiento,
		rnd:               rnd,
		log:               make([]string, 0),
		VehiculosEnEspera: make([]int, 0),
	}
}

// Implementa el método `Actualizar` requerido por la interfaz `Observador`.
func (c *ControladorSimulacion) Actualizar(vehiculo *domain.Vehiculo) {
	// Lógica para manejar la actualización cuando se genera un nuevo vehículo.
	fmt.Printf("Nuevo vehículo generado: ID = %d\n", vehiculo.ID)
	c.intentarEntrada(vehiculo)
}

// Implementa el método `ActualizarEstadoEstacionamiento` requerido por la interfaz `Observador`.
func (c *ControladorSimulacion) ActualizarEstadoEstacionamiento() {
    // Lógica para manejar actualizaciones desde el estacionamiento
    if len(c.VehiculosEnEspera) > 0 && c.estacionamiento.IntentarEntrar() {
        vehiculoID := c.VehiculosEnEspera[0]
        c.VehiculosEnEspera = c.VehiculosEnEspera[1:]
        mensaje := fmt.Sprintf("Vehículo %d ha entrado desde la espera", vehiculoID)
        fmt.Println(mensaje)
        c.agregarLog(mensaje)
    }
}


func (c *ControladorSimulacion) IniciarSimulacion() {
    id := 1
    for {
        // Genera un nuevo vehículo
        vehiculo := domain.GenerarVehiculo(id, c.rnd)
        fmt.Println("Generando vehículo:", vehiculo.ID)
        go c.intentarEntrada(vehiculo) // Llama a intentarEntrada en una goroutine
        // Simula el tiempo entre la llegada de vehículos (puede ajustarse según sea necesario)
        time.Sleep(time.Duration(c.rnd.ExpFloat64()) * time.Second)
        id++
    }
}


func (c *ControladorSimulacion) intentarEntrada(vehiculo *domain.Vehiculo) {
    fmt.Printf("Intentando entrada para vehículo %d\n", vehiculo.ID)
    if c.estacionamiento.IntentarEntrar() {
        mensaje := fmt.Sprintf("Vehículo %d ha entrado", vehiculo.ID)
        fmt.Println(mensaje)
        c.agregarLog(mensaje)
        
        // Manejar el tiempo de permanencia en una goroutine
        go func() {
            time.Sleep(time.Duration(vehiculo.Tiempo) * time.Second)
            c.estacionamiento.Salir() // Simula la salida del vehículo después de cierto tiempo
            mensaje = fmt.Sprintf("Vehículo %d ha salido", vehiculo.ID)
            fmt.Println(mensaje)
            c.agregarLog(mensaje)
        }()
    } else {
        mensaje := fmt.Sprintf("Vehículo %d esperando para entrar", vehiculo.ID)
        fmt.Println(mensaje)
        c.agregarLog(mensaje)
        c.VehiculosEnEspera = append(c.VehiculosEnEspera, vehiculo.ID)
    }
}





// Método para agregar un mensaje al log de la simulación.
func (c *ControladorSimulacion) agregarLog(mensaje string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.log = append(c.log, mensaje)
}

// Método para obtener el número de espacios ocupados en el estacionamiento.
func (c *ControladorSimulacion) EstacionamientoOcupado() int {
	return c.estacionamiento.Ocupados()
}
