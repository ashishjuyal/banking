package domain

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

func Test_should_rollback_transaction_in_case_of_failure(t *testing.T) {
	mockDB, mocksql, _ := sqlmock.New()
	defer mockDB.Close()

	//create a db client for account repository
	client := sqlx.NewDb(mockDB, "sqlmock")
	accountRepo := AccountRepositoryDb{client}

	// fake transaction instance
	trans := Transaction{"fakeTransId", "fakeAccountId", 10, WITHDRAWAL, "fakeTransDate"}

	// setting expectations and mock objects
	mocksql.ExpectBegin()

	mocksql.ExpectExec(`INSERT INTO transactions`).
		WithArgs(trans.AccountId, trans.Amount, trans.TransactionType, trans.TransactionDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mocksql.ExpectExec("UPDATE accounts").
		WithArgs(trans.Amount, trans.AccountId).
		WillReturnError(fmt.Errorf("some error"))

	mocksql.ExpectRollback()

	// now we execute our method
	if _, appError := accountRepo.SaveTransaction(trans); appError == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mocksql.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
