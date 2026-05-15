package main
import "fmt"
func main() {// Go no tiene while. El for hace todo.
    // 1. For clásico (como en C/Java)
    for i := 0; i < 5; i++ {
    fmt.Println(i)
    }
    // 2. For como "while" (solo condición)
    contador := 0
    for contador < 3 {
    fmt.Println("Contando:", contador)
    contador++
    }
    // 3. For infinito (necesita break para salir)
    for {
    fmt.Println("Bucle infinito")
    break // Salimos inmediatamente
    }
}