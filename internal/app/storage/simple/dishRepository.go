package simple

import (
	"fmt"
	"github.com/GlebSr/app/internal/app/model"
)

type DishRepository struct {
	dishes map[int]*model.Dish
}

func (d *DishRepository) Create(dish *model.Dish) (int, error) {
	dish.SetId(len(d.dishes))
	d.dishes[dish.GetId()] = dish
	return dish.GetId(), nil
}

func (d *DishRepository) Delete(id int) error {
	if _, ok := d.dishes[id]; !ok {
		return fmt.Errorf("Dish with id=%d not exist.", id)
	}
	delete(d.dishes, id)
	return nil
}

func (d *DishRepository) Edit(dish *model.Dish) error {
	if _, ok := d.dishes[dish.GetId()]; !ok {
		return fmt.Errorf("Dish with id=%d not exist.", dish.GetId())
	}
	d.dishes[dish.GetId()] = dish
	return nil
}

func (d *DishRepository) Get(id int) (*model.Dish, error) {
	return d.dishes[id], nil
}

func (d *DishRepository) GetAll() ([]*model.Dish, error) {
	dishes := make([]*model.Dish, 0)
	for _, value := range d.dishes {
		dishes = append(dishes, value)
	}
	return dishes, nil
}

func (d *DishRepository) GetIdByName(s string) (int, error) {
	//TODO implement me
	panic("implement me")
}

//type DishRepository interface {
//	Create(*model.Dish) (int, error)
//	Delete(int) error
//	Edit(*model.Dish) error
//	Get(int) (*model.Dish, error)
//	GetAll() ([]*model.Dish, error)
//	GetIdByName(string) (int, error)
//}
