package handlers

import (
	"log"

	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/models"
	"github.com/Wai-Thura-Tun/WebSocket-Using-Golang/internal/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchHandler struct{}

func NewMatchHandler() *MatchHandler {
	return &MatchHandler{}
}

func (h *MatchHandler) CreateMatch(c *fiber.Ctx) error {
	var match models.Match
	if err := c.BodyParser(&match); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	err := services.CreateMatch(match.User1ID.Hex(), match.User2ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create match.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Match created successfully",
	})
}

func (h *MatchHandler) GetMatches(c *fiber.Ctx) error {
	matches, err := services.GetMatches()
	userId := c.Locals("userId").(string)
	log.Printf("Log user id: %s", userId)

	var results []fiber.Map
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Match Not Found",
		})
	}

	for _, match := range matches {
		log.Printf("match id: %s", match.ID)
		var peerID primitive.ObjectID
		if match.User1ID.Hex() == userId {
			peerID = match.User2ID
		} else {
			peerID = match.User1ID
		}
		peer, err := services.GetUserByID(peerID.Hex())
		if err != nil {
			log.Println("fetching user by id:", err)
			continue
		}
		results = append(results, fiber.Map{
			"peer_id":   peer.ID.Hex(),
			"peer_name": peer.Username,
		})
	}
	return c.Status(fiber.StatusOK).JSON(results)
}
