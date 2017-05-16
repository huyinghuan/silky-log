package bean

import (
	"silky-log/schema"

	"github.com/go-xorm/xorm"
)

type LogBean struct{}

func (logBean *LogBean) Add(logs *schema.Log) (int64, error) {
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return 0, err
	}
	return engine.Insert(logs)
}

func (logBean *LogBean) FindAll() ([]schema.Log, error) {
	var logList []schema.Log
	var engine *xorm.Engine
	var err error
	if engine, err = GetDBConenct(); err != nil {
		return nil, err
	}
	if err := engine.Find(&logList); err != nil {
		return nil, err
	}
	return logList, nil
}
