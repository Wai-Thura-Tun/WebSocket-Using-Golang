package ws

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/config"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gateway struct {
	Clients map[string]map[*Client]bool
	Redis   *redis.Client
	PubSub  *redis.PubSub
	Mu      sync.RWMutex
}

type Client struct {
	Gateway *Gateway
	Conn    *websocket.Conn
	Send    chan Message
	UserID  string
}

type Message struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	SenderID       string `json:"sender_id" bson:"sender_id"`
	ReceiverID     string `json:"receiver_id" bson:"receiver_id"`
	Content        string `json:"content" bson:"content"`
	MediaURL       string `json:"media_url,omitempty" bson:"media_url,omitempty"`
	Type           string `json:"type" bson:"type"`
	Timestamp      int64  `json:"timestamp" bson:"timestamp"`
	ConversationID string `json:"conversation_id" bson:"conversation_id"`
}

func NewGateWay() *Gateway {
	return &Gateway{
		Clients: make(map[string]map[*Client]bool),
		Redis:   config.RedisClient,
		PubSub:  config.RedisClient.Subscribe(context.Background(), "chat:*"),
	}
}

func (g *Gateway) Run() {
	channel := g.PubSub.Channel()
	for msg := range channel {
		var m Message
		if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
			log.Printf("Unmarshal error: %v", err)
			continue
		}
		g.Mu.RLock()
		if devices, ok := g.Clients[m.ReceiverID]; ok {
			for client := range devices {
				select {
				case client.Send <- m:
				default:
					close(client.Send)
					g.Unregister(client)
				}
			}
		}
		g.Mu.RUnlock()
	}
}

func (g *Gateway) HandleMessage(m Message) {
	if m.ID == "" {
		m.ID = primitive.NewObjectID().Hex()
	}

	if m.Timestamp == 0 {
		m.Timestamp = time.Now().Unix()
	}

	if m.ConversationID == "" {
		if m.SenderID < m.ReceiverID {
			m.ConversationID = m.SenderID + "_" + m.ReceiverID
		} else {
			m.ConversationID = m.ReceiverID + "_" + m.SenderID
		}
	}

	// Publish to redis
	payload, _ := json.Marshal(m)
	if err := g.Redis.Publish(context.Background(), "chat:"+m.ReceiverID, payload).Err(); err != nil {
		log.Printf("Publish error: %v", err)
	}

	// Save to mongo
	if _, err := config.MessageCollection.InsertOne(context.Background(), m); err != nil {
		log.Printf("MongoDB insert error: %v", err)
	}
}

func (g *Gateway) Register(c *Client) {
	g.Mu.Lock()
	if _, ok := g.Clients[c.UserID]; !ok {
		g.Clients[c.UserID] = make(map[*Client]bool)
	}
	g.Clients[c.UserID][c] = true
	g.Mu.Unlock()
	g.PubSub.Subscribe(context.Background(), "chat:"+c.UserID)
}

func (g *Gateway) Unregister(c *Client) {
	g.Mu.Lock()
	if devices, ok := g.Clients[c.UserID]; ok {
		if _, ok := devices[c]; ok {
			close(c.Send)
			delete(devices, c)
			if len(devices) == 0 {
				delete(g.Clients, c.UserID)
				g.PubSub.Unsubscribe(context.Background(), "chat:"+c.UserID)
			}
		}
	}
	g.Mu.Unlock()
}
