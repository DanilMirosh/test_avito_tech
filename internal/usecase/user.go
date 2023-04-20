package usecase

import (
	"avito_test_GO/internal/entity"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserUseCase struct {
	repo UserRp
}

var _ UserContract = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRp) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) AppendBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	if sum <= 0 {
		log.Println("Cannot work with this sum number")
		return fmt.Errorf("error in sum value")
	}
	exists, err := u.repo.CheckUserBalanceExistence(ctx, user)
	if err != nil {
		return err
	}
	if !exists {
		return u.repo.CreateNewBalance(ctx, user, sum)
	}
	return u.repo.AppendBalance(ctx, user, sum)
}

func (u *UserUseCase) GetBalance(ctx context.Context, user uuid.UUID) (int64, error) {
	exists, err := u.repo.CheckUserBalanceExistence(ctx, user)
	if err != nil {
		return 0, err
	}
	if !exists {
		var sum uint64
		err = u.repo.CreateNewBalance(ctx, user, sum)
		if err != nil {
			return 0, err
		}
	}
	return u.repo.GetBalance(ctx, user)
}

func (u *UserUseCase) GetReserve(ctx context.Context, user uuid.UUID) ([]int64, error) {
	exists, err := u.repo.CheckUserReserveExistence(ctx, user)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user does not have any reserved maney %w", err)
	}
	return u.repo.GetReserve(ctx, user)
}

func (u *UserUseCase) ReserveMoney(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	enough, err := u.repo.CheckEnoughMoneyBalance(ctx, userUUID, amount)
	if err != nil {
		return err
	}
	if !enough {
		return fmt.Errorf("not enough money %w", err)
	}
	return u.repo.ReserveMoney(ctx, userUUID, serviceUUID, orderUUID, amount)
}

func (u *UserUseCase) AcceptIncome(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	exists, err := u.repo.CheckRequiredReserveExistence(ctx, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("required reserve does not exist %w", err)
	}
	return u.repo.AcceptIncome(ctx, userUUID, serviceUUID, orderUUID, amount)
}

func (u *UserUseCase) UserToUserMoneyTransfer(ctx context.Context, firstUserUUID uuid.UUID, secondUserUUID uuid.UUID, amount uint64) error {
	firstUserExists, err := u.repo.CheckUserBalanceExistence(ctx, firstUserUUID)
	if err != nil {
		return err
	}
	if !firstUserExists {
		return fmt.Errorf("sender user does not have balance %w", err)
	}
	secondUserExists, err := u.repo.CheckUserBalanceExistence(ctx, secondUserUUID)
	if err != nil {
		return err
	}
	if !secondUserExists {
		fmt.Println("Receiver user balance created. Try to transfer money again")
		return u.repo.CreateNewBalance(ctx, secondUserUUID, 0)
	}
	if amount == 0 {
		return fmt.Errorf("transfer amount must not be zero %w", err)
	}
	return u.repo.UserToUserMoneyTransfer(ctx, firstUserUUID, secondUserUUID, amount)
}

func (u *UserUseCase) UnreserveMoney(ctx context.Context, userUUID uuid.UUID, serviceUUID uuid.UUID, orderUUID uuid.UUID, amount uint64) error {
	reserveExists, err := u.repo.CheckRequiredReserveExistence(ctx, userUUID, serviceUUID, orderUUID, amount)
	if err != nil {
		return err
	}
	if !reserveExists {
		return fmt.Errorf("requested reserve does not exist")
	}
	balanceExists, err := u.repo.CheckUserBalanceExistence(ctx, userUUID)
	if err != nil {
		return err
	}
	if !balanceExists {
		fmt.Println("User balance created. Try to unreserve money again")
		return u.repo.CreateNewBalance(ctx, userUUID, 0)
	}
	return u.repo.UnreserveMoney(ctx, userUUID, serviceUUID, orderUUID, amount)
}

func (u *UserUseCase) GetTransactionListByDate(ctx context.Context, userUUID uuid.UUID, limit uint64, offset uint64) ([]entity.Transaction, error) {
	exists, err := u.repo.CheckTransactions(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user does not have any transactions %w", err)
	}
	return u.repo.GetTransactionListByDate(ctx, userUUID, limit, offset)
}

func (u *UserUseCase) GetTransactionListBySum(ctx context.Context, userUUID uuid.UUID, limit uint64, offset uint64) ([]entity.Transaction, error) {
	exists, err := u.repo.CheckTransactions(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user does not have any transactions %w", err)
	}
	return u.repo.GetTransactionListBySum(ctx, userUUID, limit, offset)
}

func (u *UserUseCase) GetAllTransactions(ctx context.Context, yearMonth time.Time) ([]entity.Report, error) {
	exists, err := u.repo.CheckAnyTransaction(ctx, yearMonth)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("report does not have any transactions %w", err)
	}
	return u.repo.GetAllTransactions(ctx, yearMonth)
}
