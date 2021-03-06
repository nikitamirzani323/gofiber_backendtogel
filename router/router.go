package router

import (
	"bitbucket.org/isbtotogroup/frontendagen_svelte/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Init() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(etag.New())
	app.Static("/", "svelte/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controllers.HealthCheck)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/home", controllers.Home)
	app.Post("/api/dashboardwinlose", controllers.Dashboard)
	app.Post("/api/reportwinlose", controllers.Reportwinlose)

	app.Post("/api/allpasaran", controllers.Pasaran)
	app.Post("/api/editpasaran", controllers.Pasarandetail)
	app.Post("/api/savepasaran", controllers.Pasaransave)
	app.Post("/api/savepasaranonline", controllers.Pasaransaveonline)
	app.Post("/api/deletepasaranonline", controllers.Pasarandeleteonline)
	app.Post("/api/savepasaranlimitline", controllers.Pasaransavelimit)
	app.Post("/api/savepasaranconf432", controllers.Pasaransaveconf432d)
	app.Post("/api/savepasaranconfcolokbebas", controllers.Pasaransaveconfcbebas)
	app.Post("/api/savepasaranconfcolokmacau", controllers.Pasaransaveconfcmacau)
	app.Post("/api/savepasaranconfcoloknaga", controllers.Pasaransaveconfcnaga)
	app.Post("/api/savepasaranconfcolokjitu", controllers.Pasaransaveconfcjitu)
	app.Post("/api/savepasaranconf5050umum", controllers.Pasaransaveconf5050umum)
	app.Post("/api/savepasaranconf5050special", controllers.Pasaransaveconf5050special)
	app.Post("/api/savepasaranconf5050kombinasi", controllers.Pasaransaveconf5050kombinasi)
	app.Post("/api/savepasaranconfmacaukombinasi", controllers.Pasaransaveconfmacaukombinasi)
	app.Post("/api/savepasaranconfdasar", controllers.Pasaransaveconfdasar)
	app.Post("/api/savepasaranconfshio", controllers.Pasaransaveconfshio)

	app.Post("/api/allperiode", controllers.Periode)
	app.Post("/api/editperiode", controllers.Periodedetail)
	app.Post("/api/saveperiode", controllers.Periodesave)
	app.Post("/api/savepasarannew", controllers.Periodesavenew)
	app.Post("/api/saveperioderevisi", controllers.Periodesaverevisi)
	app.Post("/api/cancelbet", controllers.Periodecancelbet)
	app.Post("/api/periodelistmember", controllers.Periodelistmember)
	app.Post("/api/periodelistmemberbynomor", controllers.Periodelistmemberbynomor)
	app.Post("/api/periodelistbet", controllers.Periodelistbet)
	app.Post("/api/periodelistbetstatus", controllers.Periodelistbetstatus)
	app.Post("/api/periodelistbetusername", controllers.Periodelistbetbyusername)
	app.Post("/api/periodelistbettable", controllers.Periodelistbettable)
	app.Post("/api/periodebettable", controllers.Periodebettable)
	app.Post("/api/listpasaran", controllers.Periodelistpasaran)
	app.Post("/api/listprediksi", controllers.Periodeprediksi)

	app.Post("/api/alladmin", controllers.Admin)
	app.Post("/api/editadmin", controllers.Admindetail)
	app.Post("/api/saveadmin", controllers.Adminsave)
	app.Post("/api/saveadminiplist", controllers.Adminsaveiplist)
	app.Post("/api/deleteadminiplist", controllers.Deleteiplist)

	app.Post("/api/alladminrule", controllers.Adminrule)
	app.Post("/api/editadminrule", controllers.Adminruledetail)
	app.Post("/api/saveadminrule", controllers.Adminrulesave)
	app.Post("/api/saveadminruleconf", controllers.Adminruleconfsave)

	app.Post("/api/log", controllers.Log)
	return app
}
