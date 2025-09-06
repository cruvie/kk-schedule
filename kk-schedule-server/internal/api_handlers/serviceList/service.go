package serviceList

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/internal/schedule"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

func (x *Api) Service(stage *kk_stage.Stage) ([]*kk_schedule.PBRegisterService, error) {
	span := stage.StartTrace("Service")
	defer span.End()

	service, err := schedule.GClient.ServiceList()
	return service, err
}
