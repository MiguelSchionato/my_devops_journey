package templates

import (
	"time"
	//"fmt"
)

type Task struct {
	Name  string
	ID    int
	Descr string
	Date  time.Time
	Due   time.Time
	State int
	/*
		isComplete bool
		hasStarted bool
	*/

}

type Lists struct {
	Name  string
	Descr string
	Tasks []Task
}

type Config struct {
	CurrentListPath string `json:"CurrentListPath"`
	CurrentList     string `json:"CurrentList"`
}
