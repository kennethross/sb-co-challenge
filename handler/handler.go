package handler

import (
	"fmt"
	"math/rand"
	"snake-backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type State struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
}

type Fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"` // X velocity of the snake (one of -1, 0, 1)
	VelY int `json:"velY"` // Y velocity of the snake (one of -1, 0, 1)
}

func GetNewHandler(c *fiber.Ctx) error {
	widthQuery := c.Query("width")
	heightQuery := c.Query("height")

	width, err := strconv.Atoi(widthQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Width must be a positive number", "data": err})
	}
	height, err := strconv.Atoi(heightQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Height must be a positive number", "data": err})
	}

	state := State{
		GameID: uuid.New().String(),
		Width:  width,
		Height: height,
		Score:  0,
		Fruit: Fruit{
			X: rand.Intn(width),
			Y: rand.Intn(height),
		},
		Snake: Snake{
			X:    0,
			Y:    0,
			VelX: 1,
			VelY: 0,
		},
	}

	return c.JSON(state)
}

func PostValidatorHandler(c *fiber.Ctx) error {
	type reqBody struct {
		State State        `json="state"`
		Ticks []utils.Tick `json="ticks"`
	}

	var body reqBody

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	ticks := body.Ticks
	state := body.State

	currPosition := &utils.CurrentPosition{
		X: body.State.Snake.VelX,
		Y: body.State.Snake.VelY,
	}

	for index, tick := range ticks {
		fmt.Println(tick)
		if index > 0 {
			prevTick := ticks[index-1]
			if !utils.ValidateMove(prevTick, tick) {
				return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{"status": fiber.StatusMethodNotAllowed, "message": "Invalid move"})
			}
		}
		currPosition.NewPosition(tick)

		if !utils.ValidateInBound(*currPosition, body.State.Width, body.State.Height) {
			return c.Status(fiber.StatusTeapot).JSON(fiber.Map{"status": fiber.StatusTeapot, "message": "Snake out of bound"})
		}
	}

	fmt.Println(currPosition)

	if currPosition.X == state.Fruit.X && currPosition.Y == state.Fruit.Y {
		newState := State{
			GameID: state.GameID,
			Width:  state.Width,
			Height: state.Height,
			Score:  state.Score + 1,
			Fruit: Fruit{
				X: rand.Intn(state.Height),
				Y: rand.Intn(state.Width),
			},
			Snake: body.State.Snake,
		}
		return c.Status(fiber.StatusOK).JSON(newState)
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": " Fruit not found,"})
}
