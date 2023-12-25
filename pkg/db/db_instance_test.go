package db

import (
	"github.com/stretchr/testify/assert"
	"ha/pkg/config"
	"testing"
)

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
