package session

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/pkg/errors"
)

type Session struct {
	UserName   string
	lastAction time.Time
	tries      int
	conn       *websocket.Conn
}

type sessionMem struct {
	mutex    sync.RWMutex
	sessions map[string]*Session
}

var sessionsStorage sessionMem

func NewSession(user models.User) (string, error) {
	key, err := GenerateRandomString(configuration.GetSingleton().Parameters.Security.KeySize)
	if err == nil {
		sessionsStorage.mutex.Lock()
		if uint(len(key)) == configuration.GetSingleton().Parameters.Security.KeySize || len(user.Name) >= 0 {
			sessionsStorage.sessions[key] = &Session{user.Name, time.Now(), 0, nil}
		} else {
			err = errors.New("Internal error")
		}
		sessionsStorage.mutex.Unlock()
	}
	return key, err
}

func SetConn(key string, conn *websocket.Conn) {
	sessionsStorage.mutex.Lock()
	session, ok := sessionsStorage.sessions[key]
	if ok {
		session.conn = conn
	}
	sessionsStorage.mutex.Unlock()
}

func Exists(key string) bool {
	sessionsStorage.mutex.RLock()
	_, ok := sessionsStorage.sessions[key]
	sessionsStorage.mutex.RUnlock()
	return ok
}
func GetSession(key string) (*Session, bool) {
	sessionsStorage.mutex.RLock()
	session, ok := sessionsStorage.sessions[key]
	sessionsStorage.mutex.RUnlock()
	return session, ok
}
func DeleteSession(key string) {
	sessionsStorage.mutex.Lock()
	delete(sessionsStorage.sessions, key)
	sessionsStorage.mutex.Unlock()

}

func DeleteSessionByName(userName string) {
	sessionsStorage.mutex.Lock()
	for key, session := range sessionsStorage.sessions {
		if session.UserName == userName {
			delete(sessionsStorage.sessions, key)
		}
	}
	sessionsStorage.mutex.Unlock()

}
func NewSessionInMemory() {
	sessionsStorage = sessionMem{sync.RWMutex{}, make(map[string]*Session)}
}

func NewTry(key string) {
	sessionsStorage.mutex.RLock()
	if session, ok := sessionsStorage.sessions[key]; ok {
		session.tries++
	}
	sessionsStorage.mutex.RUnlock()
}

func GetUserConn(key string) (*websocket.Conn, bool) {
	sessionsStorage.mutex.RLock()
	session, ok := sessionsStorage.sessions[key]
	if ok {
		sessionsStorage.mutex.RUnlock()
		return session.conn, ok

	}
	sessionsStorage.mutex.RUnlock()
	return nil, ok

}
func GetUserConnByName(userName string) (*websocket.Conn, bool) {
	sessionsStorage.mutex.RLock()
	for _, session := range sessionsStorage.sessions {
		if session.UserName == userName {
			sessionsStorage.mutex.RUnlock()
			return session.conn, true
		}
	}
	sessionsStorage.mutex.RUnlock()
	return nil, false

}

func GetUserName(key string) (string, error) {
	sessionsStorage.mutex.RLock()
	user, found := sessionsStorage.sessions[key]
	name := user.UserName
	sessionsStorage.mutex.RUnlock()
	if !found {
		return "", errors.New("account or session not found")
	}
	return name, nil
}
