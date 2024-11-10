package app

import (
	"fmt"
	"simulador/src/domain"
	"sync"
	"time"
	"math/rand"
)

// Observador es una interfaz que tiene un método para ser notificado sobre los cambios
type Observador interface {
	Actualizar(ocupacion []int, vehiculosEnEspera []int)
}

type ControladorSimulacion struct {
	estacionamiento   *domain.Estacionamiento
	rnd               *rand.Rand
	vehiculosEnEspera []int
	mutex             sync.Mutex
	observadores      []Observador
	// Canal para notificar cambios
	NotificarCanal chan struct{}
    log               []string  // Agrega el campo log aquí
}

// NuevoControlador crea una nueva instancia del controlador
func NuevoControlador(estacionamiento *domain.Estacionamiento, rnd *rand.Rand) *ControladorSimulacion {
	return &ControladorSimulacion{
		estacionamiento:   estacionamiento,
		rnd:               rnd,
		vehiculosEnEspera: make([]int, 0),
		observadores:      make([]Observador, 0),
		// Inicializamos el canal
		NotificarCanal:    make(chan struct{}),
        log:               make([]string, 0), // Inicializa el log como un slice vacío
	}
}

// AgregarObservador agrega un observador a la lista
func (c *ControladorSimulacion) AgregarObservador(o Observador) {
	c.observadores = append(c.observadores, o)
}


func (c *ControladorSimulacion) NotificarObservadores() {
    ocupacion := c.estacionamiento.ObtenerEspacios()  // Obtienes el estado actual de los espacios
    for _, o := range c.observadores {
        o.Actualizar(ocupacion, c.vehiculosEnEspera)  // Notificas a la interfaz con los datos actualizados
    }
}


// IniciarSimulacion empieza el proceso de generación y manejo de vehículos
func (c *ControladorSimulacion) IniciarSimulacion() {
	id := 1
	for {
		vehiculo := domain.GenerarVehiculo(id, c.rnd)
		fmt.Println("Generando vehículo:", vehiculo.ID)
		go c.intentarEntrada(vehiculo)
		time.Sleep(time.Duration(c.rnd.ExpFloat64()) * time.Second) // Usar Poisson para simular llegada
		id++
	}
}

// intentarEntrada maneja la lógica de si un vehículo puede entrar o no
func (c *ControladorSimulacion) intentarEntrada(vehiculo *domain.Vehiculo) {
	if c.estacionamiento.IntentarEntrar(vehiculo.ID) {
		mensaje := fmt.Sprintf("Vehículo %d ha entrado", vehiculo.ID)
		fmt.Println(mensaje)
		c.agregarLog(mensaje)
		time.Sleep(time.Duration(vehiculo.Tiempo) * time.Second)
		c.estacionamiento.Salir(vehiculo.ID)
		mensaje = fmt.Sprintf("Vehículo %d ha salido", vehiculo.ID)
		fmt.Println(mensaje)
		c.agregarLog(mensaje)
	} else {
		mensaje := fmt.Sprintf("Vehículo %d esperando para entrar", vehiculo.ID)
		fmt.Println(mensaje)
		c.agregarLog(mensaje)
		c.vehiculosEnEspera = append(c.vehiculosEnEspera, vehiculo.ID)
	}
}

// Agregar log de eventos
func (c *ControladorSimulacion) agregarLog(mensaje string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.log = append(c.log, mensaje) // Agrega el mensaje al log
}

// Registro devuelve todos los mensajes del log como un solo string
func (c *ControladorSimulacion) Registro() string {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    logString := ""
    for _, entry := range c.log {
        logString += entry + "\n"
    }
    return logString
}
