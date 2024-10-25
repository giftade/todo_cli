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
	file, err := os.OpenFile("task.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing %s/n", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	tasks := []*Todo{
		{1, task, time.Now()},
	}

	if fileInfo.Size() == 0 {
		err = gocsv.MarshalFile(tasks, file)
		} else {
		err = gocsv.MarshalWithoutHeaders(tasks, file)
	}
	if err != nil {
		return err
	}
return nil
}
