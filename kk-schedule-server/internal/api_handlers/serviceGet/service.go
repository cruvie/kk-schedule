package serviceGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/schedule"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) (*kk_schedule.PBRegisterService, error) {
	span := stage.StartTrace("Service")
	defer span.End()

	service, err := schedule.GClient.ServiceGet(x.In.ServiceName)
	if err != nil {
		return nil, err
	}

	return service, nil
}
