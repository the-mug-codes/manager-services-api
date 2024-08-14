package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func GetWebSocketUpgrader(context *gin.Context) (conn *websocket.Conn, err error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(request *http.Request) bool {
			return true
		},
	}
	conn, err = upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return conn, err
	}
	return conn, err
}
