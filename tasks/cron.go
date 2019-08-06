package tasks

import (
	"github.com/robfig/cron"
)

// Cron 定时器单例
var Cron *cron.Cron

//Run 运行
// func Run(job func() error) {

// 	form := time.Now().UnixNano()
// 	err := job()
// 	to := time.Now().UnixNano()
// 	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
// 	if err != nil {
// 		fmt.Printf("%s error：%dms \n", jobName, (to-form)/int64(time.Millisecond))

// 	} else {
// 		fmt.Printf("%s success: %dms \n", jobName, (to-form)/int64(time.Millisecond))
// 	}
// }

// CronJob 定时任务
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	Cron.AddFunc("0 0 0 * * *", func() { RestartDailyRank() })
	Cron.Start()

}
