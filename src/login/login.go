package main
/*
This service is meant to provide login functionality for the platform
It stores usernames and hashed passwords in a postgress sql database
*/
import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5"
	"strconv"
)

func loginHandler(dgpool *pgxpool.Pool) http.HandlerFunc{
	return func (w http.ResponseHandler, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		create := strconv.ParseBool(r.FormValue("create"))
		l := LoginAttempt{username,password}

		if !create {
			err := l.ValidateAttempt(dgpool)

			if err == nil {
				// todo actually write this part
				w.Write([]byte("login successful"))
			} else {
				w.Write([]byte("login unsuccessful"))
			}
		} else {
			
		}
	}
}

func main() {
	// Start up http

}
