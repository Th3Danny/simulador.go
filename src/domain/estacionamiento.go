package domain

import "sync"

type Estacionamiento struct {
    capacidad int
    ocupados  int
    mutex     sync.Mutex
}

// NuevoEstacionamiento crea una nueva instancia del estacionamiento.
func NuevoEstacionamiento(capacidad int) *Estacionamiento {
    return &Estacionamiento{
        capacidad: capacidad,
        ocupados:  0,
    }
}

// IntentarEntrar intenta colocar un vehículo en el estacionamiento.
func (e *Estacionamiento) IntentarEntrar(vehiculoID int) bool {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados < e.capacidad {
        e.ocupados++
        return true
    }
    return false
}

// Salir elimina un vehículo del estacionamiento.
func (e *Estacionamiento) Salir(vehiculoID int) {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    if e.ocupados > 0 {
        e.ocupados--
    }
}

// ObtenerEspacios devuelve un slice con la ocupación de los espacios
func (e *Estacionamiento) ObtenerEspacios() []int {
    e.mutex.Lock()
    defer e.mutex.Unlock()
    
    ocupacion := make([]int, e.capacidad)
    for i := 0; i < e.capacidad; i++ {
        if e.ocupados > i {
            ocupacion[i] = 1 // 1 significa que está ocupado
        } else {
            ocupacion[i] = 0 // 0 significa que está libre
        }
    }
    return ocupacion
}
