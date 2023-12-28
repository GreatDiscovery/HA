package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"ha/pkg/config"
	"ha/pkg/db/model"
	"ha/pkg/db/query"
	"testing"
	"time"
)

func init() {
	config := config.Configuration{
		MySQLHostName: "127.0.0.1",
		MySQLPort:     3306,
		MySQLUser:     "root",
		MySQLPassword: "123456",
		MySQLDB:       "test",
	}
	err := SetUp(config)
	if err != nil {
		panic(err)
	}
}

func TestSetUp(t *testing.T) {
	type args struct {
		config config.Configuration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{config: config.Configuration{
				MySQLHostName: "127.0.0.1",
				MySQLPort:     3306,
				MySQLUser:     "root",
				MySQLPassword: "123456",
				MySQLDB:       "test",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetUp(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("SetUp() error = %v, wantErr %v", err, tt.wantErr)
			}

			var result int
			tx := DB.Raw("select 1").Scan(&result)
			tx.Commit()
			assert.True(t, result == 1)
		})
	}
}

func TestDBInstance(t *testing.T) {
	query.SetDefault(DB)
	instance := model.DatabaseInstance{
		Hostname:            "127.0.0.1",
		Port:                3306,
		LastChecked:         time.Now(),
		LastSeen:            time.Now(),
		ServerID:            0,
		Version:             "",
		BinlogFormat:        "",
		LogBin:              0,
		LogSlaveUpdates:     0,
		BinaryLogFile:       "",
		BinaryLogPos:        0,
		MasterHost:          "",
		MasterPort:          0,
		SlaveSqlRunning:     0,
		SlaveIoRunning:      0,
		MasterLogFile:       "",
		ReadMasterLogPos:    0,
		RelayMasterLogFile:  "",
		ExecMasterLogPos:    0,
		SecondsBehindMaster: 0,
		SlaveLagSeconds:     0,
		NumSlaveHosts:       0,
		SlaveHosts:          "",
		ClusterName:         "",
		CreateTime:          time.Now(),
		UpdateTime:          time.Now(),
	}
	err := query.DatabaseInstance.WithContext(context.Background()).Create(&instance)
	if err != nil {
		panic(err)
	}

	find, err := query.DatabaseInstance.WithContext(context.Background()).Where(query.DatabaseInstance.Hostname.Eq("127.0.0.1")).Where(query.DatabaseInstance.Port.Eq(3306)).First()
	if err != nil {
		panic(err)
	}
	fmt.Printf("find=%+v\n", find)
	info, err := query.DatabaseInstance.WithContext(context.Background()).Delete(find)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows.affect=%v\n", info)
}
