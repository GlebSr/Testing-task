package simple

import (
	"fmt"
	"github.com/GlebSr/app/internal/app/model"
	"time"
)

type MealRepository struct {
	meals map[int]*model.Meal
}

func (m *MealRepository) Create(meal *model.Meal) (int, error) {
	meal.SetId(len(m.meals))
	m.meals[meal.GetId()] = meal
	return meal.GetId(), nil
}

func (m *MealRepository) Delete(id int) error {
	if _, ok := m.meals[id]; !ok {
		return fmt.Errorf("Meal with id=%d not exist.", id)
	}
	delete(m.meals, id)
	return nil
}

func (m *MealRepository) Edit(meal *model.Meal) error {
	if _, ok := m.meals[meal.GetId()]; !ok {
		return fmt.Errorf("Meal with id=%d not exist.", meal.GetId())
	}
	m.meals[meal.GetId()] = meal
	return nil
}

func (m *MealRepository) Get(id int) (*model.Meal, error) {
	return m.meals[id], nil
}

func (m *MealRepository) GetIdsBetween(from time.Time, to time.Time) ([]*model.Meal, error) {
	meals := make([]*model.Meal, 0)
	for _, value := range m.meals {
		if value.GetTime().Before(to) && from.Before(value.GetTime()) {
			meals = append(meals, value)
		}
	}
	return meals, nil
}
