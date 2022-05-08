package models

type Asset struct {
	Name          string
	Savings       []Saving
	Price         float64
	ProfitUSD     float64
	ProfitPercent float64
}

type Saving struct {
	Asset   Asset
	Storage Storage
	Amount  float64
}
