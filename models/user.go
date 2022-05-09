package models

type User struct {
	Id            int
	Name          string
	Username      string
	Password      string
	Assets        map[string]Asset
	AUM           float64
	ProfitUSD     float64
	ProfitPercent float64
}
