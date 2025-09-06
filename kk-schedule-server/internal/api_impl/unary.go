package api_impl

import (
	"context"

	"github.com/cruvie/kk-schedule/internal/api_handlers/jobDelete"
	"github.com/cruvie/kk-schedule/internal/api_handlers/serviceList"
	"google.golang.org/grpc"

	"gitee.com/cruvie/kk_go_kit/kk_grpc"

	"github.com/cruvie/kk-schedule/internal/api_handlers/jobDisable"
	"github.com/cruvie/kk-schedule/internal/api_handlers/jobEnable"
	"github.com/cruvie/kk-schedule/internal/api_handlers/jobGet"
	"github.com/cruvie/kk-schedule/internal/api_handlers/jobList"
	"github.com/cruvie/kk-schedule/internal/api_handlers/jobPut"
	"github.com/cruvie/kk-schedule/internal/api_handlers/jobSetSpec"
	"github.com/cruvie/kk-schedule/internal/api_handlers/serviceDelete"
	"github.com/cruvie/kk-schedule/internal/api_handlers/serviceGet"
	"github.com/cruvie/kk-schedule/internal/api_handlers/servicePut"
	"github.com/cruvie/kk-schedule/kk_schedule"
)

type server struct {
	kk_schedule.UnimplementedKKScheduleServer
}

func RegisterServer(grpcServer *grpc.Server) {
	kk_schedule.RegisterKKScheduleServer(grpcServer, &server{})
}

func (x *server) JobList(ctx context.Context, input *kk_schedule.JobList_Input) (*kk_schedule.JobList_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobList.NewApi,
	)
}

func (x *server) JobGet(ctx context.Context, input *kk_schedule.JobGet_Input) (*kk_schedule.JobGet_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobGet.NewApi,
	)
}

func (x *server) JobSetSpec(ctx context.Context, input *kk_schedule.JobSetSpec_Input) (*kk_schedule.JobSetSpec_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobSetSpec.NewApi,
	)
}

func (x *server) JobEnable(ctx context.Context, input *kk_schedule.JobEnable_Input) (*kk_schedule.JobEnable_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobEnable.NewApi,
	)
}

func (x *server) JobDisable(ctx context.Context, input *kk_schedule.JobDisable_Input) (*kk_schedule.JobDisable_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobDisable.NewApi,
	)
}

func (x *server) JobPut(ctx context.Context, input *kk_schedule.JobPut_Input) (*kk_schedule.JobPut_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobPut.NewApi,
	)
}

func (x *server) JobDelete(ctx context.Context, input *kk_schedule.JobDelete_Input) (*kk_schedule.JobDelete_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		jobDelete.NewApi,
	)
}

func (x *server) ServiceList(ctx context.Context, input *kk_schedule.ServiceList_Input) (*kk_schedule.ServiceList_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		serviceList.NewApi,
	)
}

func (x *server) ServicePut(ctx context.Context, input *kk_schedule.ServicePut_Input) (*kk_schedule.ServicePut_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		servicePut.NewApi,
	)
}

func (x *server) ServiceGet(ctx context.Context, input *kk_schedule.ServiceGet_Input) (*kk_schedule.ServiceGet_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		serviceGet.NewApi,
	)
}

func (x *server) ServiceDelete(ctx context.Context, input *kk_schedule.ServiceDelete_Input) (*kk_schedule.ServiceDelete_Output, error) {
	return kk_grpc.GrpcHandler(
		ctx,
		input,
		serviceDelete.NewApi,
	)
}
