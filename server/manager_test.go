package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kolosys/discord/server"
	"github.com/kolosys/helix"
)

func TestManagerNew(t *testing.T) {
	tests := []struct {
		name    string
		opts    *server.Options
		wantErr bool
		errVal  error
	}{
		{
			name:    "nil options",
			opts:    nil,
			wantErr: true,
		},
		{
			name: "missing callbacks",
			opts: &server.Options{
				Enabled: false,
			},
			wantErr: true,
			errVal:  server.ErrNoCallbacks,
		},
		{
			name: "valid options without HTTP",
			opts: &server.Options{
				ConnectFn: func(context.Context) error { return nil },
				StopFn:    func(context.Context) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "valid options with HTTP",
			opts: &server.Options{
				Enabled:   true,
				ConnectFn: func(context.Context) error { return nil },
				StopFn:    func(context.Context) error { return nil },
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := server.New(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && tt.errVal != nil && !errors.Is(err, tt.errVal) {
				t.Errorf("New() error = %v, want %v", err, tt.errVal)
			}
			if !tt.wantErr && m == nil {
				t.Error("New() returned nil manager for valid options")
			}
		})
	}
}

func TestManagerHTTPServer(t *testing.T) {
	t.Run("HTTP disabled returns nil", func(t *testing.T) {
		m, _ := server.New(&server.Options{
			ConnectFn: func(context.Context) error { return nil },
			StopFn:    func(context.Context) error { return nil },
		})

		if m.HTTPServer() != nil {
			t.Error("HTTPServer() should return nil when HTTP disabled")
		}
	})

	t.Run("HTTP enabled returns server", func(t *testing.T) {
		m, _ := server.New(&server.Options{
			Enabled:   true,
			ConnectFn: func(context.Context) error { return nil },
			StopFn:    func(context.Context) error { return nil },
		})

		if m.HTTPServer() == nil {
			t.Error("HTTPServer() should return server when HTTP enabled")
		}
	})
}

func TestManagerIsRunning(t *testing.T) {
	m, _ := server.New(&server.Options{
		ConnectFn: func(context.Context) error { return nil },
		StopFn:    func(context.Context) error { return nil },
	})

	if m.IsRunning() {
		t.Error("New manager should not be running")
	}
}

func TestManagerConnectError(t *testing.T) {
	connectErr := errors.New("connect failed")
	m, _ := server.New(&server.Options{
		ConnectFn: func(context.Context) error { return connectErr },
		StopFn:    func(context.Context) error { return nil },
	})

	err := m.Start(context.Background())
	if !errors.Is(err, connectErr) {
		t.Errorf("Start() error = %v, want %v", err, connectErr)
	}
}

func TestManagerStopNonRunning(t *testing.T) {
	m, _ := server.New(&server.Options{
		ConnectFn: func(context.Context) error { return nil },
		StopFn:    func(context.Context) error { return nil },
	})

	err := m.Stop(context.Background())
	if err != nil {
		t.Errorf("Stop() error = %v, want nil", err)
	}
}

func TestManagerOptions(t *testing.T) {
	t.Run("applyDefaults sets addr", func(t *testing.T) {
		opts := &server.Options{
			Enabled:   true,
			ConnectFn: func(context.Context) error { return nil },
			StopFn:    func(context.Context) error { return nil },
		}
		m, _ := server.New(opts)

		httpServer := m.HTTPServer()
		if httpServer == nil {
			t.Fatal("HTTPServer is nil")
		}
		if httpServer.Addr() != ":8080" {
			t.Errorf("Default Addr = %q, want %q", httpServer.Addr(), ":8080")
		}
	})

	t.Run("custom addr is preserved", func(t *testing.T) {
		opts := &server.Options{
			Enabled: true,
			HelixOptions: &helix.Options{
				Addr: ":3000",
			},
			ConnectFn: func(context.Context) error { return nil },
			StopFn:    func(context.Context) error { return nil },
		}
		m, _ := server.New(opts)

		httpServer := m.HTTPServer()
		if httpServer == nil {
			t.Fatal("HTTPServer is nil")
		}
		if httpServer.Addr() != ":3000" {
			t.Errorf("Custom Addr = %q, want %q", httpServer.Addr(), ":3000")
		}
	})
}
