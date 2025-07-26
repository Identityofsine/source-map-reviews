package cron

type Executable interface {
	GetName() string
	CronTime() string
	Run()
}
