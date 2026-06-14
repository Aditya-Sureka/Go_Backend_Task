package handler

import (
	"github.com/Aditya-Sureka/Go_Backend_Task/internal/models"
	"github.com/Aditya-Sureka/Go_Backend_Task/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service  *service.UserService
	Validate *validator.Validate
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		Service:  service,
		Validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(
	c *fiber.Ctx,
) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "invalid request body",
			},
		)
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.Service.CreateUser(
		c.Context(),
		req.Name,
		req.DOB,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
		},
	)
}