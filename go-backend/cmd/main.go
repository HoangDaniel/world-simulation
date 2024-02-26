package main

import (
	"encoding/json"
	"hoangdaniel/world/cmd/timeservice"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	clients = make(map[*websocket.Conn]bool)
	mutex   = &sync.Mutex{}
)

// Handler for WebSocket requests
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register new client
	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	// Keep connection alive
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}
	}
}

func broadcastTime(ts *timeservice.TimeService) {
	for {
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(ts.GetCurrentTime())
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	startTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	ts := timeservice.NewTimeService(startTime)
	ts.SetSpeed(60)
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	http.HandleFunc("/ws", handleConnections)
	go broadcastTime(ts)

	http.HandleFunc("/setSpeed", func(w http.ResponseWriter, r *http.Request) {
		var speed struct {
			Speed float64 `json:"speed"`
		}
		if err := json.NewDecoder(r.Body).Decode(&speed); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		ts.SetSpeed(speed.Speed)
	})

	http.HandleFunc("/pause", func(w http.ResponseWriter, r *http.Request) {
		ts.Pause()
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		ts.Resume()
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", corsOpts(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
