package schedule_test

import (
	"testing"

	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestForREADME(t *testing.T) {
	// create a client for kk-schedule
	conn, err := grpc.NewClient("127.0.0.1:8666",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer conn.Close() //nolint
	if err != nil {
		t.Fatal(err)
	}
	client := kk_schedule.NewKKScheduleClient(conn)

	myServiceName := "my-service"
	testJob := &kk_schedule.PBRegisterJob{
		Description: "test job",
		ServiceName: myServiceName,
		FuncName:    "Func1",
	}
	testService := &kk_schedule.PBRegisterService{
		ServiceName: myServiceName,
		Target:      "127.0.0.1:8000",
	}
	{
		// put the running service info to kk-schedule
		resp, err := client.ServicePut(t.Context(), &kk_schedule.ServicePut_Input{
			Service: testService,
		})
		assert.NoError(t, err)
		t.Log(resp)
	}
	{
		// put a job to kk-schedule with the service name
		resp, err := client.JobPut(t.Context(), &kk_schedule.JobPut_Input{
			Job: testJob,
		})
		assert.NoError(t, err)
		t.Log(resp)
	}
	{
		// set job spec
		resp, err := client.JobSetSpec(t.Context(), &kk_schedule.JobSetSpec_Input{
			ServiceName: testJob.ServiceName,
			FuncName:    testJob.FuncName,
			Spec:        "* * * * *",
		})
		assert.NoError(t, err)
		t.Log(resp)
	}
	{
		// enable job to be triggered with the spec
		resp, err := client.JobEnable(t.Context(), &kk_schedule.JobEnable_Input{
			ServiceName: testJob.ServiceName,
			FuncName:    testJob.FuncName,
		})
		assert.NoError(t, err)
		t.Log(resp)
	}
}
