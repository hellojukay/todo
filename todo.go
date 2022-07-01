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

// Engine Engine
type Engine struct {
	path string
}

// TODO TODO
type TODO struct {
	ID          int
	Description string
	CreateTime  string
}

// Init Init
func Init() (*Engine, error) {
	var home, err = os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	var path = filepath.Join(home, ".todo", "todo.txt")
	_, err = os.Stat(filepath.Join(home, ".todo"))
	if os.IsNotExist(err) {
		// create todo home
		err = crateTodoFile(path)
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

// ListALL ListALL
func (engine *Engine) ListALL() ([]TODO, error) {
	return readTodoList(engine.path)
}

// Delete Delete
func (engine *Engine) Delete(id int) error {
	list, err := engine.ListALL()
	if err != nil {
		fmt.Printf("read todo list failed, %s\n", err)
		os.Exit(1)
	}
	var newList []TODO
	for _, todo := range list {
		if todo.ID != del {
			newList = append(newList, todo)
		}
	}
	return writeTask(engine.path, newList)
}

// Add Add
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
	return writeTask(engine.path, list)
}

// crateTodoFile crateTodoFile
func crateTodoFile(path string) error {
	var dir = filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	var list = []TODO{}
	return writeTask(path, list)
}
func readTodoList(path string) ([]TODO, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		var msg = fmt.Sprintf("can not read file %s, %s", path, err)
		return nil, errors.New(msg)
	}
	var list []TODO
	err = json.Unmarshal(jsonBytes, &list)
	if err != nil {
		var msg = fmt.Sprintf("create todolist from %s , parse json failed %s", path, err)
		return nil, errors.New(msg)
	}
	return list, nil
}

func writeTask(path string, list []TODO) error {
	jsonBytes, _ := json.Marshal(list)
	var buffer bytes.Buffer
	json.Indent(&buffer, jsonBytes, "", "\t")
	err := ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	return err
}
