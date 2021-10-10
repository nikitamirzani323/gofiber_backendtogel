package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiber_backendtogel/config"
	"github.com/nikitamirzani323/gofiber_backendtogel/helpers"
)

type pasarandetail struct {
	Idpasaran int `json:"idpasaran"`
}
type pasaransave struct {
	Idpasaran         int    `json:"idpasaran"`
	Page              string `json:"page"`
	Pasaran_diundi    string `json:"pasaran_diundi"`
	Pasaran_url       string `json:"pasaran_url"`
	Pasaran_jamtutup  string `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal string `json:"pasaran_jamjadwal"`
	Pasaran_jamopen   string `json:"pasaran_jamopen"`
	Pasaran_display   int    `json:"pasaran_display"`
	Pasaran_status    string `json:"pasaran_status"`
}
type pasaransaveonline struct {
	Idpasaran   int    `json:"idpasaran" validate:"required"`
	Page        string `json:"page"`
	Haripasaran string `json:"pasaran_hariraya" validate:"required,alpha"`
}
type pasarandeleteonline struct {
	Idpasaran      int    `json:"idpasaran" validate:"required"`
	Idpasaraonline int    `json:"idpasaraonline" validate:"required,numeric"`
	Page           string `json:"page"`
}
type pasaransavelimit struct {
	Idpasaran            int    `json:"idpasaran"`
	Page                 string `json:"page"`
	Pasaran_limitline4d  int    `json:"pasaran_limitline4d"`
	Pasaran_limitline3d  int    `json:"pasaran_limitline3d"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt"`
}
type response_pasaran struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Pasaranonline interface{} `json:"pasaranonline"`
}
type response_pasaransave struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Pasaran(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_adminrule{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/allpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_adminrule)
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
func Pasarandetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandetail)
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
		SetResult(response_pasaran{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran": client.Idpasaran,
		}).
		Post(config.Path_url() + "api/editpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasaran)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":        http.StatusOK,
			"message":       result.Message,
			"record":        result.Record,
			"pasaranonline": result.Pasaranonline,
			"time":          time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":        resp.StatusCode(),
			"message":       result.Message,
			"record":        result.Record,
			"pasaranonline": result.Pasaranonline,
			"time":          time.Since(render_page).String(),
		})
	}
}
func Pasaransave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransave)
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
		SetResult(response_pasaransave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran":         client.Idpasaran,
			"page":              client.Page,
			"pasaran_diundi":    client.Pasaran_diundi,
			"pasaran_url":       client.Pasaran_url,
			"pasaran_jamtutup":  client.Pasaran_jamtutup,
			"pasaran_jamjadwal": client.Pasaran_jamjadwal,
			"pasaran_jamopen":   client.Pasaran_jamopen,
			"pasaran_display":   client.Pasaran_display,
			"pasaran_status":    client.Pasaran_status,
		}).
		Post(config.Path_url() + "api/savepasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasaransave)
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
func Pasaransaveonline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveonline)
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
		SetResult(response_pasaransave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran":        client.Idpasaran,
			"page":             client.Page,
			"pasaran_hariraya": client.Haripasaran,
		}).
		Post(config.Path_url() + "api/savepasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_pasaransave)
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
func Pasarandeleteonline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandeleteonline)
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
		SetResult(response_pasaransave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran":      client.Idpasaran,
			"idpasaraonline": client.Idpasaraonline,
			"page":           client.Page,
		}).
		Post(config.Path_url() + "api/deletepasaranonline")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_pasaransave)
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
func Pasaransavelimit(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransavelimit)
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
		SetResult(response_pasaransave{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idpasaran":            client.Idpasaran,
			"page":                 client.Page,
			"pasaran_limitline4d":  client.Pasaran_limitline4d,
			"pasaran_limitline3d":  client.Pasaran_limitline3d,
			"pasaran_limitline2d":  client.Pasaran_limitline2d,
			"pasaran_limitline2dd": client.Pasaran_limitline2dd,
			"pasaran_limitline2dt": client.Pasaran_limitline2dt,
		}).
		Post(config.Path_url() + "api/savepasaranlimitline")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response_pasaransave)
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
