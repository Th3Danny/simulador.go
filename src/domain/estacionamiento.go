package domain

import "sync"

type Estacionamiento struct {
    capacidad    int
    ocupados     int
    mutex        sync.Mutex
    espacioLibre chan struct{}
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
    return &Estacionamiento{
        capacidad:    capacidad,
        ocupados:     0,
        mutex:        sync.Mutex{},
        espacioLibre: make(chan struct{}, capacidad),
    }
}

func (e *Estacionamiento) IntentarEntrar() bool {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados < e.capacidad {
        e.ocupados++
        e.espacioLibre <- struct{}{}
        return true
    }
    return false
}

func (e *Estacionamiento) Salir() {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados > 0 {
        e.ocupados--
        <-e.espacioLibre
    }
}

// Ocupados devuelve el número actual de espacios ocupados
func (e *Estacionamiento) Ocupados() int {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    return e.ocupados
}
