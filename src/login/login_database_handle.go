package main
/*
PostgreSQL Database Handler for Logins
*/

import (
	"errors"
	"context"
	"fmt"
	"os"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LoginAttempt struct {
	username string
	password string
}

func (l *LoginAttempt) ValidateAttempt(dbpool *pgxpool.Pool) error {
	var hashed_password string
	
	err := dbpool.QueryRow(context.Background(), 
		"select password from username where username=$1",
		l.username)
	.Scan(&hashed_password)
	
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

func startDBConnection() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	
	if err != nil {
		return nil, err
	} else {
		return dgpool, nill
	}

}