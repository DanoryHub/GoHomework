package repository

type PSQLStub struct {
	Users []User
}

func (p PSQLStub) GetAllUsers() ([]User, error) {
	return p.Users, nil
}

func (p PSQLStub) GetAllAccounts(userID string) ([]Account, error) {
	for _, item := range p.Users {
		if item.ID == userID {
			return item.Accounts, nil
		}
	}
	return []Account{}, nil
}

func (p PSQLStub) GetAllTransactions(userID string, accountID string) ([]Transaction, error) {
	for _, user := range p.Users {
		if user.ID == userID {
			for _, account := range user.Accounts {
				if account.ID == accountID {
					return account.Transactions, nil
				}
			}
		}
	}
	return []Transaction{}, nil
}

func (p PSQLStub) GetAccountByID(userID string, accountID string) (Account, error) {
	for _, item := range p.Users {
		if item.ID == userID {
			for _, account := range item.Accounts {
				if account.ID == accountID {
					return account, nil
				}
			}
		}
	}
	return Account{}, nil
}

func (p PSQLStub) GetUserByID(userID string) (User, error) {
	for _, item := range p.Users {
		if item.ID == userID {
			return item, nil
		}
	}
	return User{}, nil
}

func (p PSQLStub) GetTransactionByID(userID string, accountID string, transactionID string) (Transaction, error) {
	for _, user := range p.Users {
		if user.ID == userID {
			for _, account := range user.Accounts {
				if account.ID == accountID {
					for _, transaction := range account.Transactions {
						if transaction.ID == transactionID {
							return transaction, nil
						}
					}
				}
			}
		}
	}
	return Transaction{}, nil
}

func (p *PSQLStub) UpdateUser(userID string, user User) (User, error) {
	for index, item := range p.Users {
		if item.ID == userID {
			p.Users = append(p.Users[:index], p.Users[index+1:]...)
			p.Users = append(p.Users, user)
			return user, nil
		}
	}
	return User{}, nil
}

func (p *PSQLStub) UpdateAccount(userID string, accountID string, account Account) (Account, error) {
	for usr_index, item := range p.Users {
		if item.ID == userID {
			for acc_index, account := range item.Accounts {
				if account.ID == accountID {
					p.Users[usr_index].Accounts = append(p.Users[usr_index].Accounts[:acc_index], p.Users[usr_index].Accounts[acc_index+1:]...)
					p.Users[usr_index].Accounts = append(p.Users[usr_index].Accounts, account)
					return account, nil
				}
			}
		}
	}
	return Account{}, nil
}

func (p *PSQLStub) UpdateTransaction(userID string, accountID string, transactionID string, transaction Transaction) (Transaction, error) {
	for usr_index, user := range p.Users {
		if user.ID == userID {
			for acc_index, account := range user.Accounts {
				if account.ID == accountID {
					for transaction_index, transact := range account.Transactions {
						if transact.ID == transactionID {
							p.Users[usr_index].Accounts[acc_index].Transactions = append(p.Users[usr_index].Accounts[acc_index].Transactions[:transaction_index],
								p.Users[usr_index].Accounts[acc_index].Transactions[transaction_index+1:]...)
							p.Users[usr_index].Accounts[acc_index].Transactions = append(p.Users[usr_index].Accounts[acc_index].Transactions, transaction)
							return transaction, nil
						}
					}
				}
			}
		}
	}
	return Transaction{}, nil
}

func (p *PSQLStub) DeleteUser(userID string) error {
	for index, item := range p.Users {
		if item.ID == userID {
			p.Users = append(p.Users[:index], p.Users[index+1:]...)
		}
	}
	return nil
}

func (p *PSQLStub) DeleteAccount(userID string, accountID string) error {
	for usr_index, item := range p.Users {
		if item.ID == userID {
			for acc_index, account := range item.Accounts {
				if account.ID == accountID {
					p.Users[usr_index].Accounts = append(p.Users[usr_index].Accounts[:acc_index], p.Users[usr_index].Accounts[acc_index+1:]...)
					return nil
				}
			}
		}
	}
	return nil
}

func (p *PSQLStub) DeleteTransaction(userID string, accountID string, transactionID string) error {
	for usr_index, user := range p.Users {
		if user.ID == userID {
			for acc_index, account := range user.Accounts {
				if account.ID == accountID {
					for transaction_index, transaction := range account.Transactions {
						if transaction.ID == transactionID {
							p.Users[usr_index].Accounts[acc_index].Transactions = append(p.Users[usr_index].Accounts[acc_index].Transactions[:transaction_index],
								p.Users[usr_index].Accounts[acc_index].Transactions[transaction_index+1:]...)
							return nil
						}
					}
				}
			}
		}
	}
	return nil
}

func (p *PSQLStub) CreateUser(user User) (User, error) {
	p.Users = append(p.Users, user)
	return user, nil
}

func (p *PSQLStub) CreateAccount(userID string, account Account) (Account, error) {
	for usr_index, item := range p.Users {
		if item.ID == userID {
			p.Users[usr_index].Accounts = append(p.Users[usr_index].Accounts, account)
			return account, nil
		}
	}
	return Account{}, nil
}

func (p *PSQLStub) CreateTransaction(userID string, accountID string, transaction Transaction) (Transaction, error) {
	for usr_index, user := range p.Users {
		if user.ID == userID {
			for acc_index, account := range user.Accounts {
				if account.ID == accountID {
					p.Users[usr_index].Accounts[acc_index].Transactions = append(p.Users[usr_index].Accounts[acc_index].Transactions, transaction)
					return transaction, nil
				}
			}
		}
	}
	return Transaction{}, nil
}
