package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Todo struct {
	ID      int  `csv:"id"`
	Task    interface{}  `csv:"task"`
	Created time.Time  `csv:"created"`
}

func AddTask(task interface{}) error{
	file, err := os.Create("task.csv")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing %s/n", err)
		return err
	}
	defer file.Close()

	tasks := []*Todo{
		{1, task, time.Now()},
	}
	err = gocsv.MarshalFile(tasks, file)
	if err != nil {
		return err
	}
return nil
}
