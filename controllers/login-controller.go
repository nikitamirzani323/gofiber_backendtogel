package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiber_backendtogel/config"
	"github.com/nikitamirzani323/gofiber_backendtogel/helpers"
)

type login struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Ipaddress string `json:"ipaddress" validate:"required"`
	Timezone  string `json:"timezone" validate:"required"`
}
type home struct {
	Page string `json:"page" validate:"required"`
}
type response_login struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}
type response_home struct {
	Message string `json:"message"`
}

func Login(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(login)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_login{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username":  client.Username,
			"password":  client.Password,
			"ipaddress": client.Ipaddress,
			"timezone":  client.Timezone,
		}).
		Post(config.Path_url() + "api/login")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_login)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"token":  result.Token,
		"key":    result.Key,
		"time":   time.Since(render_page).String(),
	})
}
func Home(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(home)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_home{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"page": client.Page,
		}).
		Post(config.Path_url() + "api/home")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_home)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": result.Message,
		"time":    time.Since(render_page).String(),
	})
}
