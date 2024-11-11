package app

import (
    "fmt"
    "math/rand"
    "simulador/src/domain"
    "sync"
    "time"
)

type ControladorSimulacion struct {
    estacionamiento     *domain.Estacionamiento
    rnd                 *rand.Rand
    log                 []string
    logMutex            sync.Mutex
    vehiculosEnEspera   chan *domain.Vehiculo 
    actualizaciones     chan struct{}       
}


func NuevoControlador(estacionamiento *domain.Estacionamiento, rnd *rand.Rand) *ControladorSimulacion {
    return &ControladorSimulacion{
        estacionamiento:   estacionamiento,
        rnd:               rnd,
        log:               make([]string, 0),
        vehiculosEnEspera: make(chan *domain.Vehiculo, 100), 
        actualizaciones:   make(chan struct{}, 1),           
    }
}

// Implementa el método `Actualizar` requerido por la interfaz `Observador`.
func (c *ControladorSimulacion) Actualizar(vehiculo *domain.Vehiculo) {
    fmt.Printf("Nuevo vehículo generado: ID = %d\n", vehiculo.ID)
    go func() {
        c.vehiculosEnEspera <- vehiculo 
    }()
}

// Implementa el método `ActualizarEstadoEstacionamiento` requerido por la interfaz `Observador`.
func (c *ControladorSimulacion) ActualizarEstadoEstacionamiento() {
   
    select {
    case c.actualizaciones <- struct{}{}:
        fmt.Println("Actualización del estado del estacionamiento enviada.")
    default:
       
    }
}

func (c *ControladorSimulacion) IniciarSimulacion() {
    go func() {
        id := 1
        for {
            vehiculo := domain.GenerarVehiculo(id, c.rnd)
            fmt.Println("Generando vehículo:", vehiculo.ID)
            c.Actualizar(vehiculo) 
            // Simula el tiempo entre la llegada de vehículos
            time.Sleep(time.Duration(c.rnd.ExpFloat64()) * time.Second)
            id++
        }
    }()

    
    go func() {
        for vehiculo := range c.vehiculosEnEspera {
            c.intentarEntrada(vehiculo)
        }
    }()

   
    go func() {
        for range c.actualizaciones {
            if len(c.vehiculosEnEspera) > 0 && c.estacionamiento.IntentarEntrar() {
                vehiculo := <-c.vehiculosEnEspera
                mensaje := fmt.Sprintf("Vehículo %d ha entrado desde la espera", vehiculo.ID)
                fmt.Println(mensaje)
                c.agregarLog(mensaje)
            }
        }
    }()
}

func (c *ControladorSimulacion) intentarEntrada(vehiculo *domain.Vehiculo) {
    fmt.Printf("Intentando entrada para vehículo %d\n", vehiculo.ID)
    if c.estacionamiento.IntentarEntrar() {
        mensaje := fmt.Sprintf("Vehículo %d ha entrado", vehiculo.ID)
        fmt.Println(mensaje)
        c.agregarLog(mensaje)
        
       
        go func() {
            time.Sleep(time.Duration(vehiculo.Tiempo) * time.Second)
            c.estacionamiento.Salir() 
            mensaje = fmt.Sprintf("Vehículo %d ha salido", vehiculo.ID)
            fmt.Println(mensaje)
            c.agregarLog(mensaje)
        }()
    } else {
        mensaje := fmt.Sprintf("Vehículo %d esperando para entrar", vehiculo.ID)
        fmt.Println(mensaje)
        c.agregarLog(mensaje)
        go func() {
            c.vehiculosEnEspera <- vehiculo 
        }()
    }
}


func (c *ControladorSimulacion) agregarLog(mensaje string) {
    c.logMutex.Lock()
    defer c.logMutex.Unlock()
    c.log = append(c.log, mensaje)
}


func (c *ControladorSimulacion) EstacionamientoOcupado() int {
    return c.estacionamiento.Ocupados()
}
