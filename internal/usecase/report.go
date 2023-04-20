package usecase

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"
)

type ReportUseCase struct {
	user UserContract
}

func NewReportUseCase(user UserContract) *ReportUseCase {
	return &ReportUseCase{
		user: user,
	}
}

func (r *ReportUseCase) GenerateReportByPeriod(ctx context.Context, period time.Time) (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	transactionList, err := r.user.GetAllTransactions(ctx, period)
	if err != nil {
		return nil, err
	}
	if transactionList == nil || len(transactionList) == 0 {
		return nil, fmt.Errorf("transaction list is nil or zero len")
	}
	reportTemplate := [][]string{
		{"Название сервиса", "Сумма"},
	}
	for _, transaction := range transactionList {
		reportRow := []string{transaction.ServiceName, strconv.FormatInt(transaction.ProceedSum, 10)}
		reportTemplate = append(reportTemplate, reportRow)
	}
	w := csv.NewWriter(b)
	defer w.Flush()
	for _, record := range reportTemplate {
		if err := w.Write(record); err != nil {
			return nil, err
		}
	}
	return b, nil
}
