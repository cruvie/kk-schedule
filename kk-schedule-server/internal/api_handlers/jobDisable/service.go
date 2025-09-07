package jobDisable

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	err := schedule.GClient.JobDisable(x.In.ServiceName, x.In.FuncName)
	return err
}
