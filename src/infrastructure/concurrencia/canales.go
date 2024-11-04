package concurrencia

import "sync"

// CanalSincronizado gestiona el acceso concurrente a través de un canal y un WaitGroup
type CanalSincronizado struct {
    ch  chan struct{}
    wg  sync.WaitGroup
}

// NuevoCanalSincronizado crea un canal de sincronización con el tamaño especificado
func NuevoCanalSincronizado(tamano int) *CanalSincronizado {
    return &CanalSincronizado{
        ch: make(chan struct{}, tamano),
    }
}

// Enviar agrega una tarea a la sincronización
func (c *CanalSincronizado) Enviar() {
    c.wg.Add(1)
    c.ch <- struct{}{}
}

// Recibir completa una tarea de la sincronización
func (c *CanalSincronizado) Recibir() {
    <-c.ch
    c.wg.Done()
}

// Esperar espera a que todas las tareas se completen
func (c *CanalSincronizado) Esperar() {
    c.wg.Wait()
}
