package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStore interface
type SessionStore interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
}

// Implements SessionStore interface.
type MockStore struct {
	Session *sessions.Session
	Err     error
}

func (m *MockStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return m.Session, m.Err
}