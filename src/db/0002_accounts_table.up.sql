CREATE TABLE Accounts(
account_id SERIAL PRIMARY KEY,
user_id INTEGER REFERENCES Users(id),
balance VARCHAR(50),
currency VARCHAR(50)
);
