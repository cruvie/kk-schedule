package api_impl

import (
	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

func RegisterFileDesc() {
	kk_grpc.GFileDescHub.RegisterFileDesc(kk_schedule.File_kk_schedule_service_proto)
}
