package serviceList

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

type Api struct {
	*kk_grpc.DefaultApi[kk_schedule.ServiceList_Input]
}

func NewApi() *Api {
	return &Api{
		DefaultApi: kk_grpc.NewDefaultApi[kk_schedule.ServiceList_Input](),
	}
}
