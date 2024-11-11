package domain

type Observador interface {
    Actualizar(vehiculo *Vehiculo) // Método requerido
    ActualizarEstadoEstacionamiento() // Método adicional (opcional según tu implementación)
}