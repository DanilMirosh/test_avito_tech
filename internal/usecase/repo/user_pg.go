package repo

import (
	"avito_test_GO/internal/entity"
	"avito_test_GO/internal/usecase"
	"avito_test_GO/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserRepo struct {
	*postgres.Postgres
}

var _ usecase.UserRp = (*UserRepo)(nil)

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// add here zerolog

func (u *UserRepo) CheckUserBalanceExistence(ctx context.Context, user uuid.UUID) (bool, error) {
	query := `select exists(select * from balance where user_uuid = $1)`

	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		log.Println("Cannot execute query")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			log.Println("cannot scan exists flag")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return exists, nil
}

func (u *UserRepo) CheckUserReserveExistence(ctx context.Context, user uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT * FROM reserve WHERE user_uuid = $1)`

	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		log.Println("Cannot execute query")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			log.Println("cannot scan exists flag")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return exists, nil
}

func (u *UserRepo) CheckRequiredReserveExistence(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) (bool, error) {
	query := `SELECT check_required_reserve($1, $2, $3, $4)`
	rows, err := u.Pool.Query(ctx, query, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to check required reserve existence")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var result bool
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println("cannot scan result")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return result, nil
}

func (u *UserRepo) CheckEnoughMoneyBalance(ctx context.Context, user uuid.UUID, amount uint64) (bool, error) {
	query := `SELECT check_money($1, $2)`
	rows, err := u.Pool.Query(ctx, query, user, amount)
	if err != nil {
		log.Println("Cannot execute query to check balance")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var result bool
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println("cannot scan result")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return result, nil
}

func (u *UserRepo) CreateNewBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	query := `SELECT create_new_balance($1, $2)`
	rows, err := u.Pool.Query(ctx, query, user, sum)
	if err != nil {
		log.Println("Cannot execute query to create balance")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed CreateNewBalance query")
	return nil
}

func (u *UserRepo) CreateNewReserve(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	query := `INSERT INTO reserve(user_uuid, service_uuid, order_uuid, reserve) VALUES($1, $2, $3, $4)`
	rows, err := u.Pool.Query(ctx, query, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to create reserve")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	return nil
}

func (u *UserRepo) AppendBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	query := `SELECT update_balance($1, $2)`

	rows, err := u.Pool.Query(ctx, query, user, sum)
	if err != nil {
		log.Println("Cannot execute query to append")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed AppendBalance query")
	return nil
}

func (u *UserRepo) GetBalance(ctx context.Context, user uuid.UUID) (int64, error) {
	query := `SELECT balance FROM balance WHERE user_uuid = $1`

	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		log.Println("Cannot execute query to get user balance")
		return 0, err
	}
	defer rows.Close()
	log.Println("Successfully executed GetBalance query")

	var balance int64
	for rows.Next() {
		err = rows.Scan(&balance)
		if err != nil {
			log.Println("cannot scan balance")
			return 0, fmt.Errorf("cannot scan value")
		}
	}
	return balance, nil
}

func (u *UserRepo) GetReserve(ctx context.Context, user uuid.UUID) ([]int64, error) {
	query := `SELECT reserve FROM reserve WHERE user_uuid = $1`
	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		log.Println("Cannot execute query to get user reserve")
		return nil, err
	}
	defer rows.Close()
	log.Println("Successfully executed GetReserve query")

	var reserve []int64
	for rows.Next() {
		var rs int64
		err = rows.Scan(&rs)
		if err != nil {
			log.Println("cannot scan reserve")
			return nil, fmt.Errorf("cannot scan value")
		}
		reserve = append(reserve, rs)
	}
	return reserve, nil
}

func (u *UserRepo) ReserveMoney(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	query := `SELECT reserve_money($1, $2, $3, $4)`

	rows, err := u.Pool.Query(ctx, query, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to reserve money")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed ReserveMoney query")
	return nil
}

func (u *UserRepo) AcceptIncome(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	query := `SELECT accept_income($1, $2, $3, $4)`

	rows, err := u.Pool.Query(ctx, query, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to accept income")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed AcceptIncome query")
	return nil
}

func (u *UserRepo) UserToUserMoneyTransfer(ctx context.Context, firstUserUUID uuid.UUID, secondUserUUID uuid.UUID, amount uint64) error {
	query := `SELECT user_to_user_money_transfer($1, $2, $3)`

	rows, err := u.Pool.Query(ctx, query, firstUserUUID, secondUserUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to transfer money")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed UserToUserMoneyTransfer query")
	return nil
}

func (u *UserRepo) UnreserveMoney(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	query := `SELECT unreserve_money($1, $2, $3, $4)`

	rows, err := u.Pool.Query(ctx, query, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		log.Println("Cannot execute query to unreserve money")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed UnreserveMoney query")
	return nil
}

func (u *UserRepo) CheckTransactions(ctx context.Context, userUUID uuid.UUID) (bool, error) {
	query := `select exists (select * from report where user_uuid = $1)`
	rows, err := u.Pool.Query(ctx, query, userUUID)
	if err != nil {
		log.Println("Cannot execute query to check transactions")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var result bool
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println("cannot scan result")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return result, nil
}

func (u *UserRepo) GetTransactionListByDate(ctx context.Context, userUUID uuid.UUID, limit uint64, offset uint64) ([]entity.Transaction, error) {
	query := `select service.service_name, report.money_amount, report.operation_date from report join service on report.service_uuid = service.id where user_uuid = $1 order by operation_date LIMIT $2 OFFSET $3`

	rows, err := u.Pool.Query(ctx, query, userUUID, limit, offset)
	if err != nil {
		log.Println("Cannot execute query to get transaction list")
		return nil, fmt.Errorf("cannot scan value %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed GetTransactionListByDate query")

	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.ServiceName, &transaction.MoneyAmount, &transaction.OperationDate)
		if err != nil {
			log.Println("cannot scan transactions")
			return nil, fmt.Errorf("cannot scan value %w", err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (u *UserRepo) GetTransactionListBySum(ctx context.Context, userUUID uuid.UUID, limit uint64, offset uint64) ([]entity.Transaction, error) {
	query := `select service.service_name, report.money_amount, report.operation_date from report join service on report.service_uuid = service.id where user_uuid = $1 order by money_amount LIMIT $2 OFFSET $3`

	rows, err := u.Pool.Query(ctx, query, userUUID, limit, offset)
	if err != nil {
		log.Println("Cannot execute query to get transaction list")
		return nil, fmt.Errorf("cannot scan value %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed GetTransactionListBySum query")

	var transactions []entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.ServiceName, &transaction.MoneyAmount, &transaction.OperationDate)
		if err != nil {
			log.Println("cannot scan transactions")
			return nil, fmt.Errorf("cannot scan value %w", err)
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (u *UserRepo) CheckAnyTransaction(ctx context.Context, yearMonth time.Time) (bool, error) {
	query := `select check_transactions_by_date($1)`

	rows, err := u.Pool.Query(ctx, query, yearMonth)
	if err != nil {
		log.Println("Cannot execute query to check if any transaction exist")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var result bool
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println("cannot scan result")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return result, nil
}

func (u *UserRepo) GetAllTransactions(ctx context.Context, yearMonth time.Time) ([]entity.Report, error) {
	query := `select * from get_all_transactions($1);`

	rows, err := u.Pool.Query(ctx, query, yearMonth)
	if err != nil {
		log.Println("Cannot execute query to get all transaction list")
		return nil, fmt.Errorf("cannot scan value %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed GetAllTransactions query")

	var reports []entity.Report
	for rows.Next() {
		var report entity.Report
		err = rows.Scan(&report.ServiceName, &report.ProceedSum)
		if err != nil {
			log.Println("cannot scan transactions")
			fmt.Println(err)
			return nil, fmt.Errorf("cannot scan value %v", err)
		}
		reports = append(reports, report)
	}
	return reports, nil
}
