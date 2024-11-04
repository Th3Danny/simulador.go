package utils

import (
    "math"
    "math/rand"
    "time"
)

// Poisson distribuye los tiempos de llegada usando una distribución de Poisson
func Poisson(lambda float64) time.Duration {
    rand.Seed(time.Now().UnixNano())
    interval := -math.Log(1.0-rand.Float64()) / lambda
    return time.Duration(interval * float64(time.Second))
}
