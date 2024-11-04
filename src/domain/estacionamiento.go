package domain

import "sync"

type Estacionamiento struct {
    capacidad int
    ocupados  int
    mutex     sync.Mutex
}


func NuevoEstacionamiento(capacidad int) *Estacionamiento {
    return &Estacionamiento{
        capacidad: capacidad,
        ocupados:  0,
    }
}

// Intentar entrar al estacionamiento
func (e *Estacionamiento) IntentarEntrar() bool {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados < e.capacidad {
        e.ocupados++
        return true
    }
    return false
}


func (e *Estacionamiento) Salir() {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados > 0 {
        e.ocupados--
    }
}

// Ocupados devuelve el número actual de espacios ocupados
func (e *Estacionamiento) Ocupados() int {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    return e.ocupados
}
