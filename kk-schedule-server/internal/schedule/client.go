package schedule

import (
	"errors"
	"slices"
	"strings"

	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var GClient *Client

type Client struct {
	cron   *cron.Cron
	storer StoreDriver
}

func InitGClient(cfg *Config) {
	cfg.check()
	c := &Client{
		cron:   cron.New(cfg.Opts...),
		storer: cfg.StoreDriver,
	}
	GClient = c
	GClient.initJob()
}

func (x *Client) initJob() {
	jobList, err := x.JobList("")
	if err != nil {
		panic(err)
	}
	for _, job := range jobList {
		if !job.Enabled {
			continue
		}
		err := x.JobEnable(job.ServiceName, job.FuncName)
		if err != nil {
			panic(err)
		}
	}
}

func (x *Client) Start() {
	x.cron.Start()
}

func (x *Client) Close() {
	x.cron.Stop()
}

func (x *Client) JobPut(jobs ...*kk_schedule.PBRegisterJob) error {
	for _, job := range jobs {
		err := job.Check()
		if err != nil {
			panic(err)
		}
		{ // check service exist
			_, err = x.storer.ServiceGet(job.ServiceName)
			if err != nil {
				return err
			}
		}
		newEntry := &kk_schedule.PBJob{
			EntryID:     0,
			Enabled:     false,
			Next:        nil,
			Prev:        nil,
			Spec:        "",
			ServiceName: job.ServiceName,
			Description: job.Description,
			FuncName:    job.FuncName,
		}

		entry, err := x.storer.JobGet(job.ServiceName, job.FuncName)
		if err != nil && !errors.Is(err, kk_schedule.ErrJobNotFount) {
			return err
		}

		if entry != nil {
			newEntry.EntryID = entry.EntryID
			newEntry.Enabled = entry.Enabled
			newEntry.Next = entry.Next
			newEntry.Prev = entry.Prev
			newEntry.Spec = entry.Spec
		}
		err = x.storer.JobPut(newEntry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (x *Client) JobList(serviceName string) ([]*kk_schedule.PBJob, error) {
	entries := x.cron.Entries()
	entryList, err := x.storer.JobList(serviceName)
	if err != nil {
		return nil, err
	}
	var hasSpecEntryList []int32
	var pbJobs []*kk_schedule.PBJob
	for _, entry := range entries {
		find, b := lo.Find(entryList, func(item *kk_schedule.PBJob) bool {
			return item.EntryID == int32(entry.ID)
		})
		if !b {
			continue
		} else {
			hasSpecEntryList = append(hasSpecEntryList, find.EntryID)
		}
		pbJobs = append(pbJobs, &kk_schedule.PBJob{
			EntryID:     int32(entry.ID),
			Enabled:     find.Enabled,
			Next:        timestamppb.New(entry.Next),
			Prev:        timestamppb.New(entry.Prev),
			Spec:        find.Spec,
			Description: find.Description,
			FuncName:    find.FuncName,
			ServiceName: find.ServiceName,
		})
	}

	noSpecJobList := lo.Filter(entryList, func(item *kk_schedule.PBJob, index int) bool {
		_, b := lo.Find(hasSpecEntryList, func(id int32) bool {
			return id == item.EntryID
		})
		return !b
	})
	pbJobs = append(pbJobs, noSpecJobList...)
	// sort by serviceName
	slices.SortFunc(pbJobs, func(a, b *kk_schedule.PBJob) int {
		return strings.Compare(a.ServiceName, b.ServiceName)
	})
	return pbJobs, nil
}

func (x *Client) JobGet(serviceName, funcName string) (*kk_schedule.PBJob, error) {
	entry, err := x.storer.JobGet(serviceName, funcName)
	if err != nil {
		return nil, err
	}
	cEntry := x.cron.Entry(cron.EntryID(entry.EntryID))

	if cEntry.Valid() {
		entry.Next = timestamppb.New(cEntry.Next)
		entry.Prev = timestamppb.New(cEntry.Prev)
	}

	return entry, nil
}

func (x *Client) JobSetSpec(serviceName, funcName string, spec string) error {
	job, err := x.storer.JobGet(serviceName, funcName)
	if err != nil {
		return err
	}
	if job.Spec == spec {
		return nil
	}
	job.Spec = spec

	err = x.storer.JobPut(job)
	if job.Enabled {
		err = x.JobEnable(serviceName, funcName)
		if err != nil {
			return err
		}
	}

	return err
}

func (x *Client) JobEnable(serviceName string, funcName string) error {
	entry, err := x.storer.JobGet(serviceName, funcName)
	if err != nil {
		return err
	}
	if entry.Spec == "" {
		return kk_schedule.ErrSpecIsEmpty
	}
	err = x.JobDisable(serviceName, funcName)
	if err != nil {
		return err
	}
	service, err := x.storer.ServiceGet(serviceName)
	if err != nil {
		return err
	}

	entryID, err := x.cron.AddFunc(entry.Spec, triggerFunc(service, funcName))
	if err != nil {
		return err
	}

	entry.EntryID = int32(entryID)
	entry.Enabled = true

	err = x.storer.JobPut(entry)
	return err
}

func (x *Client) JobDisable(serviceName, funcName string) error {
	entry, err := x.storer.JobGet(serviceName, funcName)
	if err != nil {
		return err
	}
	x.cron.Remove(cron.EntryID(entry.EntryID))

	entry.Enabled = false
	entry.EntryID = 0

	err = x.storer.JobPut(entry)
	return err
}

func (x *Client) JobDelete(serviceName, funcName string) error {
	// disable job
	err := x.JobDisable(serviceName, funcName)
	if err != nil {
		return err
	}
	return x.storer.JobDelete(serviceName, funcName)
}

// JobTrigger triggers a job manually
func (x *Client) JobTrigger(serviceName, funcName string) error {
	service, err := x.storer.ServiceGet(serviceName)
	if err != nil {
		return err
	}

	// Trigger the job function directly
	triggerFunc(service, funcName)()
	return nil
}

func (x *Client) ServiceList() ([]*kk_schedule.PBRegisterService, error) {
	return x.storer.ServiceList()
}

func (x *Client) ServicePut(service *kk_schedule.PBRegisterService) error {
	return x.storer.ServicePut(service)
}

func (x *Client) ServiceGet(serviceName string) (*kk_schedule.PBRegisterService, error) {
	return x.storer.ServiceGet(serviceName)
}

func (x *Client) ServiceDelete(serviceName string) error {
	// check no job in service
	jobList, err := x.JobList(serviceName)
	if err != nil {
		return err
	}
	if len(jobList) > 0 {
		return kk_schedule.ErrServiceHasJob
	}
	return x.storer.ServiceDelete(serviceName)
}
