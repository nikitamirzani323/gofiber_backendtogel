package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/frontendagen_svelte/config"
	"bitbucket.org/isbtotogroup/frontendagen_svelte/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type admindetail struct {
	Username string `json:"username" validate:"required"`
}
type adminsave struct {
	Sdata       string `json:"sdata" validate:"required"`
	Page        string `json:"page"`
	Idruleadmin int    `json:"idruleadmin" `
	Username    string `json:"username" validate:"required,alphanum,max=20"`
	Password    string `json:"password" `
	Name        string `json:"nama" validate:"required,alphanum,max=70"`
	Status      string `json:"status" validate:"required,alpha"`
}
type adminsaveiplist struct {
	Sdata     string `json:"sdata" validate:"required"`
	Page      string `json:"page"`
	Username  string `json:"username" validate:"required,alphanum,max=20"`
	Ipaddress string `json:"ipaddress" validate:"required,max=20"`
}
type response_admin struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Listruleadmin interface{} `json:"listruleadmin"`
	Listip        interface{} `json:"listip"`
}
type response_adminsave struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Admin(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_admin{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/alladmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_admin)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":        resp.StatusCode(),
			"message":       result.Message,
			"record":        result.Record,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	}
}
func Admindetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(admindetail)
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
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_admin{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": client.Username,
		}).
		Post(config.Path_url() + "api/editadmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_admin)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"listip":        result.Listip,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":        resp.StatusCode(),
			"message":       result.Message,
			"record":        result.Record,
			"listip":        result.Listip,
			"listruleadmin": result.Listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	}
}
func Adminsave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminsave)
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
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": client.Username,
		}).
		Post(config.Path_url() + "api/editadmin")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  resp.StatusCode(),
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Adminsaveiplist(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(adminsaveiplist)
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
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_adminsave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sdata":     client.Sdata,
			"page":      client.Page,
			"username":  client.Username,
			"ipaddress": client.Ipaddress,
		}).
		Post(config.Path_url() + "api/saveadminiplist")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminsave)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  resp.StatusCode(),
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	}
}
