package entity

import "time"

type Transaction struct {
	ServiceName   string    `json:"service_name"`
	MoneyAmount   uint64    `json:"money_amount"`
	OperationDate time.Time `json:"operation_date"`
}
