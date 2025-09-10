package main
/*
PostgreSQL Database Handler for Logins
*/

import (
	"errors"
	"context"
	"os"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginAttempt struct {
	username string
	password string
}

func (l *LoginAttempt) ValidateAttempt(dbpool *pgxpool.Pool) error {
	var hashed_password string
	
	err := dbpool.QueryRow(
		context.Background(), 
		"SELECT password FROM users WHERE username=$1",
		l.username,
	).Scan(&hashed_password)
	
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(l.password))
	
	if err != nil {
		return errors.New("Invalid Password")
	} else {
		return nil // All good login
	}
	
}

func (l *LoginAttempt) AddtoDB(dbpool *pgxpool.Pool) error {
	var found_username string
	hashed_password, crypt_err := bcrypt.GenerateFromPassword([]byte(l.password),bcrypt.DefaultCost)

	if crypt_err != nil {
		return errors.New("Could not encrypt password")
	}

	found := dbpool.QueryRow(
		context.Background(),
		"SELECT username FROM users WHERE username=$1",
		l.username,
	).Scan(&found_username)

	if found == pgx.ErrNoRows {
		_, err := dbpool.Exec(context.Background(), 
			"INSERT INTO users (username, password) VALUES ($1, $2)",
			l.username,
			hashed_password,
		)

		if err != nil {
			return err
		}

		if err != nil {
			return errors.New("Issue with database")
		} else {
			return nil // All good create user
		}	
	} else {
		return errors.New("Username already exists")
	}
}

func startDBConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	
	if err != nil {
		return nil, err
	} else {
		return dbpool, nil
	}

}