package main

import (
	"context"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"gitee.com/cruvie/kk_go_kit/kk_time"
	"github.com/cruvie/kk-schedule/internal/g_config"
	"github.com/cruvie/kk-schedule/internal/schedule"
)

var configSlog kk_stage.ConfigLog

func main() {
	{
		g_config.InitConfig()
	}
	stage := kk_stage.NewStage(context.Background(), "kk-schedule").SetStartTime(kk_time.NowUTCTime())
	{
		configSlog = kk_stage.ConfigLog{
			DebugMode:  g_config.Config.DebugMode,
			StartTime:  stage.StartTime,
			Lumberjack: kk_stage.DefaultLogConfig(kk_time.NowUTCTime(), "kk-schedule"),
		}
		configSlog.Init()
		defer configSlog.Close()
	}

	kk_stage.Print2Std("kk-schedule version: 0.0.1")

	kkServer := kk_server.NewKKServer(10*time.Second, stage)
	kkServer.Add("kk-schedule", 0, schedule.NewScheduleServer())
	kkServer.Add("kk-schedule-grpc", 0, NewGrpcServer(stage))
	kkServer.Add("kk-schedule-http", 0, NewHttpServer(stage))
	kkServer.Add("kk-schedule-web", 0, NewWebServer(stage))
	kkServer.ServeAndWait()
}
