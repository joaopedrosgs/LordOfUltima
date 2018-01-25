package accountserver

import (
	"errors"
	"fmt"

	"github.com/joaopedrosgs/OpenLoU/configuration"

	"encoding/json"
	"net/http"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"time"
)

type AccountServer struct {
	context *log.Entry
}

type LoginAttempt struct {
	Email    string
	Password string
}

func New() (*AccountServer, error) {
	return &AccountServer{log.WithFields(log.Fields{"Entity": "Account Server"})}, nil

}

func (s *AccountServer) StartListening(address string) {
	// Index Handler
	http.HandleFunc("/login", s.loginHandler)
	http.HandleFunc("/register", s.registerHandler)
	err := http.ListenAndServe(address, nil)
	for err != nil {
		s.context.Error("Failed to listen: " + err.Error())
		s.context.Info("Trying again in 10 seconds...")
		time.Sleep(10 * time.Second)
		err = http.ListenAndServe(address, nil)

	}
	s.context.Info("Account server has started listening")
}
func (s *AccountServer) loginHandler(writer http.ResponseWriter, request *http.Request) {
	println("aloi")
	if request.Method == "POST" {
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		attempt := &LoginAttempt{email, password}
		answer := s.NewAttempt(attempt)
		jsonAnswer, _ := json.Marshal(answer)
		fmt.Fprintf(writer, string(jsonAnswer))

	}

}
func (s *AccountServer) registerHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		login := request.PostFormValue("login")
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		err := s.CreateAccount(login, email, password)
		answer := communication.Answer{}
		if err != nil {
			answer.Data = err.Error()
		} else {
			answer.Data = success
			answer.Ok = true
		}
		jsonAnswer, _ := json.Marshal(answer)
		fmt.Fprintf(writer, string(jsonAnswer))
	}
}

// New returns an AccountServer that deals with the authentication of the user

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *AccountServer) NewAttempt(attempt *LoginAttempt) (answer *communication.Answer) {
	answer = &communication.Answer{}
	id, err := s.CheckCredentials(attempt)

	if err != nil {
		answer.Data = err.Error()
		return
	}
	key, err := GenerateRandomString(configuration.GetSingleton().Parameters.Security.KeySize)
	if err != nil {
		answer.Data = InternalError
		return
	}
	created := session.NewSession(id, key)
	if created {
		answer.Ok = true
		answer.Data = key
	}

	return
}

//CheckCredentials returns the user and nil if the credentials are correct
func (s *AccountServer) CheckCredentials(attempt *LoginAttempt) (uint, error) {
	if len(attempt.Password) == 0 || len(attempt.Email) == 0 {
		return 0, errors.New(emptyFields)
	}
	user, err := database.GetUser(attempt.Email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(attempt.Password))
	if err != nil {
		return 0, err
	}
	return user.ID, nil

}
func (s *AccountServer) CreateAccount(login string, email string, password string) error {

	if len(login) < 6 || len(email) < 8 || len(password) < 8 {
		return errors.New(shortCredentials)
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New(InternalError)
	}
	err = database.CreateUser(login, string(passwordHash), email)
	if err != nil {
		return errors.New(accountExists)
	}
	return nil

}
