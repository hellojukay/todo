package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Engine struct {
	path string
}
type TODO struct {
	ID          int
	Description string
	CreateTime  string
}

func Init() (*Engine, error) {
	var home, err = os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	var path = filepath.Join(home, ".todo", "todo.txt")
	_, err = os.Stat(filepath.Join(home, ".todo"))
	if os.IsNotExist(err) {
		// create todo home
		err = crate_todo_file(path)
		if err != nil {
			fmt.Printf("init todolist %s\n", err)
			os.Exit(1)
		}
	}
	engine := Engine{
		path: path,
	}
	return &engine, nil
}
func (engine *Engine) ListALL() ([]TODO, error) {
	return read_todo_list(engine.path)
}
func (engine *Engine) Delete(id int) error {
	list, err := engine.ListALL()
	if err != nil {
		fmt.Errorf("read todo list failed, %s\n", err)
		os.Exit(1)
	}
	var newList []TODO
	for _, todo := range list {
		if todo.ID != del {
			newList = append(newList, todo)
		}
	}
	return write_task(engine.path, newList)
}
func (engine *Engine) Add(msg string) error {
	list, err := engine.ListALL()
	if err != nil {
		var msg = fmt.Sprintf("read todo list failed, %s\n", err)
		return errors.New(msg)
	}
	id := 0
	for _, task := range list {
		if id < task.ID {
			id = task.ID
		}
	}
	id = id + 1
	var t = TODO{
		ID:          id,
		CreateTime:  time.Now().Format("2006-01-02"),
		Description: msg,
	}
	list = append(list, t)
	return write_task(engine.path, list)
}
func crate_todo_file(path string) error {
	var dir = filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	var list = []TODO{}
	return write_task(path, list)
}
func read_todo_list(path string) ([]TODO, error) {
	json_bytes, err := ioutil.ReadFile(path)
	if err != nil {
		var msg = fmt.Sprintf("can not read file %s, %s", path, err)
		return nil, errors.New(msg)
	}
	var list []TODO
	err = json.Unmarshal(json_bytes, &list)
	if err != nil {
		var msg = fmt.Sprintf("create todolist from %s , parse json failed %s", path, err)
		return nil, errors.New(msg)
	}
	return list, nil
}
func write_task(path string, list []TODO) error {
	json_bytes, _ := json.Marshal(list)
	var buffer bytes.Buffer
	json.Indent(&buffer, json_bytes, "", "\t")
	err := ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	return err
}
