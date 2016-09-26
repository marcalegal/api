package socket

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
)

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: Index(db),
		},
	}
}

// Index ...
func Index(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// socket io
		server, err := socketio.NewServer(nil)
		if err != nil {
			log.Fatal(err)
		}
		server.On("connection", func(so socketio.Socket) {
			log.Println("on connection")
			so.Join("chat")
			so.On("chat message", func(msg string) {
				log.Println("emit:", so.Emit("chat message", msg))
				so.BroadcastTo("chat", "chat message", msg)
			})
			so.On("disconnection", func() {
				log.Println("on disconnect")
			})
		})
		server.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
		})

		server.ServeHTTP(w, r)
	}
}
