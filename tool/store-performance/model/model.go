package model

type Performance struct {
	TargetTurnover          float64 `json:"target_turnover"`            // 目标
	CurrentTurnover         float64 `json:"current_turnover"`           // 实际
	LastTurnover            float64 `json:"last_turnover"`              // 上个月
	LastYearCurrentTurnover float64 `json:"last_year_current_turnover"` // 去年同月
	Number                  int64   `json:"number"`                     // 单数
	LastYearTurnover        float64 `json:"last_year_turnover"`         // 去年的营业额
	NowYearTurnover         float64 `json:"now_year_turnover"`          // 今年的营业额
}

type StoreRue struct {
	UnitPrice               float64 `json:"unit_price"`
	MonthOverMonth          float64 `json:"month_over_month"`           // 同比
	MonthOnMonth            float64 `json:"month_on_month"`             // 环比
	YearOverYear            float64 `json:"year_over_year"`             // 去年同比
	Completion              float64 `json:"completion"`                 // 完成率
	TargetTurnover          float64 `json:"target_turnover"`            // 目标
	CurrentTurnover         float64 `json:"current_turnover"`           // 实际
	LastTurnover            float64 `json:"last_turnover"`              // 上个月
	LastYearCurrentTurnover float64 `json:"last_year_current_turnover"` // 去年同月
	Number                  int64   `json:"number"`                     // 单数
}

type SignRue struct {
	UserName string `json:"user_name"`
}
