package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
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
		fmt.Errorf("can not init todolist, %s\n", err)
		os.Exit(1)
	}
}
func main() {
	if add {
		if msg == "" {
			fmt.Errorf("use -m provide descript for new task")
			fmt.Errorf("try run -a -m \"task description\"")
			os.Exit(1)
		}
		err := engine.Add(msg)
		if err != nil {
			fmt.Errorf("create a new task , %s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if del != 0 {
		err := engine.Delete(del)
		if err != nil {
			fmt.Errorf("delete task, %s", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if list {
		list, err := engine.ListALL()
		if err != nil {
			fmt.Errorf("read todo list failed, %s\n", err)
			os.Exit(1)
		}
		var w = tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', tabwriter.AlignRight)
		for _, todo := range list {
			var msg = fmt.Sprintf("%d\t%s\t%-s\t", todo.ID, todo.CreateTime, todo.Description)
			fmt.Fprintln(w, msg)
		}
		w.Flush()
		os.Exit(0)
	}
}
