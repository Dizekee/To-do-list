package main

import (
	"fmt"
	"time"

	"github.com/Dizekee/To-do-list/manager"
)

func main() {
	// Создаем менеджер задач
	tm := &manager.TaskManager{}

	fmt.Println("=== ТЕСТИРОВАНИЕ СИСТЕМЫ УПРАВЛЕНИЯ ЗАДАЧАМИ ===\n")

	// Добавляем рабочие задачи
	fmt.Println("1. Добавление рабочих задач:")
	err := tm.AddTask("work", "Завершить проект", "Закончить разработку API", 3,
		time.Now().AddDate(0, 0, 5), map[string]string{
			"project":  "ToDo App",
			"assignee": "Иван Иванов",
		})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = tm.AddTask("work", "Code review", "Проверить PR от команды", 4,
		time.Now().AddDate(0, 0, 2), map[string]string{
			"project":  "Backend",
			"assignee": "Петр Петров",
		})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Добавляем личные задачи
	fmt.Println("\n2. Добавление личных задач:")
	err = tm.AddTask("personal", "Поход в магазин", "Купить продукты на неделю", 2,
		time.Now().AddDate(0, 0, 1), map[string]string{
			"category": "Покупки",
			"location": "Супермаркет",
		})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = tm.AddTask("personal", "Спорт", "Тренировка в зале", 1,
		time.Now().AddDate(0, 0, -1), map[string]string{ // Просроченная задача
			"category": "Здоровье",
			"location": "Фитнес-клуб",
		})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Пытаемся добавить задачу с неверным приоритетом
	fmt.Println("\n3. Попытка добавления задачи с неверным приоритетом:")
	err = tm.AddTask("work", "Неверная задача", "Описание", 6,
		time.Now().AddDate(0, 0, 5), map[string]string{
			"project":  "Test",
			"assignee": "Test",
		})
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	}

	// Выводим все задачи
	fmt.Println("\n4. Все задачи в системе:")
	for i, task := range tm.GetTasksByPriority(1) { // Получаем все задачи (мин. приоритет 1)
		fmt.Printf("%d. [ID:%d] %s - Приоритет: %d, Выполнена: %t, До срока: %d дней\n",
			i+1, task.GetID(), task.GetTitle(), task.GetPriority(),
			task.IsComplited(), task.DaysUntilDue())
	}

	// Отмечаем задачу как выполненную
	fmt.Println("\n5. Отмечаем задачу ID:1 как выполненную:")
	if tm.CompleteTask(1) {
		fmt.Println("Задача ID:1 отмечена как выполненная")
	}

	// Получаем задачу по ID
	fmt.Println("\n6. Поиск задачи по ID:2")
	task, err := tm.GetTask(2)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Найдена задача: %s (Приоритет: %d)\n", task.GetTitle(), task.GetPriority())
	}

	// Просроченные задачи
	fmt.Println("\n7. Просроченные задачи:")
	overdue := tm.GetOverdueTasks()
	if len(overdue) == 0 {
		fmt.Println("Просроченных задач нет")
	} else {
		for i, task := range overdue {
			fmt.Printf("%d. %s (просрочено на %d дней)\n",
				i+1, task.GetTitle(), -task.DaysUntilDue())
		}
	}

	// Задачи по типам
	fmt.Println("\n9. Рабочие задачи:")
	workTasks := tm.GetTasksByType("work")
	for i, task := range workTasks {
		fmt.Printf("%d. %s\n", i+1, task.GetTitle())
	}

	fmt.Println("\n10. Личные задачи:")
	personalTasks := tm.GetTasksByType("personal")
	for i, task := range personalTasks {
		fmt.Printf("%d. %s\n", i+1, task.GetTitle())
	}

	// Удаление задачи
	fmt.Println("\n11. Удаление задачи ID:3:")
	if tm.RemoveTask(3) {
		fmt.Println("Задача ID:3 удалена")
	} else {
		fmt.Println("Задача ID:3 не найдена")
	}

	// Итоговый список задач
	fmt.Println("\n12. Итоговый список задач:")
	for i, task := range tm.GetTasksByPriority(1) {
		fmt.Printf("%d. [ID:%d] %s - Приоритет: %d, Выполнена: %t\n",
			i+1, task.GetID(), task.GetTitle(), task.GetPriority(), task.IsComplited())
	}

	fmt.Println("\n=== ТЕСТИРОВАНИЕ ЗАВЕРШЕНО ===")
}
