package service

import (
	"encoding/json"
	"practice/tool/store-performance/model"
)

type StoreService struct {
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) GetStorePerformance(performance *model.Performance) ([]byte, error) {

	rue := &model.StoreRue{
		UnitPrice:      s.calculateUnitPrice(performance.CurrentTurnover, performance.Number),
		Completion:     s.calculateCompletion(performance.CurrentTurnover, performance.TargetTurnover),
		MonthOnMonth:   s.calculateHuanBi(performance.CurrentTurnover, performance.LastTurnover),
		MonthOverMonth: s.calculateTongBi(performance.CurrentTurnover, performance.LastYearCurrentTurnover),
	}

	return json.Marshal(rue)
}

func (s *StoreService) calculateUnitPrice(currentTurnover float64, number int64) float64 {
	return currentTurnover / float64(number)
}

func (s *StoreService) calculateCompletion(currentTurnover, TargetTurnover float64) float64 {
	return currentTurnover / TargetTurnover
}

func (s *StoreService) calculateTongBi(turnover1, turnover2 float64) float64 {
	return (turnover1 - turnover2) / turnover2
}

func (s *StoreService) calculateHuanBi(turnover1, turnover2 float64) float64 {
	return (turnover1 - turnover2) / turnover2
}
