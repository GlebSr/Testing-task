package model

// cpfc multiplied by 100
type Dish struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Calories      int    `json:"calories"`
	Proteins      int    `json:"proteins"`
	Fats          int    `json:"fats"`
	Carbohydrates int    `json:"carbohydrates"`
}

func CreateDish(name string, calories int, proteins int, fats int, carbohydrates int) *Dish {
	if calories < 0 || proteins < 0 || fats < 0 || carbohydrates < 0 {
		return nil
	}
	return &Dish{
		Id:            -1,
		Name:          name,
		Calories:      calories,
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
	}
}

func (d *Dish) GetId() int {
	return d.Id
}

func (d *Dish) SetId(id int) {
	d.Id = id
}

func (d *Dish) Available() bool {
	return d.Name != "" && d.Calories >= 0 && d.Proteins >= 0 && d.Fats >= 0 && d.Carbohydrates >= 0
}
