package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/frontendagen_svelte/config"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_dashboard struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}

func Dashboard(c *fiber.Ctx) error {
	render_page := time.Now()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response_dashboard{}).
		SetHeader("Content-Type", "application/json").
		Post(config.Path_url() + "api/dashboardwinlose")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*response_dashboard)
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
