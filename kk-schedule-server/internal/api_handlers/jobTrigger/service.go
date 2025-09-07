package jobTrigger

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	return schedule.GClient.JobTrigger(x.In.ServiceName, x.In.FuncName)
}
