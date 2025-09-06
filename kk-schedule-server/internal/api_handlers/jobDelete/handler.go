package jobDelete

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

func (x *Api) Handler(stage *kk_stage.Stage) (*kk_schedule.JobDelete_Output, error) {
	err := x.Service(stage)
	if err != nil {
		return nil, err
	}
	return &kk_schedule.JobDelete_Output{}, nil
}
