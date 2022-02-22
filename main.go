package main

import (
	"fmt"
	"log"
	"module30/pkg/storage"
	"os"
)

func main() {
	var err error
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" +
			pwd + "@localhost:5432/tasks"
	// присвоение переменной типа интерфейс конкретной реализации БД
	db, err := storage.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	//task := storage.Task{
	//	Title:   "Learn Go language",
	//	Content: "Need to learn Golang basics",
	//}
	//task2 := storage.Task{
	//	Title:   "Learn DB integration in Go language",
	//	Content: "Need to learn MySQL and PostgreSQL Golang integration basics",
	//}
	//_, err = db.NewTask(task)
	//_, err = db.NewTask(task2)

	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All tasks")
	fmt.Println(tasks)
	tasks, err = db.Tasks(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tasks by authorID=1")
	fmt.Println(tasks)
	tasks, err = db.TasksByLabel(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tasks by labelId=0")
	fmt.Println(tasks)
	taskUpdated := storage.Task{
		Title:   "Learn SQL DB integration in Go language",
		Content: "Need to learn MySQL and PostgreSQL Golang integration advanced level",
	}
	_, err = db.UpdateTask(taskUpdated, 1)
	tasks, err = db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tasks after edit")
	fmt.Println(tasks)

	rowsAffected, err := db.RemoveTask(1)
	if rowsAffected == 0 {
		log.Fatal(err)
	}
	tasks, err = db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tasks after remove")
	fmt.Println(tasks)

	rowsAffected, err = db.RemoveTask(1)
	if rowsAffected == 0 {
		log.Fatal(err)
	}
	tasks, err = db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tasks after remove")
	fmt.Println(tasks)
}
