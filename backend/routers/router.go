package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/isbpanel_movie2/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Post("api/initmovie", controller.Init)
	app.Post("api/listmovie", controller.Listmovie)
	app.Post("api/listseason", controller.Listseason)
	app.Post("api/listepisode", controller.Listepisode)
	return app
}
