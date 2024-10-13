package server

import (
	"github.com/GlebSr/app/internal/app/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type server struct {
	app     *fiber.App
	storage *storage.Storage
}

func NewServer(storage *storage.Storage) *server {
	serv := &server{
		app:     fiber.New(),
		storage: storage,
	}
	serv.configureRouter()
	return serv
}

func (serv *server) configureRouter() {
	serv.app.Use(cors.New())
	serv.app.Static("/", "./pages")

	apiRouterV1 := serv.app.Group("/api/v1")
	dishRouter := apiRouterV1.Group("/dish")
	dishRouter.Post("", serv.handleNewDish())
	dishRouter.Put("", serv.handleEditDish())
	dishRouter.Get("", serv.handleGetDish())
	dishRouter.Get("/name/", serv.handleGetByNameDish())
	dishRouter.Get("/all", serv.handleGetAllDishes())

	dishRouter.Delete("/", serv.handleDeleteDish())
	mealRouter := apiRouterV1.Group("/meal")
	mealRouter.Post("", serv.handleNewMeal())
	mealRouter.Get("", serv.handleGetMeal())
	mealRouter.Get("/between/", serv.handleGetMeals())
	dishRouter.Delete("", serv.handleDeleteMeal())
}

func (serv *server) Listen(adr string) error {
	log.Println("Start listen")
	return serv.app.Listen(adr)
}
