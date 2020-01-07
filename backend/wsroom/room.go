package wsroom

import "sync"

type Room struct {
	Name      string
	Members   *sync.Map
	stopChan  chan bool
	joinChan  chan *Conn
	leaveChan chan *Conn
	Send      chan *RoomMessage
}

var RoomManager = struct {
	Rooms *sync.Map
}{
	Rooms: new(sync.Map),
}
