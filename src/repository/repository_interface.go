package repository

type (
	UserRepository interface {
		GetAllUsers() ([]User, error)
		GetAllAccounts(userID string) ([]Account, error)
		GetAllTransactions(userID string, accountID string) ([]Transaction, error)
		GetAccountByID(userID string, accountId string) (Account, error)
		GetUserByID(userID string) (User, error)
		GetTransactionByID(userID string, accountID string, transactionID string) (Transaction, error)
		UpdateUser(userID string, user User) (User, error)
		UpdateAccount(userID string, accountID string, account Account) (Account, error)
		UpdateTransaction(userID string, accountID string, transactionID string, transaction Transaction) (Transaction, error)
		DeleteUser(userID string) error
		DeleteAccount(userID string, accountID string) error
		DeleteTransaction(userID string, accountID string, transactionID string) error
		CreateUser(user User) (User, error)
		CreateAccount(userID string, account Account) (Account, error)
		CreateTransaction(userID string, accountID string, transaction Transaction) (Transaction, error)
	}

	User struct {
		ID string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		Accounts []Account `json:"accounts"`
	}

	Account struct{
		ID string `json:"id"`
		UserID string `json:"user_id"`
		Balance string `json:"balance"`
		Currency string `json:"currency"`
		Transactions []Transaction `json:"transactions"`
	}

	Transaction struct{
		ID string `json:"id"`
		AccountID string `json:"account_id"`
		Operation string `json:"operation"`
	}
)