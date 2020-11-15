package main

// StateInfo struct.
type StateInfo struct {
	// 状态名称
	name string
}

// Name func.
func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) setName(name string) {
	s.name = name
}

// EnableSameTransit func.
func (s *StateInfo) EnableSameTransit() bool {
	return false
}

// OnBegin func.
func (s *StateInfo) OnBegin() {
}

// OnEnd func.
func (s *StateInfo) OnEnd() {
}

// CanTransitTo func.
func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}
