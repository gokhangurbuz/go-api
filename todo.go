package main

import "time"

type Todo struct {
	Id 	  int	    `json:"Id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
