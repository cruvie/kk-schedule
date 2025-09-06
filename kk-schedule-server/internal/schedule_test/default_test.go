package schedule_test

import (
	"testing"

	"github.com/cruvie/kk-schedule/kk_schedule"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn

func getClient(t *testing.T) kk_schedule.KKScheduleClient {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:8666",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	return kk_schedule.NewKKScheduleClient(conn)
}

func down() {
	conn.Close() // nolint
}

const testAuthToken = "sdgoisdglodshlghlshlghdlskg"

var testJob = &kk_schedule.PBRegisterJob{
	Description: "test job",
	ServiceName: "test-service",
	FuncName:    "test-func",
}

var testService = &kk_schedule.PBRegisterService{
	ServiceName: "test-service",
	Target:      "127.0.0.1:8000",
	AuthToken:   testAuthToken,
}

func TestJobList(t *testing.T) {
	defer down()
	jobs, err := getClient(t).JobList(t.Context(), &kk_schedule.JobList_Input{})
	assert.NoError(t, err)
	for _, job := range jobs.JobList {
		t.Log(job)
		t.Log(job.Prev.AsTime(), job.Next.AsTime())
	}
}

func TestJobGet(t *testing.T) {
	defer down()
	job, err := getClient(t).JobGet(t.Context(), &kk_schedule.JobGet_Input{
		ServiceName: testJob.ServiceName,
		FuncName:    testJob.FuncName,
	})
	assert.NoError(t, err)
	t.Log(job.Job)
	t.Log(job.Job.Prev.AsTime(), job.Job.Next.AsTime())
}

func TestJobSetSpec(t *testing.T) {
	defer down()
	resp, err := getClient(t).JobSetSpec(t.Context(), &kk_schedule.JobSetSpec_Input{
		ServiceName: testJob.ServiceName,
		FuncName:    testJob.FuncName,
		Spec:        "* * * * *",
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestJobEnable(t *testing.T) {
	defer down()
	resp, err := getClient(t).JobEnable(t.Context(), &kk_schedule.JobEnable_Input{
		ServiceName: testJob.ServiceName,
		FuncName:    testJob.FuncName,
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestJobDisable(t *testing.T) {
	defer down()
	resp, err := getClient(t).JobDisable(t.Context(), &kk_schedule.JobDisable_Input{
		ServiceName: testJob.ServiceName,
		FuncName:    testJob.FuncName,
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestJobPut(t *testing.T) {
	defer down()
	resp, err := getClient(t).JobPut(t.Context(), &kk_schedule.JobPut_Input{
		Job: testJob,
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestServiceList(t *testing.T) {
	defer down()
	resp, err := getClient(t).ServiceList(t.Context(), &kk_schedule.ServiceList_Input{})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestServicePut(t *testing.T) {
	defer down()
	resp, err := getClient(t).ServicePut(t.Context(), &kk_schedule.ServicePut_Input{
		Service: testService,
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestServiceGet(t *testing.T) {
	defer down()
	resp, err := getClient(t).ServiceGet(t.Context(), &kk_schedule.ServiceGet_Input{
		ServiceName: testService.ServiceName,
	})
	assert.NoError(t, err)
	t.Log(resp)
}

func TestServiceDelete(t *testing.T) {
	defer down()
	resp, err := getClient(t).ServiceDelete(t.Context(), &kk_schedule.ServiceDelete_Input{
		ServiceName: testService.ServiceName,
	})
	assert.NoError(t, err)
	t.Log(resp)
}
