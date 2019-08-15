package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/andreylm/basic-go-server.git/pkg/server/modules/v1/chat/models"

	"github.com/gorilla/websocket"

	"github.com/andreylm/basic-go-server.git/pkg/db"
)

var (
	upgrager = websocket.Upgrader{
		ReadBufferSize:  1 << 10,
		WriteBufferSize: 1 << 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	once sync.Once
	hub  *models.Hub
)

// Chat - chat action
func Chat(db db.DB) http.HandlerFunc {
	once.Do(func() {
		log.Println("Running hub...")
		hub = models.NewHub()
		go hub.Run()
	})
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Creating connection....")
		conn, err := upgrager.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &models.Client{Hub: hub, Conn: conn, Send: make(chan models.ChatMessage)}
		client.Hub.Register <- client
		done := make(chan struct{})

		go client.Write(done)
		go client.Read(done)

		for {
			select {
			case <-done:
				log.Println("Exit..")
				return
			}
		}
	}
}
