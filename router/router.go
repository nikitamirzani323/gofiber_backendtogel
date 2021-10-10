package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/gofiber_backendtogel/controllers"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "svelte/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Post("/api/login", controllers.Login)
	app.Post("/api/home", controllers.Home)
	app.Post("/api/dashboardwinlose", controllers.Dashboard)

	app.Post("/api/allpasaran", controllers.Pasaran)
	app.Post("/api/editpasaran", controllers.Pasarandetail)

	app.Post("/api/alladmin", controllers.Admin)
	app.Post("/api/editadmin", controllers.Admindetail)
	app.Post("/api/saveadmin", controllers.Adminsave)
	app.Post("/api/saveadminiplist", controllers.Adminsaveiplist)

	app.Post("/api/alladminrule", controllers.Adminrule)
	app.Post("/api/editadminrule", controllers.Adminruledetail)
	app.Post("/api/saveadminrule", controllers.Adminrulesave)
	app.Post("/api/saveadminruleconf", controllers.Adminruleconfsave)
	return app
}
