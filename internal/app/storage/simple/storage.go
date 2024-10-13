package simple

import (
	"github.com/GlebSr/app/internal/app/model"
	"github.com/GlebSr/app/internal/app/storage"
)

func CreateStorage() storage.Storage {
	return storage.Storage{
		MealStorage: storage.MealStorage{MealRepository: &MealRepository{meals: make(map[int]*model.Meal)}},
		DishStorage: storage.DishStorage{DishRepository: &DishRepository{dishes: make(map[int]*model.Dish)}},
	}
}
