package job

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/schedule"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
)

func (x *ApiJobDelete) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	err := schedule.GClient.JobDelete(x.In.ServiceName, x.In.FuncName)
	return err
}
func (x *ApiJobDisable) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	err := schedule.GClient.JobDisable(x.In.ServiceName, x.In.FuncName)
	return err
}
func (x *ApiJobEnable) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	return schedule.GClient.JobEnable(x.In.ServiceName, x.In.FuncName)
}
func (x *ApiJobGet) Service(stage *kk_stage.Stage) (*kk_schedule.PBJob, error) {
	span := stage.StartTrace("Service")
	defer span.End()

	job, err := schedule.GClient.JobGet(x.In.ServiceName, x.In.FuncName)
	return job, err
}
func (x *ApiJobList) Service(stage *kk_stage.Stage) ([]*kk_schedule.PBJob, error) {
	span := stage.StartTrace("Service")
	defer span.End()

	jobList, err := schedule.GClient.JobList(x.In.ServiceName)
	return jobList, err
}

func (x *ApiJobSetSpec) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	return schedule.GClient.JobSetSpec(x.In.ServiceName, x.In.FuncName, x.In.Spec)
}
func (x *ApiJobTrigger) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	return schedule.GClient.JobTrigger(x.In.ServiceName, x.In.FuncName)
}

func (x *ApiJobPut) Service(stage *kk_stage.Stage) error {
	span := stage.StartTrace("Service")
	defer span.End()

	err := schedule.GClient.JobPut(x.In.Job)
	if err != nil {
		return err
	}

	return nil
}
