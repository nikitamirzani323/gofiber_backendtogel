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

type periodedetail struct {
	Idtrxkeluaran int `json:"idinvoice"`
}
type periodedetailmembernomor struct {
	Idtrxkeluaran int    `json:"idinvoice" validate:"required"`
	Permainan     string `json:"permainan" validate:"required"`
	Nomor         string `json:"nomor" validate:"required"`
}
type periodeSave struct {
	Sdata          string `json:"sData" validate:"required"`
	Page           string `json:"page"`
	Idtrxkeluaran  int    `json:"idinvoice" validate:"required"`
	Nomorkeluaran  string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idpasarantogel string `json:"idpasarancode" validate:"required"`
}
type periodeSaveNew struct {
	Sdata         string `json:"sData" validate:"required"`
	Page          string `json:"page"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type periodePrediksi struct {
	Nomorkeluaran string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type response_periode struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Pasaranonline interface{} `json:"pasaranonline"`
}

func Periode(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_periode{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/allperiode")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periode)
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
func Periodedetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/editperiode")
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
func Periodesave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeSave)
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
			"idinvoice":     client.Idtrxkeluaran,
			"nomorkeluaran": client.Nomorkeluaran,
			"idpasarancode": client.Idpasarantogel,
		}).
		Post(config.Path_url() + "api/saveperiode")
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
func Periodelistmember(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistmember")
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
func Periodelistmemberbynomor(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetailmembernomor)
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
			"idinvoice": client.Idtrxkeluaran,
			"permainan": client.Permainan,
			"nomor":     client.Nomor,
		}).
		Post(config.Path_url() + "api/periodelistmemberbynomor")
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
func Periodelistbet(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistbet")
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
func Periodebettable(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodedetail)
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
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodebettable")
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
