package modules

import "gopkg.in/olahol/melody.v1"

var WebSocketModule *melody.Melody

func InitWebSocketModule() {
	WebSocketModule = melody.New()
}
