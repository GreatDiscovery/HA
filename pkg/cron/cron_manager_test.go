package cron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ha/pkg/config"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	configuration := config.Configuration{}
	cronManager := NewCronManager(configuration)
	cronManager.SetUp()
	count := 0
	CronCenter.AddFunc("@every 5s", "@every 5s", func() {
		fmt.Println(time.Now())
		count++
	})
	time.Sleep(11 * time.Second)
	assert.Truef(t, count == 2, "execute twice")
	fmt.Println(CronCenter.JobEntry)
	cronManager.Shutdown()
}
