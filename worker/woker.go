package worker

import "fmt"

//任务 类型
type Task struct {
	taskId int
	f func(interface{} ) error
}

//创建任务
func NewTask(id int,f func(interface{}) error) *Task{
	return &Task{
		taskId: id,
		f:f,
	}
}

//协程池结构
type Pool struct {
	workerNum int
	EntryChan chan *Task
	workerChan chan *Task
}
//初始化
func NewPool(num int) *Pool {
	return &Pool{
		workerNum: num,
		EntryChan: make(chan *Task),
		workerChan: make(chan *Task),
	}
}

//worker执行
func (p *Pool) worker(id int)  {
	for task:=range p.workerChan{
		task.f(task.taskId)
		fmt.Println("worker ID:",id,"taskID:",task.taskId,"is done")
	}
}
//协程池启动
func (p *Pool) Run()  {
	for i:=0; i<p.workerNum;i++{
		go p.worker(i)
		fmt.Println("workerID",i,"已经启动！")
	}

	for task :=range p.EntryChan{
		p.workerChan <- task
	}

}
