package testutil

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/kolosys/discord/gateway"
)

// Assert provides testing assertion helpers.
type Assert struct {
	t *testing.T
}

// New creates a new assertion helper.
func New(t *testing.T) *Assert {
	return &Assert{t: t}
}

// Equal asserts that two values are equal.
func (a *Assert) Equal(expected, actual any, msg ...string) {
	a.t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		message := fmt.Sprintf("expected %v, got %v", expected, actual)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// NotEqual asserts that two values are not equal.
func (a *Assert) NotEqual(unexpected, actual any, msg ...string) {
	a.t.Helper()
	if reflect.DeepEqual(unexpected, actual) {
		message := fmt.Sprintf("expected not equal, got %v", actual)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// True asserts that a value is true.
func (a *Assert) True(value bool, msg ...string) {
	a.t.Helper()
	if !value {
		message := "expected true, got false"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// False asserts that a value is false.
func (a *Assert) False(value bool, msg ...string) {
	a.t.Helper()
	if value {
		message := "expected false, got true"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// Nil asserts that a value is nil.
func (a *Assert) Nil(value any, msg ...string) {
	a.t.Helper()
	if !isNil(value) {
		message := fmt.Sprintf("expected nil, got %v", value)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// NotNil asserts that a value is not nil.
func (a *Assert) NotNil(value any, msg ...string) {
	a.t.Helper()
	if isNil(value) {
		message := "expected not nil, got nil"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// NoError asserts that an error is nil.
func (a *Assert) NoError(err error, msg ...string) {
	a.t.Helper()
	if err != nil {
		message := fmt.Sprintf("unexpected error: %v", err)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// Error asserts that an error is not nil.
func (a *Assert) Error(err error, msg ...string) {
	a.t.Helper()
	if err == nil {
		message := "expected error, got nil"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// ErrorIs asserts that an error matches a sentinel error.
func (a *Assert) ErrorIs(err, target error, msg ...string) {
	a.t.Helper()
	if err != target {
		message := fmt.Sprintf("expected error %v, got %v", target, err)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// Contains asserts that a string contains a substring.
func (a *Assert) Contains(haystack, needle string, msg ...string) {
	a.t.Helper()
	if !strings.Contains(haystack, needle) {
		message := fmt.Sprintf("expected %q to contain %q", haystack, needle)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// NotContains asserts that a string does not contain a substring.
func (a *Assert) NotContains(haystack, needle string, msg ...string) {
	a.t.Helper()
	if strings.Contains(haystack, needle) {
		message := fmt.Sprintf("expected %q not to contain %q", haystack, needle)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// SliceContains asserts that a slice contains an element.
func (a *Assert) SliceContains(slice any, element any, msg ...string) {
	a.t.Helper()
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		a.t.Fatal("SliceContains: argument must be a slice")
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), element) {
			return
		}
	}

	message := fmt.Sprintf("expected slice to contain %v", element)
	if len(msg) > 0 {
		message = msg[0] + ": " + message
	}
	a.t.Error(message)
}

// SliceLen asserts that a slice has a specific length.
func (a *Assert) SliceLen(slice any, expectedLen int, msg ...string) {
	a.t.Helper()
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		a.t.Fatal("SliceLen: argument must be a slice")
	}

	if v.Len() != expectedLen {
		message := fmt.Sprintf("expected slice length %d, got %d", expectedLen, v.Len())
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// GatewayPayloadOpcode asserts that a payload has a specific opcode.
func (a *Assert) GatewayPayloadOpcode(payload *gateway.GatewayPayload, expectedOp gateway.Opcode, msg ...string) {
	a.t.Helper()
	if payload.Op != expectedOp {
		message := fmt.Sprintf("expected opcode %v, got %v", expectedOp, payload.Op)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// GatewayPayloadEventType asserts that a dispatch payload has a specific event type.
func (a *Assert) GatewayPayloadEventType(payload *gateway.GatewayPayload, expectedType string, msg ...string) {
	a.t.Helper()
	if payload.T == nil {
		message := "expected event type, got nil"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
		return
	}

	if *payload.T != expectedType {
		message := fmt.Sprintf("expected event type %s, got %s", expectedType, *payload.T)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// GatewayPayloadHasSequence asserts that a payload has a sequence number.
func (a *Assert) GatewayPayloadHasSequence(payload *gateway.GatewayPayload, msg ...string) {
	a.t.Helper()
	if payload.S == nil {
		message := "expected sequence number, got nil"
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// GatewayPayloadSequence asserts that a payload has a specific sequence number.
func (a *Assert) GatewayPayloadSequence(payload *gateway.GatewayPayload, expectedSeq int, msg ...string) {
	a.t.Helper()
	if payload.S == nil {
		message := fmt.Sprintf("expected sequence %d, got nil", expectedSeq)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
		return
	}

	if *payload.S != expectedSeq {
		message := fmt.Sprintf("expected sequence %d, got %d", expectedSeq, *payload.S)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// RequestMethodPath asserts that a received request matches method and path.
func (a *Assert) RequestMethodPath(req *ReceivedRequest, expectedMethod, expectedPath string, msg ...string) {
	a.t.Helper()
	if req.Method != expectedMethod || req.Path != expectedPath {
		message := fmt.Sprintf("expected %s %s, got %s %s", expectedMethod, expectedPath, req.Method, req.Path)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// RequestHasHeader asserts that a request contains a specific header.
func (a *Assert) RequestHasHeader(req *ReceivedRequest, headerName string, msg ...string) {
	a.t.Helper()
	if req.Header.Get(headerName) == "" {
		message := fmt.Sprintf("expected header %s to be present", headerName)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// RequestHeaderEquals asserts that a request header has a specific value.
func (a *Assert) RequestHeaderEquals(req *ReceivedRequest, headerName, expectedValue string, msg ...string) {
	a.t.Helper()
	actual := req.Header.Get(headerName)
	if actual != expectedValue {
		message := fmt.Sprintf("expected header %s=%s, got %s", headerName, expectedValue, actual)
		if len(msg) > 0 {
			message = msg[0] + ": " + message
		}
		a.t.Error(message)
	}
}

// Fail fails the test with a message.
func (a *Assert) Fail(msg string) {
	a.t.Helper()
	a.t.Error(msg)
}

// Fatalf fails the test with a formatted message and stops execution.
func (a *Assert) Fatalf(format string, args ...any) {
	a.t.Helper()
	a.t.Fatalf(format, args...)
}

// isNil checks if a value is nil, handling interface{} types correctly
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}
