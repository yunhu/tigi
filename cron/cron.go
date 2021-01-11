package cron

import "time"

func StartCron()  {
	go Detail()
}
//任务每20秒检测一次
func Detail()  {
	ticker := time.NewTicker(time.Second * 20)
	for _=range ticker.C{
		//要做的事情

	}
}