package gateway

import (
	"sync"
	"sync/atomic"
)

// Session manages the state of a Discord gateway session.
type Session struct {
	sessionID string
	sequence  atomic.Int64
	resumeURL string
	identified atomic.Bool

	mu sync.RWMutex
}

// NewSession creates a new session.
func NewSession() *Session {
	return &Session{}
}

// SetSessionID stores the session ID from the READY event.
func (s *Session) SetSessionID(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessionID = id
}

// SessionID returns the current session ID.
func (s *Session) SessionID() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sessionID
}

// SetSequence updates the sequence number.
func (s *Session) SetSequence(seq int) {
	s.sequence.Store(int64(seq))
}

// Sequence returns the current sequence number.
func (s *Session) Sequence() int {
	return int(s.sequence.Load())
}

// IncrementSequence atomically increments the sequence number and returns the new value.
func (s *Session) IncrementSequence() int {
	return int(s.sequence.Add(1))
}

// SetResumeURL stores the resume gateway URL from the READY event.
func (s *Session) SetResumeURL(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.resumeURL = url
}

// ResumeURL returns the resume gateway URL.
func (s *Session) ResumeURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.resumeURL
}

// MarkIdentified marks the session as identified.
func (s *Session) MarkIdentified() {
	s.identified.Store(true)
}

// IsIdentified returns whether the session has been identified.
func (s *Session) IsIdentified() bool {
	return s.identified.Load()
}

// CanResume returns whether the session can be resumed.
// A session can be resumed if it has a session ID, a resume URL, and has been identified.
func (s *Session) CanResume() bool {
	if !s.IsIdentified() {
		return false
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.sessionID != "" && s.resumeURL != ""
}

// Reset clears the session state.
func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessionID = ""
	s.resumeURL = ""
	s.sequence.Store(0)
	s.identified.Store(false)
}
