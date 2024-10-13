CREATE TABLE dishes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    calories DECIMAL(10, 2),
    proteins DECIMAL(10, 2),
    fats DECIMAL(10, 2),
    carbohydrates DECIMAL(10, 2)
);

CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    meal_time TIMESTAMP NOT NULL
);

CREATE TABLE meal_dishes (
    meal_id INT REFERENCES meals(id) ON DELETE CASCADE,
    dish_id INT REFERENCES dishes(id) ON DELETE CASCADE,
    weight_in_grams DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (meal_id, dish_id)
);
