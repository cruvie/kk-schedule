package schedule

import (
	"context"
	"log/slog"

	"gitee.com/cruvie/kk_go_kit/kk_grpc"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func triggerClient(service *kk_schedule.PBRegisterService) (conn *grpc.ClientConn, client kk_schedule.KKScheduleTriggerClient, err error) {
	var opts []grpc.DialOption
	if service.AuthToken != "" {
		opts = append(opts, grpc.WithAuthority(service.AuthToken))
	}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err = grpc.NewClient(service.Target, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, kk_schedule.NewKKScheduleTriggerClient(conn), nil
}

func triggerFunc(service *kk_schedule.PBRegisterService, funcName string) func() {
	return func() {
		conn, client, err := triggerClient(service)
		defer func() {
			err := conn.Close()
			if err != nil {
				slog.Error(err.Error())
			}
		}()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		stage := kk_stage.NewStage(context.Background(), "kk-schedule")
		ctx, cancelFunc := kk_grpc.NewCallGrpcCtx(stage)
		defer cancelFunc()
		_, err = client.Trigger(ctx, &kk_schedule.Trigger_Input{
			FuncName: funcName,
		})
		if err != nil {
			slog.Error(err.Error())
		}
	}
}
