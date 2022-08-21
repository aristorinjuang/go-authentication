package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/entity"
	"github.com/aristorinjuang/go-authentication/internal/factory/dto"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db *sql.DB
}

func (m *mysql) Get(email *valueobject.Email) (*entity.User, error) {
	rows, err := m.db.Query(
		"SELECT first_name, last_name, hash FROM users WHERE email = ? LIMIT 1",
		email.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := dto.NewUser(email)
	for rows.Next() {
		rows.Scan(&user.FirstName, &user.LastName, &user.Hash)
	}

	return user.Entity(), nil
}

func (m *mysql) Create(user *entity.User) error {
	_, err := m.db.Exec(
		"INSERT INTO users(email, first_name, last_name, hash, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())",
		user.Email.String(),
		user.Name.First,
		user.Name.Last,
		user.Password.Hash,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewMySQL(db *sql.DB) *mysql {
	return &mysql{db}
}

func ConnectMySQL(database *config.Database) (*sql.DB, error) {
	db, err := sql.Open("mysql", database.Source)
	if err != nil {
		return nil, err
	}

	if database.MaxIdleConns > 0 {
		db.SetMaxIdleConns(database.MaxIdleConns)
	}
	if database.MaxOpenConns > 0 {
		db.SetMaxOpenConns(database.MaxOpenConns)
	}
	if database.ConnMaxLifetime.Nanoseconds() > 0 {
		db.SetConnMaxLifetime(database.ConnMaxLifetime)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
