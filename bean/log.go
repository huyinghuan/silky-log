package bean

import (
	"silky-log/schema"
)

type LogBean struct{}

func (logBean *LogBean) Add(logs *schema.Log) (int64, error) {
	engine := GetDBconnect()
	return engine.Insert(logs)
}

func (logBean *LogBean) FindAll() ([]schema.Log, error) {
	var logList []schema.Log
	engine := GetDBconnect()
	if err := engine.Find(&logList); err != nil {
		return nil, err
	}
	return logList, nil
}
