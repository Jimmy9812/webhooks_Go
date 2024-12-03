package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookPayload struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método sea POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializar el JSON recibido
	var payload WebhookPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	// Procesar el evento recibido
	fmt.Printf("Evento recibido: %s\n", payload.Event)
	fmt.Printf("Datos: %s\n", payload.Data)

	// Responder al emisor del Webhook
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook recibido correctamente"))
}

func main() {
	// Crear una ruta para el Webhook
	http.HandleFunc("/webhook", webhookHandler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
