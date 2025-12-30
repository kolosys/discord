package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kolosys/helix"
)

type Manager struct {
	httpServer  *helix.Server
	connectFn   func(context.Context) error
	stopFn      func(context.Context) error
	httpEnabled bool

	mu      sync.RWMutex
	running bool
	once    sync.Once
}

func New(opts *Options) (*Manager, error) {
	if opts == nil {
		return nil, fmt.Errorf("server: options required")
	}
	if opts.ConnectFn == nil || opts.StopFn == nil {
		return nil, ErrNoCallbacks
	}

	m := &Manager{
		connectFn:   opts.ConnectFn,
		stopFn:      opts.StopFn,
		httpEnabled: opts.Enabled,
	}

	if opts.Enabled {
		opts.applyDefaults()
		m.httpServer = helix.Default(opts.HelixOptions)
	}

	return m, nil
}

func (m *Manager) Start(ctx context.Context) error {
	m.mu.Lock()
	if m.running {
		m.mu.Unlock()
		return ErrAlreadyRunning
	}
	m.running = true
	m.mu.Unlock()

	if m.httpEnabled {
		return m.startWithHTTP(ctx)
	}
	return m.startWithoutHTTP(ctx)
}

func (m *Manager) startWithHTTP(ctx context.Context) error {
	m.httpServer.OnStart(func(s *helix.Server) {
		if err := m.connectFn(ctx); err != nil {
			fmt.Printf("server: failed to connect: %v\n", err)
			go func() { s.Shutdown(context.Background()) }()
		}
	})

	m.httpServer.OnStop(func(shutdownCtx context.Context, s *helix.Server) {
		m.mu.RLock()
		running := m.running
		m.mu.RUnlock()
		if running {
			if err := m.stopFn(shutdownCtx); err != nil {
				fmt.Printf("server: error during shutdown: %v\n", err)
			}
			m.mu.Lock()
			m.running = false
			m.mu.Unlock()
		}
	})

	return m.httpServer.Start()
}

func (m *Manager) startWithoutHTTP(ctx context.Context) error {
	if err := m.connectFn(ctx); err != nil {
		m.mu.Lock()
		m.running = false
		m.mu.Unlock()
		return err
	}

	return m.waitForShutdown()
}

func (m *Manager) waitForShutdown() error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return m.stopAndCleanup(shutdownCtx)
}

func (m *Manager) Stop(ctx context.Context) error {
	m.mu.Lock()
	if !m.running {
		m.mu.Unlock()
		return nil
	}
	m.mu.Unlock()

	if m.httpEnabled {
		return m.httpServer.Shutdown(ctx)
	}

	return m.stopAndCleanup(ctx)
}

func (m *Manager) stopAndCleanup(ctx context.Context) error {
	var err error
	m.once.Do(func() {
		err = m.stopFn(ctx)
		m.mu.Lock()
		m.running = false
		m.mu.Unlock()
	})
	return err
}

func (m *Manager) IsRunning() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.running
}

func (m *Manager) HTTPServer() *helix.Server {
	return m.httpServer
}
