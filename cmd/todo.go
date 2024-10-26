package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Todo struct {
	ID      int    `csv:"id"`
	Task    string `csv:"task"`
	Done    bool   `csv:"done"`
	Created string `csv:"created"`
}

func AddTask(task string) {
	file, err := os.OpenFile("task.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	defer file.Close()

	existingFile, err := os.Open("task.csv")
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	defer existingFile.Close()

	fileInfo, err := existingFile.Stat()
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var tasks []*Todo

	if fileInfo.Size() > 0 {

		if err := gocsv.UnmarshalFile(existingFile, &tasks); err != nil {
			fmt.Printf("err: %s", err)
			return
		}
	}
	var newID int
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newID = lastTask.ID + 1
	} else {
		newID = 1
	}

	newTask := &Todo{
		ID:      newID,
		Task:    task,
		Done:    false,
		Created: time.Now().Format("02 Jan 2006, 3:04 PM"),
	}

	if fileInfo.Size() == 0 {
		err = gocsv.MarshalFile([]*Todo{newTask}, file)
	} else {
		err = gocsv.MarshalWithoutHeaders([]*Todo{newTask}, file)
	}

	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
}

func ListTask() {
	file, err := os.Open("task.csv")
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	if fileInfo.Size() <= 0 {
		fmt.Println("No tasks...")
		return
	}

	var tasks []*Todo

	if err := gocsv.UnmarshalFile(file, &tasks); err != nil {
		fmt.Printf("err: %s", err)
		return
	}

}
