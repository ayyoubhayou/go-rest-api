package routes

import (
	"github.com/go-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/movie", controllers.CreateMovie)
	app.Get("/movie/:title", controllers.GetAMovie)
	app.Put("/movie/:movieId", controllers.EditAMovie)
	app.Delete("/movie/:movieId", controllers.DeleteAMovie)
	app.Get("/movies", controllers.GetAllMovies)
}
