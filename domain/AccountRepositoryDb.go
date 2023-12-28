package domain

import (
	"RouterBasics/errs"
	"RouterBasics/logger"

	//"strconv"

	"gorm.io/gorm"
)

type AccountRepositoryDb struct {
	client *gorm.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	// result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	//account := Account{CustomerId: a.CustomerId, OpeningDate: a.OpeningDate, AccountType: a.AccountType, Amount: a.Amount, Status: a.Status}

	//result := d.client.Create(&account)

	if result.Error != nil {
		logger.Error("Error while creating new account: " + result.Error.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &a, nil
}

/**
 * transaction = make an entry in the transaction table + update the balance in the accounts table
 */
//  func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
// 	// starting the database transaction block
// 	tx := d.client.Begin()
// 	if tx.Error != nil {
// 		logger.Error("Error while starting a new transaction for bank account transaction: " + tx.Error.Error())
// 		return nil, errs.NewUnexpectedError("Unexpected database error")
// 	}

// 	// inserting bank account transaction
// 	result := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date)
// 											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

// 	// updating account balance
// 	if t.IsWithdrawal() {
// 		_ = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
// 	} else {
// 		updateResult = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
// 	}

// 	// in case of error Rollback, and changes from both the tables will be reverted
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error("Error while saving transaction: " + err.Error())
// 		return nil, errs.NewUnexpectedError("Unexpected database error")
// 	}
// 	// commit the transaction when all is good
// 	err := tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error("Error while commiting transaction for bank account: " + err.Error.Error())
// 		return nil, errs.NewUnexpectedError("Unexpected database error")
// 	}
// 	// getting the last transaction ID from the transaction table
// 	// transactionId, err := result.LastInsertId()
// 	// if err != nil {
// 	// 	logger.Error("Error while getting the last transaction id: " + err.Error())
// 	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
// 	// }

// 	// Getting the latest account information from the accounts table
// 	account, appErr := d.FindBy(t.AccountId)
// 	if appErr != nil {
// 		return nil, appErr
// 	}
// 	t.TransactionId = strconv.FormatInt(transactionId, 10)

// 	// updating the transaction struct with the latest balance
// 	t.Amount = account.Amount
// 	return &t, nil
// }

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	// sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	// err := d.client.Get(&account, sqlGetAccount, accountId)
	err := d.client.Where("account_id = ?", accountId).First(&account)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(db *gorm.DB) AccountRepositoryDb {

	return AccountRepositoryDb{db}
}
