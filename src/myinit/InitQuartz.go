package myinit

import (
	"github.com/bamzi/jobrunner"
	"github_com_Taoey_iris_cli/src/system/job"
)

func InitQuartz() {
	jobrunner.Start()
	jobrunner.Schedule("@every 2s", job.PrintTime{})
}
