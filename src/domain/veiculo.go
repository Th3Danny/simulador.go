package domain

import "math/rand"

type Vehiculo struct {
    ID     int // ID único del vehículo
    Tiempo int // Tiempo en el que el vehículo permanecerá en el estacionamiento
}

func GenerarVehiculo(id int, rnd *rand.Rand) *Vehiculo {
    // Genera un tiempo de estacionamiento aleatorio entre 3 y 5 segundos
    tiempo := rnd.Intn(3) + 3
    vehiculo := &Vehiculo{
        ID:     id,
        Tiempo: tiempo,
    }
    // Aquí deberías tener acceso a una instancia de `Estacionamiento` para notificar, por ejemplo:
    // estacionamiento.NotificarObservadores(vehiculo) (según tu lógica)
    return vehiculo
}
