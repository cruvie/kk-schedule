package schedule_test

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"testing"

	"github.com/cruvie/kk-schedule/kk_schedule"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	kk_schedule.UnimplementedKKScheduleTriggerServer
}

func (server) Trigger(ctx context.Context, input *kk_schedule.Trigger_Input) (*kk_schedule.Trigger_Output, error) {
	slog.Info("Trigger received", "FuncName", input.FuncName)
	switch input.FuncName {
	case "Func1":
		go Func1()
	default:
		return nil, kk_schedule.ErrJobNotFount
	}
	return &kk_schedule.Trigger_Output{}, nil
}

func Func1() {
	slog.Info("Func1 start")
	defer slog.Info("Func1 end")
}

func authorityAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	var authority string
	if auth := md[":authority"]; len(auth) > 0 {
		authority = auth[0]
	} else {
		return nil, status.Error(codes.Unauthenticated, "missing authority")
	}

	if authority != testAuthToken {
		return nil, status.Error(codes.Unauthenticated, "invalid authority")
	}

	return handler(ctx, req)
}

func TestClientServer(t *testing.T) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authorityAuthInterceptor,
		),
	)
	defer grpcServer.GracefulStop()
	kk_schedule.RegisterKKScheduleTriggerServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
