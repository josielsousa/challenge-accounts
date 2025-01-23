package types

// MockAPILogProvider - Mock provider logs.
type MockAPILogProvider struct {
	OnInfo  func(info string)
	OnError func(trace string, erro error)
}

// Info - Mock provider log info.
func (m *MockAPILogProvider) Info(info string) {
	m.OnInfo(info)
}

// Error - Mock provider log info.
func (m *MockAPILogProvider) Error(trace string, erro error) {
	m.OnError(trace, erro)
}
