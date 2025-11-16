package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var quotesArray = []string{
	"La vida es bella",
	"El camino es el destino",
	"El obstáculo es el camino",
	"Menos es más",
	"El tiempo es oro",
	"La acción es la clave",
	"Piensa diferente",
	"Cree en ti",
	"Sé el cambio",
	"Vive el presente",
	"Ama sin límites",
	"Sigue tu pasión",
	"Crea tu suerte",
	"Hoy es el día",
	"Nunca te rindas",
	"Sueña grande",
	"La calma es poder",
	"Sé auténtico",
	"La fe mueve montañas",
	"Aprende del error",
	"La paciencia es virtud",
	"Busca la paz",
	"El silencio enseña",
	"Sé tu prioridad",
	"Respira profundo",
	"El arte sana",
	"Perdona rápido",
	"Agradece siempre",
	"Todo es posible",
	"La verdad libera",
	"Ríe a menudo",
	"Valora lo simple",
	"La duda paraliza",
	"El miedo es temporal",
	"Hazlo ahora",
	"Disfruta el viaje",
	"Sé flexible",
	"El éxito es tuyo",
	"Comienza de nuevo",
	"Sé impecable",
	"El esfuerzo vale",
	"Mira hacia adelante",
	"Escucha tu voz",
	"La meta es clara",
	"Supera tus límites",
	"La bondad vuelve",
	"Decide hoy",
	"Sé valiente",
	"La energía fluye",
	"Acepta el reto",
	"La música une",
	"El sol sale",
	"Deja ir",
	"Siembra y cosecha",
	"Tu salud primero",
	"La magia existe",
	"Encuentra la luz",
	"La meta es el inicio",
	"Conoce tu valor",
	"El futuro espera",
	"Haz tu parte",
	"La fuerza interior",
	"Cambia la perspectiva",
	"El enfoque gana",
	"Caminante, no hay camino",
	"La soledad nutre",
	"Sé generoso",
	"El amor es todo",
	"Da lo mejor",
	"El detalle importa",
	"La quietud ayuda",
	"Observa el mundo",
	"El presente es regalo",
	"La esperanza vive",
	"Sé persistente",
	"El poder reside en ti",
	"Céntrate en lo bueno",
	"La mente es clave",
	"Trabaja duro",
	"Sé sabio",
	"La aventura espera",
	"El propósito guía",
	"La libertad interior",
	"Abraza el cambio",
	"El progreso es lento",
	"La belleza está dentro",
	"No te compares",
	"Tu paz es sagrada",
	"Actúa con intención",
	"La resiliencia crece",
	"Busca el conocimiento",
	"Sé responsable",
	"El equilibrio es arte",
	"Confía en el proceso",
	"La creatividad fluye",
	"Mejora continua",
	"Sé agradecido siempre",
	"El respeto es base",
}

var generatedQuotes = []string{}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/api/quote", generateQuoteHandler)
	http.HandleFunc("/api/quotes", getQuotesHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateQuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	numeroAleatorio := generarNumeroAleatorio()
	quote := quotesArray[numeroAleatorio]
	generatedQuotes = append(generatedQuotes, quote)

	response := map[string]interface{}{
		"quote": quote,
		"total": len(generatedQuotes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getQuotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"quotes": generatedQuotes,
		"total":  len(generatedQuotes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generarNumeroAleatorio() int {
	return rand.Intn(len(quotesArray))
}
