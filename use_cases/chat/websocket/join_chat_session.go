package chat_websocket

import (
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	websocket "github.com/the-mug-codes/service-manager-api/adapters/websocket"
	entity "github.com/the-mug-codes/service-manager-api/entities"
)

func JoinChatSection(chatMessage entity.ChatMessageRepository, websocketChat *websocket.WebsocketChat, wsConnection *ws.Conn, sessionID uuid.UUID) (err error) {
	client := websocket.CreateChatClientConnection(wsConnection, sessionID, make(chan *entity.ChatMessage))
	websocketChat.Sessions.Register <- client
	go client.WriteMessage(chatMessage)
	client.ReadMessage(websocketChat.Sessions)
	return err
}
