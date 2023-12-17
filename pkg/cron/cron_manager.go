package cron

import (
	"context"
	"errors"
	"github.com/robfig/cron/v3"
	"ha/pkg/config"
	"ha/pkg/log"
	"ha/pkg/service"
)

type cronManager struct {
	service.MangerConfig
	cron     *cron.Cron
	JobEntry map[string]cron.EntryID
}

var CronCenter cronManager

func NewCronManager(configuration config.Configuration) service.CronManager {
	c := cron.New()
	CronCenter = cronManager{
		cron:     c,
		JobEntry: make(map[string]cron.EntryID)}
	return &CronCenter
}

func (c *cronManager) SetUp() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if c.Initialized {
		return
	}
	c.start()
	c.Initialized = true
}

func (c *cronManager) Shutdown() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if !c.Initialized {
		return
	}
	c.cron.Stop()
}

func (c *cronManager) start() {
	c.cron.Start()
}

func (c *cronManager) AddFunc(jobName string, spec string, cmd func()) error {
	if !c.Initialized {
		return service.UnInitialized
	}
	if jobName == "" {
		return errors.New("must have a cron job name")
	}
	if _, ok := c.JobEntry[jobName]; ok {
		return errors.New("jobName is not unique")
	}
	jobId, err := c.cron.AddFunc(spec, cmd)
	if err != nil {
		log.G(context.TODO()).Error("add cron job failed ", err)
		return err
	}
	c.JobEntry[jobName] = jobId
	return nil
}
