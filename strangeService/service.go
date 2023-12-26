package strangeService

//import "StudyProject/strangeService/methods"
//
//func Start() {
//
//	methods.Methods()
//
//}

import (
	"fmt"
	"math/rand"
)

// Структура для представления задачи
type Task struct {
	ID        int
	Desc      string
	Completed bool
}

// Функция для создания новой задачи
func NewTask(Desc string) *Task {
	// Генерируем ID для задачи
	taskID := int(rand.Intn(3))

	// Создаем новую задачу с указанным описанием и ID
	task := &Task{
		ID:        taskID,
		Desc:      Desc,
		Completed: false,
	}

	return task
}

// Функция для установки статуса задачи как выполненной
func SetTaskCompleted(task *Task) {
	// Устанавливаем статус задачи в выполненный
	task.Completed = true
}

// Функция для вывода информации о задаче
func PrintTask(task *Task) {
	if task == nil {
		// Если задача не указана, выводим сообщение об этом
		fmt.Println("No task")
		return
	}

	// Выводим информацию о задаче в формате: ID, Desc, Completed
	fmt.Printf("ID: %d, Desc: %s, Completed: %t\n", task.ID, task.Desc, task.Completed)
}

// Структура для хранения списка задач
type TaskList struct {
	Tasks []Task
}

// Функция для добавления новой задачи в список
func AddTask(list *TaskList, Desc string) {
	// Создаем новую задачу с указанным описанием
	task := NewTask(Desc)

	// Добавляем новую задачу в список
	list.Tasks = append(list.Tasks, *task)
}

// Функция для вывода всех задач в списке
func PrintTaskList(list *TaskList) {
	// Проходим по каждой задаче в списке и выводим ее информацию
	for _, task := range list.Tasks {
		PrintTask(&task)
	}
}

// Функция для удаления задачи из списка по ее ID
func DeleteTask(list *TaskList, id int) {
	// Проходим по каждой задаче в списке и ищем задачу с указанным ID
	for i, task := range list.Tasks {
		if task.ID == id {
			// Если задача найдена, удаляем ее из списка
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			break
		}
	}
}

// Точка входа в программу
func main() {
	// Создаем новый список задач
	list := TaskList{Tasks: []Task{}}

	// Добавляем несколько задач в список
	Desc := "First task"
	task := NewTask(Desc)
	list.Tasks = append(list.Tasks, *task)

	Desc = "Second task"
	task = NewTask(Desc)
	list.Tasks = append(list.Tasks, *task)

	Desc = "Third task"
	task = NewTask(Desc)
	list.Tasks = append(list.Tasks, *task)

	// Выводим все задачи в списке
	fmt.Println("All tasks:")
	PrintTaskList(&list)

	// Удаляем задачу с ID 2
	fmt.Println("\nDelete task with ID: 2")
	DeleteTask(&list, 2)

	// Выводим все задачи после удаления
	fmt.Println("All tasks after deletion:")
	PrintTaskList(&list)
}
