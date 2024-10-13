package model

import "time"

type Meal struct {
	Id       int         `json:"id"`
	MealTime time.Time   `json:"meal_time"`
	Dishes   map[int]int `json:"dishes"`
}

func CreateMeal(dishes map[int]int) *Meal {
	if dishes == nil || len(dishes) == 0 {
		return nil
	}
	return &Meal{
		Id:       -1,
		MealTime: time.Now(),
		Dishes:   dishes,
	}
}

func (m *Meal) GetId() int {
	return m.Id
}

func (m *Meal) SetId(id int) {
	m.Id = id
}

func (m *Meal) SetTime(mealTime time.Time) {
	m.MealTime = mealTime
}

func (m *Meal) GetTime() time.Time {
	return m.MealTime
}

func (m *Meal) Available() bool {
	return !m.MealTime.IsZero() && m.MealTime.Before(time.Now()) && m.Dishes != nil && len(m.Dishes) > 0
}
