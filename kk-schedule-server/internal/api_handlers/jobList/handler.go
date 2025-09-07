package jobList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*kk_schedule.JobList_Output, error) {
	jobList, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &kk_schedule.JobList_Output{
		JobList: jobList,
	}, nil
}
