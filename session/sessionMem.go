package session

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/models"
)

type sessionMem struct {
	mutex    sync.RWMutex
	sessions map[string]*models.Session
}

var sessionsStorage sessionMem

func NewSession(usr *models.User, conn *websocket.Conn) (*models.Session, error) {
	key, err := GenerateRandomString(configuration.GetSingleton().Parameters.Security.KeySize)
	if err != nil {
		return nil, err
	}
	sessionsStorage.mutex.Lock()
	session := &models.Session{User: usr, LastAction: time.Now(), Conn: conn}
	sessionsStorage.sessions[key] = session
	sessionsStorage.mutex.Unlock()

	return session, nil
}

func Exists(key string) bool {
	sessionsStorage.mutex.RLock()
	_, ok := sessionsStorage.sessions[key]
	sessionsStorage.mutex.RUnlock()
	return ok
}
func GetSession(key string) (*models.Session, bool) {
	sessionsStorage.mutex.RLock()
	session, ok := sessionsStorage.sessions[key]
	sessionsStorage.mutex.RUnlock()
	return session, ok
}
func CloseSession(key string) {
	sessionsStorage.mutex.Lock()
	sessionsStorage.sessions[key].Close()
	delete(sessionsStorage.sessions, key)
	sessionsStorage.mutex.Unlock()

}

func DeleteSessionByUser(user *models.User) {
	sessionsStorage.mutex.Lock()
	for key, session := range sessionsStorage.sessions {
		if session.User == user {

			delete(sessionsStorage.sessions, key)
		}
	}
	sessionsStorage.mutex.Unlock()

}
func NewSessionInMemory() {
	sessionsStorage = sessionMem{sync.RWMutex{}, make(map[string]*models.Session)}
}
