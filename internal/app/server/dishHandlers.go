package server

import (
	"encoding/json"
	"github.com/GlebSr/app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (serv *server) handleNewDish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("NewDish")
		dish := new(model.Dish)
		log.Println(string(ctx.Body()))
		if err := json.Unmarshal(ctx.Body(), dish); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		log.Println(*dish)
		if !dish.Available() {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		_, err := serv.storage.Dish().Create(dish)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (serv *server) handleEditDish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("EditDish")
		dish := new(model.Dish)
		log.Println(string(ctx.Body()))
		if err := json.Unmarshal(ctx.Body(), dish); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		log.Println(*dish)
		if !dish.Available() {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		err := serv.storage.Dish().Edit(dish)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (serv *server) handleGetDish() fiber.Handler {
	log.Println("GetDish")
	return func(ctx *fiber.Ctx) error {
		dish, err := serv.storage.Dish().Get(ctx.QueryInt("id", -1))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		dishJson, err := json.Marshal(dish)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Send(dishJson)
	}
}

func (serv *server) handleGetByNameDish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("GetByNameDish")
		dishId, err := serv.storage.Dish().GetIdByName(ctx.Query("name"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		dishJson, err := json.Marshal(struct {
			id int `json:"id"`
		}{
			id: dishId,
		})
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Send(dishJson)
	}
}

func (serv *server) handleGetAllDishes() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("GetAllDishes")
		dishes, err := serv.storage.Dish().GetAll()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		dishesJson, err := json.Marshal(dishes)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Send(dishesJson)
	}
}

func (serv *server) handleDeleteDish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("DeleteDish")
		err := serv.storage.Dish().Delete(ctx.QueryInt("id", -1))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}
