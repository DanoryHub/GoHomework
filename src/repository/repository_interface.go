package repository

type (
	UserRepository interface {
		GetAllUsers() ([]User, error)
		GetAllAccounts(userId string) ([]Account, error)
		FindAccountById(userID string, accountId string) (Account, error)
		FindUserByID(userId string) (User, error)
		UpdateUser(userId string, user User) (User, error)
		UpdateAccount(userID string, accountId string, account Account) (Account, error)
		DeleteUser(userId string) error
		DeleteAccount(userId string, accountId string) error
	}

	User struct {
		ID string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Email string `json:"email"`
		SMS string `json:"sms"`
		Accounts []Account `json:"accounts"`
	}

	Account struct{
		ID string `json:"id"`
		UserID string `json:"user_id"`
		Balance string `json:"balance"`
		Currency string `json:"currency"`
	}
)