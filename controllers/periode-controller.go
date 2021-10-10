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
type periodesavenew struct {
	Sdata         string `json:"sData" validate:"required"`
	Page          string `json:"page"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type periodeprediksi struct {
	Nomorkeluaran string `json:"nomorkeluaran" validate:"required,min=4,max=4"`
	Idcomppasaran int    `json:"pasaran_code" validate:"required"`
}
type response_periode struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Pasaranonline interface{} `json:"pasaranonline"`
}
type response_periodedetail struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type response_periodelistbet struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Totalbet    int         `json:"totalbet"`
	Subtotal    int         `json:"subtotal"`
	Subtotalwin int         `json:"subtotalwin"`
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/editperiode")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":         client.Sdata,
			"idinvoice":     client.Idtrxkeluaran,
			"nomorkeluaran": client.Nomorkeluaran,
			"idpasarancode": client.Idpasarantogel,
		}).
		Post(config.Path_url() + "api/saveperiode")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_periodedetail)
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
func Periodesavenew(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodesavenew)
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"sData":        client.Sdata,
			"page":         client.Page,
			"pasaran_code": client.Idcomppasaran,
		}).
		Post(config.Path_url() + "api/savepasarannew")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_periodedetail)
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistmember")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
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
		SetResult(response_periodedetail{}).
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
	result := resp.Result().(*response_periodedetail)
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
		SetResult(response_periodelistbet{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodelistbet")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodelistbet)
	if result.Status == 200 {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":      http.StatusOK,
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
		})
	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":      resp.StatusCode(),
			"message":     result.Message,
			"record":      result.Record,
			"totalbet":    result.Totalbet,
			"subtotal":    result.Subtotal,
			"subtotalwin": result.Subtotalwin,
			"time":        time.Since(render_page).String(),
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idinvoice": client.Idtrxkeluaran,
		}).
		Post(config.Path_url() + "api/periodebettable")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
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
func Periodelistpasaran(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_periode{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/listpasaran")
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
func Periodeprediksi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(periodeprediksi)
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
		SetResult(response_periodedetail{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"nomorkeluaran": client.Nomorkeluaran,
			"Idcomppasaran": client.Idcomppasaran,
		}).
		Post(config.Path_url() + "api/listprediksi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response_periodedetail)
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
