package main
import "fmt"
func main() {
// Declaración larga: especificas el tipo tú mismo.
var nombre string = "Gopher"
var edad int = 25
// Declaración corta: Go infiere el tipo automáticamente.
// El := solo se puede usar DENTRO de una función.
puntos := 100
activo := true
// %s = texto, %d = número entero, %v = cualquier tipo
fmt.Printf("Jugador: %s, Edad: %d\n", nombre, edad)
fmt.Printf("Puntos: %d, Activo: %v\n", puntos, activo)
}