package worker

import "fmt"

var WP *Pool

type TaskData struct {
	TaskId int
}

//任务 类型
type Task struct {
	Data *TaskData
	F    func(data *TaskData) error
}

//创建任务
func NewTask(data *TaskData, f func(*TaskData) error) *Task {
	return &Task{
		Data: data,
		F:    f,
	}
}

//协程池结构
type Pool struct {
	workerNum  int
	workerChan chan *Task
}

//初始化
func NewPool(num int) *Pool {
	WP = &Pool{
		workerNum:  num,
		workerChan: make(chan *Task),
	}
	return WP
}

//worker执行
func (p *Pool) worker(id int) {
	for task := range p.workerChan {
		task.F(task.Data)
		fmt.Println("worker ID:", id, "taskID:", task.Data, "is done")
	}
}
func (p *Pool) SengToWorker(task *Task) {
	p.workerChan <- task
	return

}

//协程池启动
func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
		fmt.Println("workerID", i, "已经启动！")
	}
}
