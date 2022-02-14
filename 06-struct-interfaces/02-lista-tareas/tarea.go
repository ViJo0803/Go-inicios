package main

import "fmt"

//Lista de tareas

type taskList struct {
	tasks []*Task
}

//metodo agregar tarea
func (tl *taskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

//metodo eliminar tarea
func (tl *taskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

//Tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

//metodo
func (t *Task) toPrint() {
	fmt.Printf("Nombre : %s\n Descripcion: %s\n Completado; %t \n", t.name, t.desc, t.completed)
}

func (t *Task) markcompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de go",
		desc:      "Completar curso de go este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso HTML",
		desc:      "Completar curso de html esta semana",
		completed: true,
	}

	lista := taskList{}

	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println(lista)
	lista.removeTask(1)
	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}

	//t1.toPrint()
	//t2.toPrint()
}
