package main

import (
	"fmt"

	"github.com/google/uuid"

	"StudyProject/TaskManager/model"
)

func main() {
	// Накидали новых таск
	newTask("Kazhetsya, moi status izmenitsya")
	newTask("Kazhetsya, menya udalyat")
	newTask("Chto proishodit")

	// Вывели все таски
	fmt.Println("Список таск:")
	ALLtaskInfomation()
	fmt.Println("")

	// Вывели список тасок после изменения статуса одной из них
	fmt.Println("Список таск после изменений статуса:")
	changeStatus(0)
	ALLtaskInfomation()
	fmt.Println("")

	// Вывели список тасок после удаления одной из них
	fmt.Println("Список таск после удаления:")
	deleteTask(TasksMassive[1].ID)
	ALLtaskInfomation()
}

// TasksMassive создали хранилище (в виде массива) под таски
var TasksMassive []model.Task

// Создаём новую таску
func newTask(desc string) int {
	// Создаём новый объект task типа "Task"
	task := model.Task{
		ID:        uuid.New(),
		Desc:      desc,
		Completed: false,
	}
	// Добавляем созданную таску в массив таск TasksMassive
	TasksMassive = append(TasksMassive, task)

	// Возвращаем ID созданной таски, в виде её порядкового номера в массиве TasksMassive
	return len(TasksMassive) - 1
}

// ALLtaskInfomation функция вывода информации обо всех тасках
func ALLtaskInfomation() {
	for i := range TasksMassive {
		taskInfomation(i)
	}
}

// Функция вывод информации о выбранной таски
func taskInfomation(id int) {
	fmt.Printf("ID: %s, Desc: %s, Completed: %t\n", TasksMassive[id].ID, TasksMassive[id].Desc, TasksMassive[id].Completed)
}

// Функция по замене статуса выбранной таски
func changeStatus(id int) {
	TasksMassive[id].Completed = true
}

// Удаление из массива TasksMassive таски, чей uuid указан
func deleteTask(id uuid.UUID) {
	for i, task := range TasksMassive {
		if task.ID == id {
			TasksMassive = append(TasksMassive[:i], TasksMassive[i+1:]...)
			break
		}
	}
}
