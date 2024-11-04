package domain

import "math/rand"

type Vehiculo struct {
    ID     int
    Tiempo int
}

func GenerarVehiculo(id int, rnd *rand.Rand) *Vehiculo {
    tiempo := rnd.Intn(3) + 3 // Aleatorio entre 3 y 5 segundos
    return &Vehiculo{
        ID:     id,
        Tiempo: tiempo,
    }
}
