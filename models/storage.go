package models

import "time"

type Storage struct {
	Name        string
	Type        string
	Trades      []Trade
	Deposits    []Deposit
	Withdrawals []Withdrawal
}

type Trade struct {
	Type      string
	AssetFrom string
	AssetTo   string
	Amount    float64
	Price     float64
	Fee       float64
	Date      time.Time
}

type Deposit struct {
	Transfer
}

type Withdrawal struct {
	Transfer
	Fee float64
}

type Transfer struct {
	Asset  string
	Amount float64
	From   string
	To     string
	Date   time.Time
}


