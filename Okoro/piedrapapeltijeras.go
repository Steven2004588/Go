package main
import (
    "fmt"
    "math/rand"
    "strings" // Para manejar texto
    "time"
)
func main() {
    // Semilla para que el azar cambie siempre
    rand.Seed(time.Now().UnixNano())
    // Opciones guardadas en una lista (arreglo)
    opciones := [3]string{"piedra", "papel", "tijera"}
    fmt.Println("¡Bienvenido a Piedra, Papel o Tijera!")
    for {
        fmt.Print("\nTu jugada (o escribe 'salir'): ")
        var usuario string
        fmt.Scan(&usuario)

        // Convertir a minúsculas por si escriben "PIEDRA"
        usuario = strings.ToLower(usuario)

        if usuario == "salir" {
            break // Termina el juego
        }
        // PC elige un número del 0 al 2 y lo traduce a palabra
        pc := opciones[rand.Intn(3)]
        fmt.Printf("PC eligió: %s\n", pc)

        // Comparar quién gana
        if usuario == pc {
            fmt.Println("Empate! 🤝")
        } else if (usuario == "piedra" && pc == "tijera") ||
            (usuario == "papel" && pc == "piedra") ||
            (usuario == "tijera" && pc == "papel") {
            fmt.Println("Ganaste! 🎉")
        } else {
            fmt.Println("Perdiste... 😰")
        }
    }
}