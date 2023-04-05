package bot

//go:generate go-bindata -pkg bot -o assets.go assets/

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Bot struct {
	conn     *websocket.Conn
	handler  func(*Bot, []byte)
	outgoing chan []byte
}

func NewBot(handler func(bot *Bot, msg []byte)) (*Bot, error) {
	u := url.URL{Scheme: "wss", Host: "production.highrise.game", Path: "/web/webapi"}

	headers := http.Header{}
	headers.Set("api-token", getEnv("BOT_TOKEN", ""))
	headers.Set("room-id", getEnv("ROOM_ID", ""))

	rootCA, err := Asset("assets/gts1p5.pem")
	if err != nil {
		return nil, err
	}

	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)

	tlsConfig := &tls.Config{
		RootCAs:            rootCAPool,
		InsecureSkipVerify: false,
	}

	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = tlsConfig

	conn, _, err := dialer.Dial(u.String(), headers)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		conn:     conn,
		handler:  handler,
		outgoing: make(chan []byte, 100),
	}

	go bot.listen()
	go bot.timeOut()
	go bot.writeLoop()

	return bot, nil
}

func (b *Bot) Close() {
	close(b.outgoing)
	b.conn.Close()
}

func (b *Bot) listen() {
	for {
		messageType, message, err := b.conn.ReadMessage()
		if err != nil {
			// log.Printf("could not read message from WebSocket connection: %v", err)
			log.Fatalf("could not read message from WebSocket connection: %v", err)
		}

		switch messageType {
		case websocket.TextMessage:
			log.Printf("received text message: %v", message)
			go b.handler(b, message)
		case websocket.BinaryMessage:
			log.Printf("received binary message: %v", message)
			go b.handler(b, message)
		case websocket.CloseMessage:
			log.Fatalf("received close message")
		}
	}
}

func (b *Bot) timeOut() {
	for {
		b.conn.WriteMessage(websocket.TextMessage, []byte(`{"_type": "KeepaliveRequest"}`))
		time.Sleep(15 * time.Second)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if fallback == "" {
		panic("Missing env variable: " + key)
	}
	return fallback
}

func (b *Bot) writeLoop() {
	for message := range b.outgoing {
		if err := b.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("could not write message to WebSocket connection: %v", err)
		}
	}
}
