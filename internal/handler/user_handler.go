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
		c.UserContext(),
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

func (h *UserHandler) GetUser(
c *fiber.Ctx,
) error {


id, err := c.ParamsInt("id")

if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"error": "invalid user id",
		},
	)
}

user, err := h.Service.GetUser(
	c.UserContext(),
	int32(id),
)

if err != nil {
	return c.Status(fiber.StatusNotFound).JSON(
		fiber.Map{
			"error": "user not found",
		},
	)
}

return c.JSON(
	fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  service.CalculateAge(user.Dob),
	},
)


}

func (h *UserHandler) ListUsers(
c *fiber.Ctx,
) error {


users, err := h.Service.ListUsers(
	c.UserContext(),
)

if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(
		fiber.Map{
			"error": err.Error(),
		},
	)
}

response := make([]fiber.Map, 0)

for _, user := range users {

	response = append(
		response,
		fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
			"age":  service.CalculateAge(user.Dob),
		},
	)
}

return c.JSON(response)

}

func (h *UserHandler) UpdateUser(
c *fiber.Ctx,
) error {


id, err := c.ParamsInt("id")

if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"error": "invalid user id",
		},
	)
}

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

user, err := h.Service.UpdateUser(
	c.UserContext(),
	int32(id),
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

return c.JSON(
	fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	},
)

}

func (h *UserHandler) DeleteUser(
c *fiber.Ctx,
) error {


id, err := c.ParamsInt("id")

if err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(
		fiber.Map{
			"error": "invalid user id",
		},
	)
}

err = h.Service.DeleteUser(
	c.UserContext(),
	int32(id),
)

if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(
		fiber.Map{
			"error": err.Error(),
		},
	)
}

return c.SendStatus(
	fiber.StatusNoContent,
)


}



