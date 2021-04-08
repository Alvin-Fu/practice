package service

import (
	"encoding/json"
	"practice/tool/store-performance/model"
	"testing"
)

func TestStoreService_GetStorePerformance(t *testing.T) {
	ser := NewStoreService()
	performance := &model.Performance{
		TargetTurnover:          390000,
		CurrentTurnover:         318444.4,
		LastTurnover:            311730,
		LastYearCurrentTurnover: 402686.3,
		Number:                  372,
	}
	data, _ := json.Marshal(performance)
	ser.GetStorePerformance(data)
}
