package jobPut

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	err := schedule.GClient.JobPut(x.In.Job)
	if err != nil {
		return err
	}

	return nil
}
