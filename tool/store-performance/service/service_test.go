package service

import (
	"practice/tool/store-performance/model"
	"strconv"
	"strings"
	"testing"

	"fmt"
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
	//data, _ := json.Marshal(performance)
	mo := ser.GetStorePerformance(performance)
	t.Logf("%#v", mo)
}

func TestNewStoreService(t *testing.T) {
	str := strconv.FormatFloat(0.8165241025641026, 'f', -1, 64)
	if str == "" {
		return
	}
	t.Log(str)
	tmp := strings.Split(str, ".")
	if len(tmp) < 2 {
		return
	}

	t.Log(tmp[0], tmp[1][:3])
	fa, err := strconv.ParseFloat(tmp[0]+"."+tmp[1][:2], 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(fa)
	if tmp[1][2] >= '5' {
		fa += 0.01
	}
	t.Log(fa)
}

func TestDecimal(t *testing.T) {
	t.Log(Decimal(0.8165241025641026))
}
