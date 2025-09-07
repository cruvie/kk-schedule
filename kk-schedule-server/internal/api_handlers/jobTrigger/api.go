package jobTrigger

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
)

type Api struct {
	*kk_grpc.DefaultApi[kk_schedule.JobTrigger_Input]
}

func NewApi() *Api {
	return &Api{
		DefaultApi: kk_grpc.NewDefaultApi[kk_schedule.JobTrigger_Input](),
	}
}
