package types

type Executable interface {
	GetName() string
	CronTime() string
	Run()
}
