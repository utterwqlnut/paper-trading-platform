package main
/*
This service is meant to provide login functionality for the platform
It stores usernames and hashed passwords in a postgress sql database
*/
import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
)

func loginHandler(dbpool *pgxpool.Pool) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		create,_ := strconv.ParseBool(r.FormValue("create"))

		l := LoginAttempt{username,password}

		if !create {
			err := l.ValidateAttempt(dbpool)

			if err == nil {
				w.Write([]byte("login succesful"))
				// do all the other stuff
			} else {
				w.Write([]byte(err.Error()))
			}
		} else {
			err := l.AddtoDB(dbpool)

			if err == nil {
				w.Write([]byte("user created succesfully"))

				// Do all the other stuff
			} else {
				w.Write([]byte(fmt.Sprintf(err.Error())))
			}
		}
	}
}

func main() {
	// Start up http
	dbpool, err := startDBConnection()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("DB Succesfully started")
	defer dbpool.Close()

	http.Handle("/login", loginHandler(dbpool))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
