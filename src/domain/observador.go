package domain

type Observador interface {
    Actualizar(vehiculo *Vehiculo) 
    ActualizarEstadoEstacionamiento() 
}