package assert

import (
	"fmt"
	"testing"
)

type mockTestingT struct {
	failures []string
}

func (m *mockTestingT) Fatalf(format string, args ...interface{}) {
	m.failures = append(m.failures, fmt.Sprintf(format, args...))
}

func TestDeepEqual(t *testing.T) {
	tt := &mockTestingT{}
	DeepEqual(tt, []string{"a", "b", "c"}, []string{"e", "f"})
	Contains(t, tt.failures[0], "([]string) (len=2 cap=2)")
}
