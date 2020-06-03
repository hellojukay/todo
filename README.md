# todo
a command line todolist write in go.

![image](https://travis-ci.org/hellojukay/todo.svg?branch=master)
# usage
```shell
hellojukay@local todo (master) $ todo --help
Usage of todo:
  -a	add a new task to todolist
  -d int
    	delete task from todolist
  -list
    	list todo
  -m string
    	task description
```
list tasks
```shell
hellojukay@local todo (master) $ todo --list
1  2020-06-03  模拟镜像打包 
2  2020-06-03  k8s集群搭建
3  2020-06-03  申请新的 k8s 集群节点
4  2020-06-03  无法用 sudo 切换到 root 的问题
```
