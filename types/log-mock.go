package types

// MockAPILogProvider - Mock provider logs
type MockAPILogProvider struct {
	OnInfo  func(info string)
	OnError func(trace string, erro error)
}

// Info - Mock provider log info
func (m *MockAPILogProvider) Info(info string) {
	m.OnInfo(log)
}

// Error - Mock provider log info
func (l *MockAPILogProvider) Error(trace string, erro error) {
	l.OnError(trace, erro)
}
