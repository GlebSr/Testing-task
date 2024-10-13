let chart; // Хранение экземпляра графика

document.getElementById('fetch-meals').addEventListener('click', () => {
    const fromDate = new Date(document.getElementById('from-date').value).getTime() / 1000;
    const toDate = new Date(document.getElementById('to-date').value).getTime() / 1000;

    fetch(`http://localhost:8080/api/v1/meal/between/?from=${fromDate}&to=${toDate}`)
        .then(response => response.json())
        .then(meals => {
            const mealData = {};

            // Сбор всех id блюд для последующего запроса
            const dishIds = new Set();

            meals.forEach(meal => {
                Object.keys(meal.dishes).forEach(dishId => {
                    dishIds.add(dishId); // Добавляем id блюда
                });
            });

            // Получение данных о каждом блюде
            Promise.all([...dishIds].map(id => fetch(`http://localhost:8080/api/v1/dish?id=${id}`)))
                .then(responses => Promise.all(responses.map(res => res.json())))
                .then(dishes => {
                    const dishMap = {};
                    dishes.forEach(dish => {
                        dishMap[dish.id] = dish; // Создаём словарь для быстрого доступа к блюдам
                    });

                    meals.forEach(meal => {
                        const mealTime = new Date(meal.mealTime).toDateString(); // Форматируем дату
                        const calories = Object.keys(meal.dishes).reduce((sum, dishId) => {
                            const weight = meal.dishes[dishId]; // Получаем массу блюда
                            const dish = dishMap[dishId]; // Получаем информацию о блюде из словаря

                            if (dish) {
                                return sum + (dish.calories * weight / 100); // Считаем калории
                            }
                            return sum;
                        }, 0); // Считаем сумму калорий

                        if (mealData[mealTime]) {
                            mealData[mealTime] += calories; // Суммируем калории по дате
                        } else {
                            mealData[mealTime] = calories; // Создаём новую запись
                        }
                    });

                    renderChart(mealData); // Рендерим график
                });
        });
});

// Функция для отрисовки графика
function renderChart(mealData) {
    const ctx = document.getElementById('mealChart').getContext('2d');

    const labels = Object.keys(mealData);
    const data = Object.values(mealData);

    // Если график уже существует, то удаляем его перед созданием нового
    if (chart) {
        chart.destroy(); // Удаляем старый график
    }

    // Создаём новый график
    chart = new Chart(ctx, {
        type: 'line', // Тип графика
        data: {
            labels: labels,
            datasets: [{
                label: 'Сумма калорий по дням',
                data: data,
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderWidth: 2
            }]
        },
        options: {
            responsive: true,
            scales: {
                y: {
                    beginAtZero: true,
                    title: {
                        display: true,
                        text: 'Калории'
                    }
                },
                x: {
                    title: {
                        display: true,
                        text: 'Дата'
                    }
                }
            }
        }
    });
}
