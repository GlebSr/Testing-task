package storage

type DishStorage struct {
	DishRepository DishRepository
}

func (d *DishStorage) Dish() DishRepository {
	return d.DishRepository
}

type MealStorage struct {
	MealRepository MealRepository
}

func (m *MealStorage) Meal() MealRepository {
	return m.MealRepository
}

type Storage struct {
	DishStorage DishStorage
	MealStorage MealStorage
}

func (receiver *Storage) Dish() DishRepository {
	return receiver.DishStorage.Dish()
}
func (receiver *Storage) Meal() MealRepository {
	return receiver.MealStorage.Meal()
}
