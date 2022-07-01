package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var add bool
var del int
var list bool
var msg string
var engine *Engine

func init() {
	flag.BoolVar(&add, "a", false, "add a new task to todolist ")
	flag.IntVar(&del, "d", 0, "delete task from todolist")
	flag.BoolVar(&list, "list", false, "list todo")
	flag.StringVar(&msg, "m", "", "task description")
	flag.Parse()
	var err error
	engine, err = Init()
	if err != nil {
		fmt.Printf("can not init todolist, %s\n", err)
		os.Exit(1)
	}
}
func main() {
	if add {
		if msg == "" {
			fmt.Printf("use -m provide descript for new task")
			fmt.Printf("try run -a -m \"task description\"")
			os.Exit(1)
		}
		err := engine.Add(msg)
		if err != nil {
			fmt.Printf("create a new task , %s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if del != 0 {
		err := engine.Delete(del)
		if err != nil {
			fmt.Printf("delete task, %s", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if list {
		list, err := engine.ListALL()
		if err != nil {
			fmt.Printf("read todo list failed, %s\n", err)
			os.Exit(1)
		}
		for _, todo := range list {
			var msg = fmt.Sprintf("%3d  %s  %s", todo.ID, todo.CreateTime, strings.TrimPrefix(todo.Description, " "))
			fmt.Fprintln(os.Stdout, msg)
		}
		os.Exit(0)
	}
}
