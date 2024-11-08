package app

import (
    "simulador/src/domain"
    "math/rand"
    "fmt"
    "sync"
    "time"
)


type ControladorSimulacion struct {
    estacionamiento    *domain.Estacionamiento
    rnd                *rand.Rand
    log                []string
    mutex              sync.Mutex
    VehiculosEnEspera  []int 
}


func NuevoControlador(estacionamiento *domain.Estacionamiento, rnd *rand.Rand) *ControladorSimulacion {
    return &ControladorSimulacion{
        estacionamiento:   estacionamiento,
        rnd:               rnd,
        log:               make([]string, 0),
        VehiculosEnEspera: make([]int, 0), 
    }
}

func (c *ControladorSimulacion) IniciarSimulacion() {
    id := 1
    for {
        vehiculo := domain.GenerarVehiculo(id, c.rnd)
        fmt.Println("Generando vehículo:", vehiculo.ID) 
        go c.intentarEntrada(vehiculo)
        time.Sleep(time.Duration(c.rnd.ExpFloat64()) * time.Second) 
        id++
    }
}


func (c *ControladorSimulacion) intentarEntrada(vehiculo *domain.Vehiculo) {
    if c.estacionamiento.IntentarEntrar() {
        mensaje := fmt.Sprintf("Vehículo %d ha entrado", vehiculo.ID)
        fmt.Println(mensaje) 
        c.agregarLog(mensaje)
        time.Sleep(time.Duration(vehiculo.Tiempo) * time.Second)
        c.estacionamiento.Salir()
        mensaje = fmt.Sprintf("Vehículo %d ha salido", vehiculo.ID)
        fmt.Println(mensaje) 
        c.agregarLog(mensaje)
    } else {
        mensaje := fmt.Sprintf("Vehículo %d esperando para entrar", vehiculo.ID)
        fmt.Println(mensaje) 
        c.agregarLog(mensaje)
    }
}


func (c *ControladorSimulacion) EstacionamientoOcupado() int {
    ocupados := c.estacionamiento.Ocupados() 
    fmt.Println("Número de espacios ocupados:", ocupados) 
    return ocupados
}


func (c *ControladorSimulacion) agregarLog(mensaje string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.log = append(c.log, mensaje)
}

func (c *ControladorSimulacion) Registro() string {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    logString := ""
    for _, entry := range c.log {
        logString += entry + "\n"
    }
    return logString
}
