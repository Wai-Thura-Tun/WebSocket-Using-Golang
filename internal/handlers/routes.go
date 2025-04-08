package handlers

import (
	"log"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/middlewares"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, g *ws.Gateway) {
	userHandler := NewUserHandler()
	authHandler := NewAuthHandler()
	matchHandler := NewMatchHandler()

	app.Post("/register", userHandler.CreateUser)
	app.Post("/login", authHandler.Login)
	app.Post("/match/create", matchHandler.CreateMatch)

	app.Use(middlewares.AuthMiddleware)

	app.Get("/matches", matchHandler.GetMatches)
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		userID := c.Locals("userId").(string)
		client := &ws.Client{
			Gateway: g,
			Conn:    c,
			Send:    make(chan ws.Message, 256),
			UserID:  userID,
		}
		g.Register(client)
		defer g.Unregister(client)
		log.Printf("Clients count: %v", len(g.Clients))

		// Send message to client throght websocket
		go func() {
			for msg := range client.Send {
				if err := c.WriteJSON(msg); err != nil {
					log.Printf("Write error: %v", err)
					return
				}
			}
		}()

		// Get message from client and publish to receiver channel
		var msg ws.Message
		for {
			if err := c.ReadJSON(&msg); err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			msg.SenderID = userID
			g.HandleMessage(msg)
		}
	}))
}
