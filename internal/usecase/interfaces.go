package usecase

import (
	"avito_test_GO/internal/entity"
	"bytes"
	"context"
	"github.com/google/uuid"
	"time"
)

type (
	UserRp interface {
		CheckUserBalanceExistence(context.Context, uuid.UUID) (bool, error)
		CheckUserReserveExistence(context.Context, uuid.UUID) (bool, error)
		CheckRequiredReserveExistence(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) (bool, error)
		CreateNewBalance(context.Context, uuid.UUID, uint64) error
		CheckEnoughMoneyBalance(context.Context, uuid.UUID, uint64) (bool, error)
		CreateNewReserve(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		GetReserve(context.Context, uuid.UUID) ([]int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		UserToUserMoneyTransfer(context.Context, uuid.UUID, uuid.UUID, uint64) error
		UnreserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		GetTransactionListByDate(context.Context, uuid.UUID, uint64, uint64) ([]entity.Transaction, error)
		CheckTransactions(context.Context, uuid.UUID) (bool, error)
		GetTransactionListBySum(context.Context, uuid.UUID, uint64, uint64) ([]entity.Transaction, error)
		CheckAnyTransaction(context.Context, time.Time) (bool, error)
		GetAllTransactions(context.Context, time.Time) ([]entity.Report, error)
	}

	UserContract interface {
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		GetReserve(context.Context, uuid.UUID) ([]int64, error)
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		UserToUserMoneyTransfer(context.Context, uuid.UUID, uuid.UUID, uint64) error
		UnreserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		GetTransactionListByDate(context.Context, uuid.UUID, uint64, uint64) ([]entity.Transaction, error)
		GetTransactionListBySum(context.Context, uuid.UUID, uint64, uint64) ([]entity.Transaction, error)
		GetAllTransactions(context.Context, time.Time) ([]entity.Report, error)
	}

	ReportContract interface {
		GenerateReportByPeriod(context.Context, time.Time) (*bytes.Buffer, error)
	}
)
