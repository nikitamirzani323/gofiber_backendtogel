package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiber_backendtogel/config"
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
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
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
