package wsroom

import (
	"net/http"
	"sync"
	"time"

	"github.com/chuckpreslar/emission"
	"github.com/gorilla/websocket"
)

type Conn struct {
	Cookie *sync.Map
	Socket *websocket.Conn
	ID     string
	Send   chan []byte
	Rooms  *sync.Map
}

type CookieReader func(*http.Request) map[string]string

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = pongWait * 9 / 10
	maxMessageSize = 1024 * 1024 * 1024
)

var (
	ConnManager = struct {
		Conns *sync.Map
	}{
		Conns: new(sync.Map),
	}
	Emitter = emission.NewEmitter()
)

func (c *Conn) readPump() {
	defer func() {
		c.Rooms.Range(func(key, value interface{}) bool {
			room, ok := RoomManager.Rooms.Load(key.(string))
			if ok {
				room.(*Room).Leave(c)
			}
			return true
		})
		c.Socket.Close()
	}()
}
