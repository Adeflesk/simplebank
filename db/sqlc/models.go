// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `db:"id"`
	Owner     string    `db:"owner"`
	Balance   int64     `db:"balance"`
	Currency  string    `db:"currency"`
	CreatedAt time.Time `db:"created_at"`
}

type Entry struct {
	ID        int64 `db:"id"`
	AccountID int64 `db:"account_id"`
	// It can be negative or positive
	Amount   int64        `db:"amount"`
	CreateAt sql.NullTime `db:"create_at"`
}

type Transfer struct {
	ID            int64 `db:"id"`
	FromAccountID int64 `db:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id"`
	// must be postive
	Amount    int64     `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}

type User struct {
	Username          string    `db:"username"`
	HashedPassword    string    `db:"hashed_password"`
	FullName          string    `db:"full_name"`
	Email             string    `db:"email"`
	PasswordChangedAt time.Time `db:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at"`
}
