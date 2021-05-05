package service

import (
	"practice/tool/store-performance/model"
	"strconv"
	"strings"

	"fmt"
)

type StoreService struct {
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) GetStorePerformance(performance *model.Performance) *model.StoreRue {
	return &model.StoreRue{
		UnitPrice:      s.calculateUnitPrice(performance.CurrentTurnover, performance.Number),
		Completion:     s.calculateCompletion(performance.CurrentTurnover, performance.TargetTurnover),
		MonthOnMonth:   s.calculateHuanBi(performance.CurrentTurnover, performance.LastTurnover),
		MonthOverMonth: s.calculateTongBi(performance.CurrentTurnover, performance.LastYearCurrentTurnover),
	}

}

func (s *StoreService) calculateUnitPrice(currentTurnover float64, number int64) float64 {
	return Decimal(currentTurnover / float64(number))
}

func (s *StoreService) calculateCompletion(currentTurnover, TargetTurnover float64) float64 {
	return Decimal(currentTurnover / TargetTurnover)
}

func (s *StoreService) calculateTongBi(turnover1, turnover2 float64) float64 {
	return Decimal((turnover1 - turnover2) / turnover2)
}

func (s *StoreService) calculateHuanBi(turnover1, turnover2 float64) float64 {
	return Decimal((turnover1 - turnover2) / turnover2)
}

func (s *StoreService) roundingOff(num float64) float64 {
	str := strconv.FormatFloat(num, 'f', -1, 64)
	if str == "" {
		return num
	}
	tmp := strings.Split(str, ".")
	if len(tmp) < 2 {
		return num
	}
	var f float64
	if tmp[1][2] >= '5' {
		f = 0.01
	}
	fa, err := strconv.ParseFloat(tmp[0]+"."+tmp[1][:2], 64)
	if err != nil {
		return num
	}
	return fa + f
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
