package storage

import (
	"github.com/GlebSr/app/internal/app/model"
	"time"
)

type DishRepository interface {
	Create(*model.Dish) (int, error)
	Delete(int) error
	Edit(*model.Dish) error
	Get(int) (*model.Dish, error)
	GetAll() ([]*model.Dish, error)
	GetIdByName(string) (int, error)
}

type MealRepository interface {
	Create(*model.Meal) (int, error)
	Delete(int) error
	Edit(*model.Meal) error
	Get(int) (*model.Meal, error)
	GetIdsBetween(time.Time, time.Time) ([]*model.Meal, error)
}
