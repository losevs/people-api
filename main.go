package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/losevs/people-api/handlers"
	"github.com/losevs/people-api/logger"
)

func main() {
	defer logger.Logfile.Close()
	app := fiber.New()

	setupRoutes(app)

	logger.Logg.Info("Starting the app on port localhost:80")
	log.Fatalln(app.Listen(":80"))
}

func setupRoutes(app *fiber.App) {
	show := app.Group("/show") // /show
	show.Get("/", handlers.ShowAll)
	show.Get("/age/asc", handlers.AgeType)
	show.Get("/:id", handlers.ShowByID)
	//PAGINATION
	pagShow := show.Group("/pag") // /show/pag
	pagShow.Get("/:page", handlers.Pag)
	pagShow.Get("/men/:page", handlers.MenPag)
	pagShow.Get("/wmen/:page", handlers.WMenPag)
	//FILTERS:
	filtShow := show.Group("/filt") // /show/filt
	filtShow.Get("/sex/:sex", handlers.FiltSex)
	filtShow.Get("/country/:country", handlers.FiltCountry)
	filtShow.Get("/age/:age", handlers.FiltAge)

	app.Post("/new", handlers.AddNew)

	app.Patch("/change/:id", handlers.PatchByID)

	app.Delete("del/:id", handlers.DeleteByID)
}

/*
PAGINATION:
- всех по 4 на странице

-- Мужчин 5 на странице
-- Женщин 5 на странице

FILTERS:
- по полу
- по стране

- по возрасту = по возрастанию?

*/
