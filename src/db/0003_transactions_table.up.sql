CREATE TABLE Transactions(
transaction_id SERIAL PRIMARY KEY,
account_id INTEGER REFERENCES Accounts(account_id),
operation VARCHAR(50)
); 
