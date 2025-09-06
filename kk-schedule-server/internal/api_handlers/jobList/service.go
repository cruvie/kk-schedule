package jobList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/internal/schedule"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) ([]*kk_schedule.PBJob, error) {
	span := stage.StartTrace("Service")
	defer span.End()

	jobList, err := schedule.GClient.JobList(x.In.ServiceName)
	return jobList, err
}
