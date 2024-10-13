package server

import (
	"encoding/json"
	"github.com/GlebSr/app/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

func (serv *server) handleNewMeal() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		meal := new(model.Meal)
		log.Println(string(ctx.Body()))
		if err := ctx.BodyParser(meal); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		meal.SetTime(time.Now())
		log.Println("aval")
		if !meal.Available() {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		_, err := serv.storage.Meal().Create(meal)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (serv *server) handleGetMeal() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		meal, err := serv.storage.Meal().Get(ctx.QueryInt("id", -1))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		dishJson, err := json.Marshal(meal)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Send(dishJson)
	}
}

func (serv *server) handleGetMeals() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		from, err := strconv.ParseInt(ctx.Query("from", "0"), 10, 64)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		to, err := strconv.ParseInt(ctx.Query("to", "0"), 10, 64)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		meals, err := serv.storage.Meal().GetIdsBetween(time.Unix(from, 0), time.Unix(to, 0))
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		mealsJson, err := json.Marshal(meals)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Send(mealsJson)
	}
}

func (serv *server) handleDeleteMeal() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := serv.storage.Meal().Delete(ctx.QueryInt("id", -1))
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}
