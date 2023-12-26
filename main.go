package main

import (
	"fmt"

	"github.com/google/uuid"

	"StudyProject/TaskManager/model"
)

// TasksMassive Создали типа хранилище (в виде массива) под таски
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

func main() {
	newTask("Chto proishodit")
	newTask("Chto za uzhas tut tvoritsya")
	newTask("Prosto za chto")

	//id := newTask("ПП работает фигово")
	//fmt.Println(TasksMassive[id])
	//fmt.Println(TasksMassive)
	ALLtaskInfomation()
	fmt.Println("просто текст для прерывания потока данных")
	fmt.Println("просто текст для прерывания потока данных")
	fmt.Println("просто текст для прерывания потока данных")
	//changeStatus(id)
	////fmt.Println(TasksMassive)
	//
	//taskInfomation(id)
	//
	deleteTask(TasksMassive[1].ID)
	ALLtaskInfomation()
}

// Функция на замемену статуса выбранной таски
//func changeStatus(id int) {
//	TasksMassive[id].Completed = true
//}

// Функция вывод информации о выбранной таски
func taskInfomation(id int) {
	fmt.Printf("ID: %s, Desc: %s, Completed: %t\n", TasksMassive[id].ID, TasksMassive[id].Desc, TasksMassive[id].Completed)

}

// ALLtaskInfomation Функция вывод информации о всех тасках
func ALLtaskInfomation() {
	for i := range TasksMassive {
		taskInfomation(i)
	}
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

//func main() {
//
//	//y := culculation.GetMaxNumber(123456789)
//	//fmt.Println(y)
//
//	//x := culculation.GetReverseOrder(123)
//	//fmt.Println(x)
//
//	//z := culculation.RemoveTheNumber(1234567, 3)
//	//fmt.Println(z)
//	//fmt.Println(len(z))
//
//	strangeService.Start()
//}
