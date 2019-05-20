package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DBRepo struct {
	DB *sql.DB
}

func DBInit(DBUserName, DBPassword string) *sql.DB {
	db, err := sql.Open("postgres", "user="+DBUserName+" password="+DBPassword+" dbname=gobankdb sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func (repo DBRepo) GetAllUsers() ([]User, error) {
	var (
		user  User
		users []User
	)

	UserData, err := repo.DB.Query("SELECT * FROM Users")
	defer UserData.Close()
	if err != nil {
		return []User{}, err
	}

	for UserData.Next() {
		err = UserData.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone)
		if err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}
	return users, err
}

func (repo DBRepo) GetAllAccounts(userID string) ([]Account, error) {
	var (
		account  Account
		accounts []Account
	)

	AccountData, err := repo.DB.Query("SELECT * FROM Accounts WHERE user_id = $1", userID)
	defer AccountData.Close()
	if err != nil {
		return []Account{}, err
	}

	for AccountData.Next() {
		err = AccountData.Scan(&account.ID, &account.UserID, &account.Balance, &account.Currency)
		if err != nil {
			return []Account{}, err
		}
		accounts = append(accounts, account)
	}
	return accounts, err

}

func (repo DBRepo) GetAllTransactions(UserId, AccountID string) ([]Transaction, error) {
	var (
		transaction  Transaction
		transactions []Transaction
	)

	TransactionData, err := repo.DB.Query("SELECT * FROM Transactions WHERE account_id = $1", AccountID)
	defer TransactionData.Close()
	if err != nil {
		return []Transaction{}, err
	}

	for TransactionData.Next() {
		err = TransactionData.Scan(&transaction.ID, &transaction.AccountID, &transaction.Operation)
		if err != nil {
			return []Transaction{}, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, err

}

func (repo DBRepo) GetAccountByID(userID, accountID string) (Account, error) {
	var account Account

	AccountData, err := repo.DB.Query("SELECT * FROM Accounts WHERE account_id = $1", accountID)
	defer AccountData.Close()
	if err != nil {
		return Account{}, err
	}
	for AccountData.Next() {
		err = AccountData.Scan(&account.ID, &account.UserID, &account.Balance, &account.Currency)
		if err != nil {
			return Account{}, err
		}

	}
	return account, err
}

func (repo DBRepo) GetUserByID(userID string) (User, error) {
	var user User

	UserData, err := repo.DB.Query("SELECT * FROM Users WHERE id = $1", userID)
	defer UserData.Close()
	if err != nil {
		return User{}, err
	}
	for UserData.Next() {
		err = UserData.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone)
		if err != nil {
			return User{}, err
		}

	}
	return user, err
}

func (repo DBRepo) GetTransactionByID(userID, accountID, transactionID string) (Transaction, error) {
	var transaction Transaction

	TransactionData, err := repo.DB.Query("SELECT * FROM Transactions WHERE transaction_id = $1", transactionID)
	defer TransactionData.Close()
	if err != nil {
		return Transaction{}, err
	}
	for TransactionData.Next() {
		err = TransactionData.Scan(&transaction.ID, &transaction.AccountID, &transaction.Operation)
		if err != nil {
			return Transaction{}, err
		}

	}
	return transaction, err
}

func (repo DBRepo) UpdateUser(userID string, user User) (User, error) {
	var out_user User
	repo.DB.Exec("DELETE FROM Accounts WHERE user_id = $1", userID)
	repo.DB.Exec("DELETE FROM Users WHERE id = $1", userID)
	repo.DB.Exec("INSERT INTO Users(id,firstname,lastname,email,telephone) VALUES($1,$2,$3,$4,$5)",
		&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone)
	UserData, err := repo.DB.Query("SELECT * FROM Users WHERE id = $1", user.ID)
	for UserData.Next() {
		err = UserData.Scan(&out_user.ID, &out_user.Firstname, &out_user.Lastname, &out_user.Email, &out_user.Phone)
	}
	return out_user, err
}

func (repo DBRepo) UpdateAccount(userID string, accountID string, account Account) (Account, error) {
	var out_account Account
	repo.DB.Exec("DELETE FROM Transactions WHERE account_id = $1", accountID)
	repo.DB.Exec("DELETE FROM Accounts WHERE account_id = $1", accountID)
	repo.DB.Exec("INSERT INTO Accounts(account_id,user_id,balance,currency) VALUES($1,$2,$3,$4)",
		&account.ID, &account.UserID, &account.Balance, &account.Currency)
	AccountData, err := repo.DB.Query("SELECT * FROM Accounts WHERE account_id = $1", account.ID)
	for AccountData.Next() {
		err = AccountData.Scan(&out_account.ID, &out_account.UserID, &out_account.Balance, &out_account.Currency)
	}
	return out_account, err

}

func (repo DBRepo) UpdateTransaction(userID string, accountID string, transactionID string, transaction Transaction) (Transaction, error) {
	var out_transaction Transaction
	repo.DB.Exec("DELETE FROM Transactions WHERE transaction_id = $1", transactionID)
	repo.DB.Exec("INSERT INTO Transactions(transaction_id, account_id, operation) VALUES($1,$2,$3)",
		&transaction.ID, &transaction.AccountID, &transaction.Operation)
	TransactionData, err := repo.DB.Query("SELECT * FROM Transactions WHERE transaction_id = $1", transaction.ID)
	for TransactionData.Next() {
		err = TransactionData.Scan(&out_transaction.ID, &out_transaction.AccountID, &out_transaction.Operation)
	}
	return out_transaction, err

}

func (repo DBRepo) DeleteUser(userID string) error {
	_, err := repo.DB.Exec("DELETE FROM Users WHERE id = $1", userID)
	return err
}

func (repo DBRepo) DeleteAccount(userID string, accountID string) error {
	_, err := repo.DB.Exec("DELETE FROM Accounts WHERE account_id = $1", accountID)
	return err
}

func (repo DBRepo) DeleteTransaction(userID string, accountID string, transactionID string) error {
	_, err := repo.DB.Exec("DELETE FROM Transactions WHERE transaction_id = $1", transactionID)
	return err
}

func (repo DBRepo) CreateUser(user User) (User, error) {
	var out_user User
	repo.DB.Exec("INSERT INTO Users(id,firstname,lastname,email,telephone) VALUES($1,$2,$3,$4,$5)",
		&user.ID, &user.Firstname, &user.Lastname, &user.Email, user.Phone)
	UserData, err := repo.DB.Query("SELECT * FROM Users WHERE id = $1", user.ID)
	for UserData.Next() {
		err = UserData.Scan(&out_user.ID, &out_user.Firstname, &out_user.Lastname, &out_user.Email, &out_user.Phone)
	}
	return out_user, err
}

func (repo DBRepo) CreateAccount(userID string, account Account) (Account, error) {
	var out_account Account
	repo.DB.Exec("INSERT INTO Accounts(account_id,user_id,balance,currency) VALUES($1,$2,$3,$4)",
		&account.ID, &account.UserID, &account.Balance, &account.Currency)
	AccountData, err := repo.DB.Query("SELECT * FROM Accounts WHERE account_id = $1", account.ID)
	for AccountData.Next() {
		err = AccountData.Scan(&out_account.ID, &out_account.UserID, &out_account.Balance, &out_account.Currency)
	}
	return out_account, err
}

func (repo DBRepo) CreateTransaction(userID string, accountID string, transaction Transaction) (Transaction, error) {
	var out_transaction Transaction
	repo.DB.Exec("INSERT INTO Transactions(transaction_id, account_id, operation) VALUES($1,$2,$3)",
		&transaction.ID, &transaction.AccountID, &transaction.Operation)
	TransactionData, err := repo.DB.Query("SELECT * FROM Transactions WHERE transaction_id = $1", transaction.ID)
	for TransactionData.Next() {
		err = TransactionData.Scan(&out_transaction.ID, &out_transaction.AccountID, &out_transaction.Operation)
	}
	return out_transaction, err

}
