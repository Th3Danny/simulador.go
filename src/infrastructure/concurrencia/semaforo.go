package concurrencia

type Semaforo struct {
    ch chan struct{}
}

// NuevoSemaforo crea un nuevo semáforo con el número de permisos especificado
func NuevoSemaforo(permisos int) *Semaforo {
    return &Semaforo{
        ch: make(chan struct{}, permisos),
    }
}

// Adquirir intenta obtener un permiso del semáforo
func (s *Semaforo) Adquirir() {
    s.ch <- struct{}{}
}

// Liberar libera un permiso en el semáforo
func (s *Semaforo) Liberar() {
    <-s.ch
}
