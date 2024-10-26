package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/gocarina/gocsv"
)

type Todo struct {
	ID      int    `csv:"id"`
	Task    string `csv:"task"`
	Done    bool   `csv:"done"`
	Created string `csv:"created"`
}

func AddTask(task string) (id int, err error) {
	file, err := os.OpenFile("task.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("err: %s", err)
		return 0, err
	}
	defer file.Close()

	existingFile, err := os.Open("task.csv")
	if err != nil {
		fmt.Printf("err: %s", err)
		return 0, err
	}
	defer existingFile.Close()

	fileInfo, err := existingFile.Stat()
	if err != nil {
		fmt.Printf("err: %s", err)
		return 0, err
	}

	var tasks []*Todo

	if fileInfo.Size() > 0 {

		if err := gocsv.UnmarshalFile(existingFile, &tasks); err != nil {
			fmt.Printf("err: %s", err)
			return 0, err
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
		return 0, err
	}
	return newID, nil
}

func ListTask(showCompletedTask bool) error {
	file, err := os.Open("task.csv")
	if err != nil {
		fmt.Printf("err: %s", err)
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err: %s", err)
		return err
	}

	if fileInfo.Size() <= 0 {
		fmt.Println("No tasks...")
		return err
	}

	var tasks []Todo

	if err := gocsv.UnmarshalFile(file, &tasks); err != nil {
		fmt.Printf("err: %s", err)
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	if showCompletedTask {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
		for _, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", task.ID, task.Task, task.Created, task.Done)
		}
		w.Flush()
		return nil
	}

	fmt.Fprintln(w, "ID\tTask\tCreated")
	for _, task := range tasks {
		if !task.Done {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Task, task.Created)
		}
	}
	w.Flush()
	return nil
}
