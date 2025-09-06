package serviceGet

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*kk_schedule.ServiceGet_Output, error) {
	service, err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &kk_schedule.ServiceGet_Output{
		Service: service,
	}, nil
}
