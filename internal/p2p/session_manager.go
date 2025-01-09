package p2p

import (
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type TssSession interface {
	Id() string
	Receive(request *SubmitRequest) error
	RegisterIdChangeListener(func(oldId, newId string))
}

type SessionManager struct {
	sessions map[string]TssSession
	mu       sync.RWMutex
}

func NewSessionManager(sessions ...TssSession) *SessionManager {
	manager := &SessionManager{
		sessions: make(map[string]TssSession),
	}

	for _, session := range sessions {
		manager.Add(session)
	}

	return manager
}

func (m *SessionManager) Add(session TssSession) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[session.Id()] = session
	session.RegisterIdChangeListener(m.onIdChange)
}

func (m *SessionManager) Get(id string) TssSession {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if session, exists := m.sessions[id]; exists {
		return session
	}

	return nil
}

func (m *SessionManager) Receive(request *SubmitRequest) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if session, exists := m.sessions[request.SessionId]; exists {
		return session.Receive(request)
	}

	return ErrSessionNotFound
}

func (m *SessionManager) onIdChange(oldId, newId string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, ok := m.sessions[oldId]
	if !ok {
		return
	}

	delete(m.sessions, oldId)
	m.sessions[newId] = session
	// id change listener remains the same
}