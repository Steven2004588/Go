package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var opciones = []string{"piedra", "papel", "tijera"}

func elegirComputador() string {
	rand.Seed(time.Now().UnixNano())
	return opciones[rand.Intn(3)]
}

func decidirGanador(jugador, computador string) string {
	if jugador == computador {
		return "empate"
	}
	gana := map[string]string{
		"piedra": "tijera",
		"tijera": "papel",
		"papel":  "piedra",
	}
	if gana[jugador] == computador {
		return "ganaste"
	}
	return "perdiste"
}

func manejarInicio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func manejarJugar(w http.ResponseWriter, r *http.Request) {
	jugador := r.FormValue("jugador")

	if jugador != "piedra" && jugador != "papel" && jugador != "tijera" {
		fmt.Fprintf(w, "Opción inválida")
		return
	}

	computador := elegirComputador()
	resultado := decidirGanador(jugador, computador)

	emojis := map[string]string{
		"piedra": "🪨",
		"papel":  "📄",
		"tijera": "✂️",
	}
	mensajes := map[string]string{
		"ganaste":  "¡Ganaste! 🏆",
		"perdiste": "Perdiste... 💀",
		"empate":   "¡Empate! 🤝",
	}
	colores := map[string]string{
		"ganaste":  "#4ade80",
		"perdiste": "#e84545",
		"empate":   "#facc15",
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>Resultado</title>
  <style>
    @import url('https://fonts.googleapis.com/css2?family=Bebas+Neue&family=DM+Sans:wght@400;600&display=swap');
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: 'DM Sans', sans-serif;
      background: #0d0d0f;
      color: #eee;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
    }
    .tarjeta {
      background: #18181c;
      border: 1px solid #2a2a30;
      border-radius: 20px;
      padding: 2.5rem 3rem;
      text-align: center;
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      box-shadow: 0 0 40px rgba(0,0,0,0.6);
      animation: aparecer 0.4s ease;
    }
    @keyframes aparecer {
      from { opacity: 0; transform: translateY(20px); }
      to   { opacity: 1; transform: translateY(0); }
    }
    .versus {
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 2.5rem;
    }
    .jugador span { font-size: 3.5rem; display: block; }
    .vs {
      font-family: 'Bebas Neue', sans-serif;
      font-size: 2rem;
      color: #e84545;
    }
    .label {
      font-size: 0.7rem;
      color: #555;
      letter-spacing: 1px;
      text-transform: uppercase;
      margin-bottom: 4px;
    }
    .resultado {
      font-family: 'Bebas Neue', sans-serif;
      font-size: 2.5rem;
      color: %s;
    }
    .boton {
      padding: 0.75rem 2rem;
      background: #e84545;
      color: white;
      border: none;
      border-radius: 100px;
      font-size: 1rem;
      font-weight: 600;
      cursor: pointer;
      text-decoration: none;
      display: inline-block;
      transition: background 0.2s;
    }
    .boton:hover { background: #c73232; }
  </style>
</head>
<body>
  <div class="tarjeta">
    <div class="versus">
      <div class="jugador">
        <div class="label">Tú</div>
        <span>%s</span>
      </div>
      <div class="vs">VS</div>
      <div class="jugador">
        <div class="label">PC</div>
        <span>%s</span>
      </div>
    </div>
    <div class="resultado">%s</div>
    <a class="boton" href="/">Jugar de nuevo</a>
  </div>
</body>
</html>`,
		colores[resultado],
		emojis[jugador],
		emojis[computador],
		mensajes[resultado],
	)
}

func main() {
	http.HandleFunc("/", manejarInicio)
	http.HandleFunc("/jugar", manejarJugar)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}