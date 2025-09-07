package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cruvie/kk-schedule/kk-schedule-server/internal/g_config"
	"github.com/cruvie/kk-schedule/kk-schedule-server/kk_schedule"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	storeServiceKey = "kk-schedule-service"
	storeJobKey     = "kk-schedule-job"
)

type StoreEtcd struct {
	Client *clientv3.Client
}

func NewStoreEtcd() *StoreEtcd {
	config := clientv3.Config{
		Endpoints:   g_config.Config.StoreEtcd.Endpoints,
		Username:    g_config.Config.StoreEtcd.UserName,
		Password:    g_config.Config.StoreEtcd.Password,
		DialTimeout: 2 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	return &StoreEtcd{
		Client: client,
	}
}

func (x *StoreEtcd) getJobKey(entry *kk_schedule.PBJob) string {
	return fmt.Sprintf("%s/%s/%s", storeJobKey, entry.ServiceName, entry.FuncName)
}

func (x *StoreEtcd) JobList(serviceName string) ([]*kk_schedule.PBJob, error) {
	resp, err := x.Client.Get(context.Background(), fmt.Sprintf("%s/%s", storeJobKey, serviceName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var jobs []*kk_schedule.PBJob
	for _, kv := range resp.Kvs {
		var v kk_schedule.PBJob
		if err := json.Unmarshal(kv.Value, &v); err != nil {
			return nil, err
		}
		jobs = append(jobs, &v)
	}

	return jobs, nil
}

func (x *StoreEtcd) JobGet(serviceName, funcName string) (*kk_schedule.PBJob, error) {
	key := fmt.Sprintf("%s/%s/%s", storeJobKey, serviceName, funcName)
	resp, err := x.Client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, kk_schedule.ErrJobNotFount
	}

	var entry kk_schedule.PBJob
	if err := json.Unmarshal(resp.Kvs[0].Value, &entry); err != nil {
		return nil, err
	}

	return &entry, nil
}

func (x *StoreEtcd) JobPut(entry *kk_schedule.PBJob) error {
	key := x.getJobKey(entry)
	value, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = x.Client.Put(context.Background(), key, string(value))
	return err
}

func (x *StoreEtcd) JobDelete(serviceName, funcName string) error {
	key := fmt.Sprintf("%s/%s/%s", storeJobKey, serviceName, funcName)
	_, err := x.Client.Delete(context.Background(), key)
	return err
}

func (x *StoreEtcd) ServicePut(v *kk_schedule.PBRegisterService) error {
	key := fmt.Sprintf("%s/%s", storeServiceKey, v.ServiceName)
	value, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = x.Client.Put(context.Background(), key, string(value))
	return err
}

func (x *StoreEtcd) ServiceGet(serviceName string) (*kk_schedule.PBRegisterService, error) {
	key := fmt.Sprintf("%s/%s", storeServiceKey, serviceName)
	resp, err := x.Client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, kk_schedule.ErrServiceNotFount
	}

	var v kk_schedule.PBRegisterService
	err = json.Unmarshal(resp.Kvs[0].Value, &v)
	return &v, err
}

func (x *StoreEtcd) ServiceList() ([]*kk_schedule.PBRegisterService, error) {
	key := storeServiceKey
	resp, err := x.Client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var services []*kk_schedule.PBRegisterService
	for _, kv := range resp.Kvs {
		var v kk_schedule.PBRegisterService
		err := json.Unmarshal(kv.Value, &v)
		if err != nil {
			return nil, err
		}
		services = append(services, &v)
	}
	return services, nil
}

func (x *StoreEtcd) ServiceDelete(serviceName string) error {
	if serviceName == "" {
		return kk_schedule.ErrServiceNameEmpty
	}
	key := fmt.Sprintf("%s/%s", storeServiceKey, serviceName)
	_, err := x.Client.Delete(context.Background(), key)
	return err
}
