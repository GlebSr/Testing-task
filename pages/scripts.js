document.addEventListener('DOMContentLoaded', () => {
    loadDishes(); // Загружаем список блюд при загрузке страницы

    // Добавление нового блюда
    const dishForm = document.getElementById('dish-form');
    dishForm.addEventListener('submit', event => {
        event.preventDefault();

        const dishData = {
            name: dishForm.name.value,
            calories: parseInt(dishForm.calories.value, 10),
            proteins: parseInt(dishForm.proteins.value, 10),
            fats: parseInt(dishForm.fats.value, 10),
            carbohydrates: parseInt(dishForm.carbohydrates.value, 10)
        };

        fetch('http://localhost:8080/api/v1/dish', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dishData)
        }).then(response => {
            if (response.ok) {
                alert('Блюдо добавлено!');
                dishForm.reset(); // Очищаем форму после отправки
                loadDishes();     // Обновляем список блюд
            }
        });
    });

    // Добавляем обработчик отправки приёма пищи
    document.getElementById('submit-meal').addEventListener('click', () => {
        submitMeal();
    });
});

// Загружаем список блюд
function loadDishes() {
    fetch('http://localhost:8080/api/v1/dish/all')
        .then(response => response.json())
        .then(data => {
            const dishList = document.getElementById('dish-list');
            dishList.innerHTML = ''; // Очищаем список перед обновлением
            data.forEach(dish => {
                const dishItem = document.createElement('div');
                dishItem.classList.add('dish-item');
                dishItem.innerHTML = `
                    ${dish.name} - КБЖУ: ${dish.calories} / ${dish.proteins} / ${dish.fats} / ${dish.carbohydrates}
                    <button onclick="addToMeal(${dish.id}, '${dish.name}')">Добавить</button>
                    <button onclick="editDish(${dish.id})">Редактировать</button>
                    <button onclick="deleteDish(${dish.id})">Удалить</button>
                `;
                dishList.appendChild(dishItem);
            });
        });
}

// Добавление блюда в приём пищи
let meal = {}; // Храним блюда с их массой (граммы)

function addToMeal(dishId, dishName) {
    const grams = prompt(`Введите количество граммов для блюда "${dishName}":`);
    if (grams && !isNaN(grams)) {
        meal[dishId] = parseInt(grams, 10);
        updateMealSummary();
    }
}

// Обновляем сводку по приему пищи
function updateMealSummary() {
    const mealSummary = document.getElementById('meal-summary-items');
    mealSummary.innerHTML = '';

    let totalCalories = 0;
    let totalProteins = 0;
    let totalFats = 0;
    let totalCarbohydrates = 0;

    // Для каждого блюда в приёме пищи
    const dishIds = Object.keys(meal);
    if (dishIds.length === 0) {
        document.getElementById('total-kbju').innerText = '0'; // Если нет блюд, показываем 0
        return;
    }

    dishIds.forEach(dishId => {
        const grams = meal[dishId];
        fetch(`http://localhost:8080/api/v1/dish?id=${dishId}`)
            .then(response => response.json())
            .then(dish => {
                const dishSummary = document.createElement('div');
                dishSummary.classList.add('meal-summary-item');
                dishSummary.innerHTML = `
                    ${dish.name}: ${grams} грамм
                    <button onclick="removeFromMeal(${dishId})">Удалить</button>
                `;
                mealSummary.appendChild(dishSummary);

                // Рассчитываем общее КБЖУ
                totalCalories += (dish.calories * grams) / 100;
                totalProteins += (dish.proteins * grams) / 100;
                totalFats += (dish.fats * grams) / 100;
                totalCarbohydrates += (dish.carbohydrates * grams) / 100;

                // Обновляем общее КБЖУ
                document.getElementById('total-kbju').innerText = `
                    Калории: ${totalCalories.toFixed(2)}, 
                    Белки: ${totalProteins.toFixed(2)}, 
                    Жиры: ${totalFats.toFixed(2)}, 
                    Углеводы: ${totalCarbohydrates.toFixed(2)}
                `;
            });
    });
}

// Удаление блюда из приёма пищи
function removeFromMeal(dishId) {
    delete meal[dishId];
    updateMealSummary();
}

// Отправка приёма пищи
function submitMeal() {
    const mealData = {
        dishes: meal
    };

    fetch('http://localhost:8080/api/v1/meal', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(mealData)
    }).then(response => {
        if (response.ok) {
            alert('Приём пищи успешно отправлен!');
            document.getElementById('submit-meal').classList.add('success');
            setTimeout(() => {
                document.getElementById('submit-meal').classList.remove('success');
                meal = {}; // Очищаем список приёмов пищи
                updateMealSummary(); // Очищаем сводку
            }, 2000);
        }
    });
}

// Редактирование блюда
function editDish(dishId) {
    const newName = prompt('Введите новое название блюда:');
    const newCalories = prompt('Введите новое количество калорий:');
    const newProteins = prompt('Введите новое количество белков:');
    const newFats = prompt('Введите новое количество жиров:');
    const newCarbohydrates = prompt('Введите новое количество углеводов:');

    if (newName && newCalories && newProteins && newFats && newCarbohydrates) {
        const updatedDish = {
            name: newName,
            calories: parseInt(newCalories, 10),
            proteins: parseInt(newProteins, 10),
            fats: parseInt(newFats, 10),
            carbohydrates: parseInt(newCarbohydrates, 10)
        };

        fetch(`http://localhost:8080/api/v1/dish?id=${dishId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(updatedDish)
        }).then(response => {
            if (response.ok) {
                alert('Блюдо успешно обновлено!');
                loadDishes(); // Обновляем список блюд
            }
        });
    }
}

// Удаление блюда
function deleteDish(dishId) {
    if (confirm('Вы уверены, что хотите удалить это блюдо?')) {
        fetch(`http://localhost:8080/api/v1/dish?id=${dishId}`, {
            method: 'DELETE'
        }).then(response => {
            if (response.ok) {
                alert('Блюдо удалено!');
                loadDishes(); // Обновляем список блюд
            }
        });
    }
}
