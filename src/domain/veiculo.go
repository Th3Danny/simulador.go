package domain

import "math/rand"

type Vehiculo struct {
    ID     int 
    Tiempo int 
}

func GenerarVehiculo(id int, rnd *rand.Rand) *Vehiculo {
    // Genera un tiempo de estacionamiento aleatorio entre 3 y 5 segundos
    tiempo := rnd.Intn(3) + 3
    vehiculo := &Vehiculo{
        ID:     id,
        Tiempo: tiempo,
    }
    
    return vehiculo
}
