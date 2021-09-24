package main
import (
    "fmt"
    "math"
)
func MySqrt(x float64) float64 {
    z := 1.0
    var t float64
    for {
        z, t = z - (z*z - x) / (2*z), z
        if math.Abs(t-z) < 1e-6 {
            break
        }
    }
    return z
}
func main() {
    assumption := MySqrt(2)
    exact := math.Sqrt(2)
    fmt.Printf("Assumption: %v, Exact: %v, Error: %v", assumption, exact, math.Abs(assumption - exact))
}