package domain

import(
	"sync"
	"fmt"
) 

type Estacionamiento struct {
    capacidad    int
    ocupados     int
    mutex        sync.Mutex
    observadores []Observador
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
    return &Estacionamiento{
        capacidad:    capacidad,
        ocupados:     0,
        observadores: make([]Observador, 0),
    }
}

func (e *Estacionamiento) AgregarObservador(o Observador) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    e.observadores = append(e.observadores, o)
}

func (e *Estacionamiento) NotificarObservadores() {
    for _, observador := range e.observadores {
        observador.ActualizarEstadoEstacionamiento()
    }
}

func (e *Estacionamiento) IntentarEntrar() bool {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados < e.capacidad {
        e.ocupados++
        fmt.Println("Vehículo entró, ocupados:", e.ocupados)
        e.NotificarObservadores()
        return true
    }
    return false
}

func (e *Estacionamiento) Salir() {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados > 0 {
        e.ocupados--
        fmt.Println("Vehículo salió, ocupados:", e.ocupados)
        e.NotificarObservadores()
    }
}



func (e *Estacionamiento) Ocupados() int {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    return e.ocupados
}
